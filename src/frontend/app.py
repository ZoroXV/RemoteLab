from flask import *
import requests

app = Flask(__name__)
raspberrypi_server_ip = 'RPI_IP'
raspberrypi_server_port = '8080'
raspberrypi_server_address = 'http://' + raspberrypi_server_ip + ':' + raspberrypi_server_port

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/upload_command', methods=["POST"])
def upload_command():
    data_form = request.form.to_dict(flat=False)
    data = {
        'port': data_form['port'][0],
        'fqbn': data_form['fqbn'][0],
        'filename': data_form['filename'][0]
    }
    print(data)
    response = requests.post(raspberrypi_server_address +'/command/upload', json=data)
    return redirect(url_for('index'), code=302)

@app.route('/upload_binary', methods=["POST"])
def upload_binary():
    uploaded_file = request.files['file']
    if uploaded_file.filename != '':
        print(uploaded_file.filename)
        uploaded_file.save(uploaded_file.filename)

    data = {'name': uploaded_file.filename}
    files = {'file': open(uploaded_file.filename, 'rb')}
    response = requests.post(raspberrypi_server_address + '/uploadfile', data=data, files=files)
    return redirect(url_for('index'), code=302)

if __name__ == '__main__':
    app.run(debug=False, host='0.0.0.0')
