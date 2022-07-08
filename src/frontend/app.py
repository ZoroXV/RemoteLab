from flask import *

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
    uploaded_file = request.files['file']
    if uploaded_file.filename != '':
        uploaded_file.save(uploaded_file.filename)
    return "Sucess"

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')