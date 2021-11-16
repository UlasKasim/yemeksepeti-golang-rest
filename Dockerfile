# create image from the official Go image
FROM golang:alpine
RUN apk add --update tzdata \
    bash wget curl git;
# Create binary directory, install glide and fresh

ENV CGO_ENABLED 0    

RUN mkdir -p $$GOPATH/bin && \
    curl https://glide.sh/get | sh && \
    go get github.com/pilu/fresh
# define work directory
ADD . /go/src/yemeksepeti-golang-rest
WORKDIR /go/src/yemeksepeti-golang-rest
# serve the app
CMD go install && fresh -c runner.conf main.go