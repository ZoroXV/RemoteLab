from flask import  Flask, render_template, request, json

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/upload_command', methods=["POST"])
def upload_command():
    data = request.form
    return data

@app.route('/upload_binary', methods=["POST"])
def upload_binary():
    data = request.form
    return data

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')