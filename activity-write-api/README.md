# Testing Locally

## Building locally

Run the command below to run the on stand alone fashion

```sh
go install ./...
go build
activity-write-api
```

## Building a docker image

Run the command below to build a docker image.

```sh
docker run --rm -v "$PWD":/go/src/github.com/treeder/dockergo -w /go/src/github.com/treeder/dockergo iron/go:dev go build -o write-api
docker build -t go-social/write-api:latest .
```

## Running the image

Run the command below to run the docker image on port 8090

```sh
docker run --rm -p 8090:8080 go-social/write-api
```

## Cleaning artifacts

Run the commands below to remove dev artifacts

```sh
rm activity-write-api
docker rmi go-social/write-api
docker rmi iron/go:dev
docker rmi iron/go:latest
```
