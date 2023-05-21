import os

from flask import Flask

app = Flask(__name__)
HOSTNAME = os.environ.get('HOSTNAME', 'Default')

@app.route('/')
def home():
    return f'{HOSTNAME} - 我喜欢中国菜'

@app.route('/ready')
def ready():
    return 'Ready'

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 8080))
    app.run(debug=True, host='0.0.0.0', port=port)
