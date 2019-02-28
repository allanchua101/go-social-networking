FROM golang:alpine AS daemon-build-env
RUN mkdir /go/src/app

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/activity-write-daemon

COPY . .

RUN dep ensure && go build -o write-daemon

# final stage
FROM alpine
WORKDIR /app
COPY --from=daemon-build-env /go/src/activity-write-daemon /app/
ENTRYPOINT ./write-daemon