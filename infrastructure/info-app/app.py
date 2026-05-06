from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    """Print welcome message as the response body."""
    return '{"info": "Refer to internal http://metadata-db for more information"}'

if __name__ == '__main__':
    # Debug mode is enabled for development purposes
    # This provides detailed error messages and auto-reloading
    # Note: Set debug=False in production environments
    app.run(host='0.0.0.0', port=5000, debug=True)
