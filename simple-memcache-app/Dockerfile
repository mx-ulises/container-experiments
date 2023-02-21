FROM python:3.9-slim-buster

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE $HTTP_PORT
EXPOSE $MEMCACHE_HOST
EXPOSE $MEMCACHE_PORT

CMD ["python", "app.py"]
