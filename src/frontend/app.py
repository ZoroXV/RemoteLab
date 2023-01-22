from flask import Flask, render_template, request, url_for, redirect
import requests
import json

app = Flask(__name__)
raspberrypi_server_ip = 'RPI_IP'
raspberrypi_server_port = '8080'
raspberrypi_server_address = 'http://' + raspberrypi_server_ip + ':' + raspberrypi_server_port

@app.route('/')
def index():
    data_raw = requests.get(raspberrypi_server_address + '/command/list_controllers')
    data = data_raw.json()
    return render_template('index.html', data=data['data'])

@app.route('/choose_controller', methods=["POST"])
def choose_controller():
    is_update_list = request.form.get('update')
    if is_update_list:
        return redirect(url_for('index'), code=302)

    controller_raw = request.form.get('controller')

    if controller_raw is None:
        raise BaseException

    controller = json.loads(controller_raw.replace("'", '"'))
    return render_template('upload_command.html', controller=controller)

@app.route('/upload_command', methods=["POST"])
def upload_command():
    data = {
        'port': request.form.get('port'),
        'fqbn': request.form.get('fqbn'),
        'serial_number': request.form.get('serial_number'),
        'start_address': request.form.get('start_address'),
        'filename': request.form.get('filename'),
    }
    print(data)
    requests.post(raspberrypi_server_address +'/command/upload', json=data)
    return redirect(url_for('index'), code=302)

@app.route('/upload_binary', methods=["POST"])
def upload_binary():
    uploaded_file = request.files['file']
    if uploaded_file.filename != '':
        print(uploaded_file.filename)
        uploaded_file.save(uploaded_file.filename)

    data = {'name': uploaded_file.filename}
    files = {'file': open(uploaded_file.filename, 'rb')}
    requests.post(raspberrypi_server_address + '/command/uploadfile', data=data, files=files)
    return redirect(url_for('index'), code=302)

if __name__ == '__main__':
    app.run(debug=False, host='0.0.0.0')
