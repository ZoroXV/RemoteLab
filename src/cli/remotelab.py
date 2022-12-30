#!/usr/bin/env python3

import argparse
import json
from urllib.error import HTTPError, URLError
from urllib.request import urlopen, Request
import binascii
import os
import sys

def parseargs():
    parser = argparse.ArgumentParser(prog='./remotelab.py', description='RemoteLab CLI tool.')

    commands = parser.add_subparsers(dest='commands')

    flash_command = commands.add_parser('flash')
    flash_command.add_argument('filename', nargs=1, help='the name of the file to flash on the microcontroller')
    flash_command.add_argument('-a', '--address', required=True, help='the ip address of the server')
    flash_command.add_argument('-f', '--fqbn', dest='fqbn', required=True, help='the type of the card, following the names of the `arduino-cli` (ex: "arduino:avr:uno")')
    flash_command.add_argument('-p', '--port', dest='port', required=True, help='the port on which the card is linked (ex: "/dev/ttyUSB0")')
    
    upload_file_command = commands.add_parser('upload')
    upload_file_command.add_argument('filepath', nargs='+', help='the full path of the file(s) to upload on the server')
    upload_file_command.add_argument('-a', '--address', required=True, help='the ip address of the server')

    list_controllers_command = commands.add_parser('list')
    list_controllers_command.add_argument('-a', '--address', required=True, help='the ip address of the server')

    return parser.parse_args()

def encode_multipart_formdata(filename, filecontent):
    boundary = binascii.hexlify(os.urandom(16)).decode('ascii')

    body = ('--%s\r\n'
            'Content-Disposition: form-data; name=\"name\"\r\n'
            '\r\n'
            '%s\r\n'
            '--%s\r\n'
            'Content-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\n'
            '\r\n'
            '%s\r\n'
            '--%s--\r\n' % (boundary, filename, boundary, filename, filecontent, boundary))

    content_type = 'multipart/form-data; boundary=%s' % boundary

    return body, content_type

def request(address, urlpath, method='GET', data=None, content_type=None):
    url = 'http://' + address + urlpath
    
    headers = {}
    if content_type != None:
        headers = {'Content-Type': content_type}
    if data != None:
        data = data.encode()

    httprequest = Request(url, method=method, data=data, headers=headers)

    try:
        with urlopen(httprequest) as response:
            return response.status, json.load(response)
    except HTTPError as err:
        return err.code, json.load(err)
    except URLError as err:
        return -1, err.reason


def flash(address, filename, port, fqbn):
    print('Flash', filename, 'on port', port, '... ')

    data = {
        'port': port,
        'fqbn': fqbn,
        'filename': filename
    }
    status, data = request(address, '/command/upload', 'POST', json.dumps(data), 'application/json')
    
    if status != 200 or data['status'] != 'OK':
        print('ERROR')
        print('' + data['message'], file=sys.stderr)
    else:
        print('OK')

def upload_file(address, filepath):
    print('Upload file', filepath, '... ', end='')

    with open(filepath) as fp:
        filecontent = fp.read()

    body, content_type = encode_multipart_formdata(os.path.basename(filepath), filecontent)
    status, data = request(address, '/uploadfile', 'POST', body, content_type)
    
    if status != 200 or data['status'] != 'OK':
        print('ERROR')
        print('' + data['message'], file=sys.stderr)
        return -1
    else:
        print('OK')
        return 0

def upload_files(address, filepaths):
    for filepath in filepaths:
        if upload_file(address, filepath) == -1:
            return

def list_controllers(address):
    status, data = request(address, '/command/list_controllers')
    
    if status != 200 or data['status'] != 'OK':
        print('ERROR')
        print('' + data['message'], file=sys.stderr)
    else:
        for dev in data['data']:
            print(
                '\033[91;4m%s:\033[0m\n'
                '\t* \033[94;4mVendor:\033[0m %s\n'
                '\t* \033[94;4mPort:\033[0m %s\n'
                '\t* \033[94;4mFqbn:\033[0m'
                 % (dev['product_name'], dev['vendor_name'], dev['port']))
            for fqbn in dev['fqbn']:
                print('\t\t* %s\n' % fqbn, end='')

def main ():
    args = parseargs()
    
    if args.commands == 'flash':
        flash(args.address, args.filename[0], args.port, args.fqbn)
    elif args.commands == 'upload':
        upload_files(args.address, args.filepath)
    elif args.commands == 'list':
        list_controllers(args.address)

if __name__ == '__main__':
    main()