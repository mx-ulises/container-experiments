# Sample node.js app (KUAR)

## Description

Sample `Hello World!` app from Kubernetes: Up and Running, chapter 2.

To build the container use:

```
docker build -t simple-node .
```

To Run the container use:

```
docker run --rm -p 3000:3000 simple-node
```
