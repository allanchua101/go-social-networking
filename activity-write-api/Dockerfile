FROM golang:alpine AS build-env
RUN mkdir /go/src/app

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/app
COPY ./Gopkg.toml /go/src/app

WORKDIR /go/src/app 

RUN dep ensure && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app /app/
ENTRYPOINT ./goapp