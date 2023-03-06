import os, time, sys

from flask import Flask, jsonify, abort, request
from werkzeug.exceptions import HTTPException
import etcd3

app = Flask(__name__)
etcd_client = None
etcd_lock = None
etcd_lock_key = None
counter = 0


@app.errorhandler(503)
def lock_not_available(e):
    # Return errors as Json objects, in this case the 500 errors
    return jsonify({'error': str(e), 'counter': counter}), 500

@app.errorhandler(408)
def too_long(e):
    # Return errors as Json objects, in this case the 400 errors
    return jsonify({'error': str(e), 'counter': counter}), 408


@app.route('/increase_counter', methods=['POST'])
def increase_counter():
    # Request to increase the counter once at the time
    global counter
    sleep_t = int(request.form.get('sleep'))
    # Initialize lock
    etcd_lock = etcd_client.lock(etcd_lock_key, ttl=etcd_lock_ttl)
    try:
        # Acquire the lock
        if etcd_lock.acquire(timeout=None):
            # Critical section - increase the counter after the sleep
            time.sleep(sleep_t)
            if etcd_lock.is_acquired():
                counter += 1
            else:
                abort(400)
    except Exception as e:
        # Error happened, let's check
        if isinstance(e, HTTPException):
            # Request provided a long sleep time, aborting with 4XX
            abort(408, description='Wait was too long')
        else:
            # Lock was in usage, aborting with 503
            abort(503, description='Function not available')
    finally:
        # Release lock (happen only if acquired)
        released = etcd_lock.release()
        print(f'release for sleep({sleep_t}): {released}', file=sys.stderr)
    return jsonify({"counter": counter})


@app.route('/read_counter', methods=['GET'])
def read_counter():
    # Request to return the current counter value
    return jsonify({"counter": counter})


if __name__ == '__main__':
    # Initialize etcd client
    etcd_host = os.environ.get('ETCD_HOST', '127.0.0.1')
    etcd_port = int(os.environ.get('ETCD_PORT', 2379))
    etcd_client = etcd3.client(host=etcd_host, port=etcd_port)
    # Initialize etcd Lock globals
    etcd_lock_key = os.environ.get('ETCD_LOCK_KEY', '/my/lock')
    etcd_lock_ttl = int(os.environ.get('ETCD_LOCK_TTL', 60))
    # Initialize Flask App
    http_port = int(os.environ.get('HTTP_PORT', 8080))
    app.run(debug=True, host='0.0.0.0', port=http_port)
