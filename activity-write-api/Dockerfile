FROM golang:alpine AS build-env
RUN mkdir /go/src/app

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/activity-write-api

COPY . .

RUN dep ensure && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/activity-write-api /app/
EXPOSE 8080
ENTRYPOINT ./goapp