# RabbitMQ as messaging medium.

RabbitMQ is selected for development simulation, in real-life you would want to consume a messaging medium from cloud providers (SQS and ) to offload scaling and maintenance of messaging infrastructure.

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
