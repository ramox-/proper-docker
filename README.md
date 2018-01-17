# proper-docker
Design of truly minimalistic containers using my favourite docker image 'Scratch' (0 bytes)


main.go contains a very simple golang json-api using the gorilla/mux package
It exposes only a '/person' endpoint that allows basic operations (HTTP GET/POST/DELETE)

The aim of this was not to write an amazing api but rather to show how Docker and Golang goes hand-in-hand in creating minimal and maintainable containerized applications.


# compile into a single binary
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

# build image from Dockerfile
```
docker build -t proper-docker .
```

You should have a very minimal container, as opposed to when building from an image including a full linux distribution and adding lots of dependencies/frameworks to it.
Not including a full OS will almost never have any drawbacks, the docker image will be small as all it contains is your binary statically compiled with any dependencies.
Increased ease of maintainability as you will never have to worry about OS level patches for your source image.
