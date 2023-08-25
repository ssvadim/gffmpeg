from fastapi import FastAPI, UploadFile, Form, HTTPException
from starlette.responses import FileResponse
import subprocess

app = FastAPI()

@app.post('/')
async def convert_audio(file: UploadFile = File(...),
                        output_format: str = Form('mp3'),
                        params: str = Form('-i')):
    with open('/tmp/input_file', 'wb') as buffer:
        buffer.write(await file.read())
    ffmpeg_command = ['ffmpeg', '-y'] + params.split(" ") + ['/tmp/input_file', f'/tmp/output.{output_format}']
    process = subprocess.run(ffmpeg_command)
    if process.returncode != 0:
        raise HTTPException(status_code=500, detail="FFmpeg failed to convert the file.")
    return FileResponse(f'/tmp/output.{output_format}', filename=f'output.{output_format}')
