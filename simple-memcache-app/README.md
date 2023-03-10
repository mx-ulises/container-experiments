
  1. Build App image:
  ```
  export APP_NAME=myapp
  export APP_VERSION=1.0
  docker build -t $APP_NAME:$APP_VERSION .
  ```

  1. Create a container running memcache.
  ```
  export MEMCACHE_CONTAINER_NAME=my-memcache
  docker run --name $MEMCACHE_CONTAINER_NAME -d memcached memcached -m 64
  ```

  1. Get IP Address:
  ```
  docker container inspect $MEMCACHE_CONTAINER_NAME | grep IPAddress
  ```

  1. Deploy App container
  ```
  export HTTP_PORT=8080
  export MEMCACHE_HOST=172.17.0.2
  export MEMCACHE_PORT=11211
  docker run --rm -p $HTTP_PORT:$HTTP_PORT -e HTTP_PORT=$HTTP_PORT -e MEMCACHE_HOST=$MEMCACHE_HOST -e MEMCACHE_PORT=$MEMCACHE_PORT $APP_NAME:$APP_VERSION
  ```
