#!/bin/sh

# Create Rabbitmq user
( sleep 30 ; \
rabbitmqctl add_user $RABBITMQ_USER $RABBITMQ_PASS 2>/dev/null ; \
sleep 10 ; \
rabbitmqctl set_user_tags $RABBITMQ_USER administrator ; \
sleep 10 ; \
rabbitmqctl set_permissions -p / $RABBITMQ_USER  ".*" ".*" ".*" ; \
echo "*** User '$RABBITMQ_USER' with password '$RABBITMQ_PASS' completed. ***" ; \
echo "*** Log in the WebUI at port 15672 (example: http:/localhost:15672) ***" ; \
rabbitmqadmin declare queue name=$WRITE_API_QUEUE_NAME  durable=true ; \
# Read Projections
rabbitmqadmin declare queue name=$MY_FEED_PROJECTOR_QUEUE_NAME durable=true ; \
rabbitmqadmin declare queue name=$FRIENDS_FEED_PROJECTOR_QUEUE_NAME durable=true ; \
# Exchange Declarations
rabbitmqadmin declare exchange name=$ACTIVITY_EXCHANGE_NAME type=fanout ; \
# Binding Declarations
rabbitmqadmin declare binding destination_type="queue" source=$ACTIVITY_EXCHANGE_NAME destination=$MY_FEED_PROJECTOR_QUEUE_NAME ; \
rabbitmqadmin declare binding destination_type="queue" source=$ACTIVITY_EXCHANGE_NAME destination=$FRIENDS_FEED_PROJECTOR_QUEUE_NAME) &

# $@ is used to pass arguments to the rabbitmq-server command.
# For example if you use it like this: docker run -d rabbitmq arg1 arg2,
# it will be as you run in the container rabbitmq-server arg1 arg2
rabbitmq-server $@