ARG GO_VERSION=1.21
ARG GO_IMAGE_TYPE=alpine
ARG MAINTAINER="Sayed Yeamin Arafat <sayed.yeamin@brainstation-23.com>"

FROM golang:${GO_VERSION}-${GO_IMAGE_TYPE} as app-base

ARG TIMEZONE
ARG GOOS
ARG GOARCH
ARG CGO_ENABLED
ARG DOCKER_BUILD_MODE

ENV GO111MODULE=on \
    CGO_ENABLED=${CGO_ENABLED} \
    GOOS=${GOOS} \
    GOARCH=${GOARCH} \
    TZ=${TIMEZONE}

LABEL org.opencontainers.image.authors=${MAINTAINER}

# Install additional dependencies if needed
#RUN apk update && \
#    apk add nano curl

WORKDIR /usr/src/app

# COPY go.mod, go.sum and download the dependencies
COPY ./codes/go.* ./
RUN go mod download && go mod verify

# COPY All things inside the project and build
COPY ./codes .
RUN go build -o /goapp .

EXPOSE 8080

# Run the executable
#ENTRYPOINT [ "/goapp" ]

# Run Go program, just like locally
ENTRYPOINT ["go","run","main.go"]


