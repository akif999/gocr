FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/akif999/gocr
COPY . /usr/local/go/src/github.com/akif999/gocr
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/gocr ./gocr


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/akif999/gocr/build/gocr /bin/gocr
CMD ["gocr", "up"]
