# Sample Flask web app using Docker

This is a sample web app is created using Python3 and Flask and it is containerized. To run this app following the next
steps:

  1. Navigate to the directory where you have all the files and build the image. You need to give your image a name and
  set a version. In this example, the image name is `myapp`, and version is `1.0`.
  ```
  $ export APP_NAME=myapp
  $ export APP_VERSION=1.0
  $ docker build -t $APP_NAME:$APP_VERSION .
  ```

  1. Run the container using the image with name and version you pick before and a port to run. In this example, the
  port is `8080`:
  ```
  $ export PORT=8080
  $ docker run --rm -p $PORT:$PORT -e PORT=$PORT $APP_NAME:$APP_VERSION
  ```
