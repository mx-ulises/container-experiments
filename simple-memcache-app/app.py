import os

from flask import Flask, jsonify, abort, request
from pymemcache.client.base import Client

app = Flask(__name__)
memcache_client = None

@app.errorhandler(404)
def resource_not_found(e):
    # Return errors as Json objects, in this case the 404 errors
    return jsonify(error=str(e)), 404


@app.route('/keys/<key>', methods=['GET', 'POST'])
def serve_key_request(key: str):
    # Request handler for the keys, we can do GET and POST
    if request.method == 'GET':
        value = memcache_client.get(key)
        if value:
            return jsonify({key: value.decode()})
        else:
            abort(404, description='Key not found.')
    if request.method == 'POST':
        value = request.form.get('value')
        memcache_client.set(key, value, expire=60)
        return jsonify({key: value})


if __name__ == '__main__':
    # Initialize Memcache client
    memcache_host = os.environ.get('MEMCACHE_HOST', '127.0.0.1')
    memcache_port = int(os.environ.get('MEMCACHE_PORT', 11211))
    memcache_client = Client((memcache_host, memcache_port))
    # Initialize Flask App
    http_port = int(os.environ.get('HTTP_PORT', 8080))
    app.run(debug=True, host='0.0.0.0', port=http_port)
