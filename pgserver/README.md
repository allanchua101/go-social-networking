# PostgreSQL-bassed Event Store

This guide explains the purpose of this section and how to work around the application's event store.

## Pre-requisites

You need to have the following software installed on your machine.

- Linux / Mac Machine (Because of mounting issues associated with Windows.)
- Docker
- Docker Compose

### Containerized debugging database

Utilising a containerized database offer the following advantages:

- Host machine is safe from installation pollution.
- Compatibility issues associated with other applications being developed on other PostgreSQL versions.
- No vendor lock-in.
- Some projects have use-cases where data should stay in certain geographical locations.

## Launching PostgreDB Container

To bootup the PostgreDB container, run `docker-compose up` locally.

```sh
docker-compose up

# You could also run in detached mode
# using the code below

docker-compose up -d
```

## Configuring the credentials

To configure the credentials for your local setup, you can navigate to the `docker-compose.yml` file and manipulate the following configurations:

```yaml
environment:
  POSTGRES_PASSWORD: godevrocks
  POSTGRES_USER: dev01
  POSTGRES_DB: gosocialdb
```

## Solving conflicting ports

If you are experiencing conflicting ports when booting up the PostgreDB container, tweak the following section in the `docker-compose.yml` file

```yaml
# Default port configuration that
# conflicts in your machine.
ports:
  - 5555:5432
```

to something that looks like

```yaml
# External port was replaced (5555 -> 5455), this enables you
# to solve conflicting ports caused by multiple consumers of 5555 ports.
ports:
  - 5455:5432
```
