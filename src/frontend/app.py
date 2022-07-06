import flask

app = flask.Flask(__name__)

@app.route('/')
def index():
    return "Hello from Flask"

if __name__ == '__main__':
    app.run(host='0.0.0.0')
