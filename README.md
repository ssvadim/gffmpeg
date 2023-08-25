# Media Converter

Media Converter is a simple web service built with FastAPI. It allows you to convert media files to different formats using FFmpeg. The application accepts a file to be converted, the desired output format, and optional FFmpeg parameters.

The application is designed to be deployed on Google Cloud Run, but can run anywhere that supports Python 3.7+.


## Usage

To use the service, make a POST request to the root endpoint ('/') with the following data:

- `file`: The file to be converted (multipart/form-data)
- `output_format` (form): The desired output format (default: 'mp3')
- `params` (form): Optional FFmpeg parameters (default: '-i')

Example using curl:

```bash
curl -X POST \
     -F "file=@/path/to/your/file" \
     -F "output_format=mp4" \
     -F "params=-i -vcodec h264" \
     http://localhost:8000/ --output output.mp4
```
Replace /path/to/your/file with the actual path to your file. The output file will be saved as output.mp4 in the current directory.

Note: This example assumes the application is running locally on port 8000. Update the URL if it's hosted elsewhere.

Deployment on Google Cloud Run
Follow Google's official documentation for building and deploying an application on Cloud Run.

Disclaimer
This application is provided as-is, without warranty of any kind. Use at your own risk. Always ensure you have the rights to convert and distribute any media files used.
