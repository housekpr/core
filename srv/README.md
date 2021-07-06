# Server notes

## How to start (requires linux)

    go run ./server.go

## How to build

If building on different platform, it's important to define the target operating system and cpu architecture.
For example building on MS Win or Apple MacOS for RaspberryPI device would require following variables.

    GOOS=linux GOARCH=arm go build ./server.go

## Gqlgen

The app uses [gqlgen](https://github.com/99designs/gqlgen) for the api layer.

### Init

    go mod init github.com/housekpr/core/srv
    go get github.com/99designs/gqlgen
    go get github.com/vektah/gqlparser/v2@v2.1.0
    go run github.com/99designs/gqlgen init

### schema the scheme

go run github.com/99designs/gqlgen generate


## Gpiod 

GPIO is managed over [gpiod](https://github.com/warthog618/gpiod). It's library for accessing GPIO pins/lines on Linux platforms using the GPIO character device.
