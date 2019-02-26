# GO DB Mongo Server

This guide describes critical information about simulation mongo server.

## Master Projection Connection String

Use the connection string below to connect to the main / master document store used for projection data.

```sh
mongodb://localhost:27018/go-social-db
```

## Slave Projection Connection String

Use the connection string below to connect to the slave / failover document store used for disaster failovers.

```sh
mongodb://localhost:27018/go-social-slave
```

## Connecting to the container

Use the following command to connect to the simulation container in your local machine:

```sh
docker exec -it go-social-mongo sh

# Run your standard mongo shell commands.
```
