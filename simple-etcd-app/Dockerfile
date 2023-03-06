FROM python:3.9-slim-buster

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE $HTTP_PORT
EXPOSE $ETCD_HOST
EXPOSE $ETCD_PORT
EXPOSE $ETCD_LOCK_KEY
EXPOSE $ETCD_LOCK_TTL

ENV PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python

CMD ["python", "app.py"]
