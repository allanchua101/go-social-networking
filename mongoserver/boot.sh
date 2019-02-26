#!/bin/sh

# Wait for container to bootup
(sleep 30 ; \
mongo $PROJECTION_MAIN_NAME --eval "db.bootTable.insert({ Name : \"Seeded Main Projection\" })"; \
mongo $PROJECTION_SLAVE_NAME --eval "db.bootTable.insert({ Name : \"Seeded Slave Projection\" })";
) &

# $@ is used to pass arguments to the mongod command.
# For example if you use it like this: docker run -d mongod arg1 arg2,
# it will be as you run in the container mongod arg1 arg2
mongod $@
