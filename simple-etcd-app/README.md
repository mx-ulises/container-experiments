  1. Build App image:
  ```
  export APP_NAME=myapp
  export APP_VERSION=1.0
  docker build -t $APP_NAME:$APP_VERSION .
  ```

  2. Create a container running etcd.
  ```
  export ETCD_CONTAINER_NAME=my-etcd
  export ETCD_PORT=2379
  docker run -d \
  --name $ETCD_CONTAINER_NAME \
  -p $ETCD_PORT:$ETCD_PORT \
  quay.io/coreos/etcd:v3.4.15 \
  /usr/local/bin/etcd \
  --name $ETCD_CONTAINER_NAME \
  --advertise-client-urls http://0.0.0.0:$ETCD_PORT \
  --listen-client-urls http://0.0.0.0:$ETCD_PORT
  ```

  3. Get IP Address:
  ```
  docker container inspect $ETCD_CONTAINER_NAME | grep IPAddress
  ```

  4. Deploy App container
  ```
  export HTTP_PORT=8080
  export ETCD_HOST=172.17.0.2
  export ETCD_PORT=2379
  export ETCD_LOCK_KEY=/my/lock
  export ETCD_LOCK_TTL=60
  docker run --rm -p $HTTP_PORT:$HTTP_PORT -e HTTP_PORT=$HTTP_PORT -e ETCD_HOST=$ETCD_HOST -e ETCD_PORT=$ETCD_PORT -e ETCD_LOCK_KEY=$ETCD_LOCK_KEY -e ETCD_LOCK_TTL=$ETCD_LOCK_TTL $APP_NAME:$APP_VERSION
  ```
  