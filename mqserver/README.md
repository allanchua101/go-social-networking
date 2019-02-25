# MQ-SERVER for Topics

This folder is used as the persistent drive of the RabbitMQs used for pub/sub implementations.

## Connecting to a running RabbitMQ container

To connect to the RabbitMQ container, run the command below:

```sh
docker exec -it go-social-topic-mq /bin/bash
```

## Checking node health

To check the health of the MQ node, run the command below inside the container:

```sh
rabbitmqctl node_health_check
```
