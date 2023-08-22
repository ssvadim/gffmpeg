from fastapi import FastAPI, UploadFile, File, Form
from starlette.responses import FileResponse
import subprocess

app = FastAPI()

@app.post('/')
async def convert_audio(file: UploadFile = File(...), output_format: str = Form('mp3')):
    with open('/tmp/input_file', 'wb') as buffer:
        buffer.write(await file.read())
    subprocess.run(['ffmpeg', '-i', '/tmp/input_file', f'/tmp/output.{output_format}'])
    return FileResponse(f'/tmp/output.{output_format}', filename=f'output.{output_format}')
