# CI/CD Setup

## Building the builder container locally
    
The builder container is based on debian testing slim and includes golang and gcc-arm-linux-gnueabi in order to enable CGO cross-compiling for ARM (required to use sqlite3).

    cd $GOPATH/src/github.com/housekpr/core/cicd
    docker build -t housekpr/builder .
    docker run -v $GOPATH/src/github.com/housekpr:/go/src/github.com/housekpr --name housekpr-builder housekpr/builder

## Running the build container 

In order to save build time and enable faster local development, the base image of the container provides the required debian packages to run the build and is expecting a mounted project under the default gopath