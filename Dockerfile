FROM python:3.11.4-alpine3.18

RUN apk update && \
    apk add --no-cache ffmpeg

COPY . /app
WORKDIR /app
RUN pip install --no-cache-dir -r requirements.txt

CMD exec uvicorn --host 0.0.0.0 --port $PORT main:app
