FROM python:3.11-slim

RUN apt-get update && \
    apt-get install -y ffmpeg

COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt

CMD exec uvicorn --host 0.0.0.0 --port $PORT main:app
