from flask import Flask, request
import subprocess
import os

app = Flask(__name__)

@app.route('/', methods=['POST'])
def convert_audio():
    file = request.files['file']
    output_format = request.form.get('format', 'mp3')
    file.save('/tmp/input_file')
    subprocess.run(['ffmpeg', '-i', '/tmp/input_file', f'/tmp/output.{output_format}'])
    return send_file(f'/tmp/output.{output_format}', attachment_filename=f'output.{output_format}')
