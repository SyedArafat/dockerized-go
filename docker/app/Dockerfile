ARG GO_VERSION=1.21
ARG GO_IMAGE_TYPE=alpine

FROM golang:${GO_VERSION}-${GO_IMAGE_TYPE} as dev

ARG MAINTAINER="Sayed Yeamin Arafat <sayed.yeamin@brainstation-23.com>"

LABEL org.opencontainers.image.authors=${MAINTAINER}

ARG TIMEZONE="Asia/Dhaka"
ARG GOOS="linux"
ARG GOARCH="amd64"
ARG CGO_ENABLED=0

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY

ENV http_proxy="${HTTP_PROXY}" \
    https_proxy="${HTTPS_PROXY}" \
    no_proxy="${NO_PROXY}"

ENV GO111MODULE=on \
    CGO_ENABLED=${CGO_ENABLED} \
    GOOS=${GOOS} \
    GOARCH=${GOARCH} \
    TZ=${TIMEZONE}

WORKDIR /usr/src/app

# COPY go.mod, go.sum and download the dependencies
COPY ./codes/go.* ./
RUN go mod download && go mod verify

# COPY All things inside the project and build
COPY ./codes .
RUN go build -o /goapp ./cmd/main/

CMD ["go","run","./cmd/main/main.go"]

FROM alpine:latest as prod
COPY --from=dev /goapp /goapp

# Run the executable
CMD [ "/goapp" ]



