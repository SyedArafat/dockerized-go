ARG GO_VERSION=1.21
ARG GO_IMAGE_TYPE=alpine
ARG MAINTAINER="Sayed Yeamin Arafat <sayed.yeamin@brainstation-23.com>"

FROM golang:${GO_VERSION}-${GO_IMAGE_TYPE} as app-base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ=asd

#ENV TZ ${TIMEZONE}

LABEL org.opencontainers.image.authors=${MAINTAINER}


#RUN apk update && \
#    apk add nano curl

WORKDIR /usr/src/app

# COPY go.mod, go.sum and download the dependencies
COPY ./codes/go.* ./
RUN go mod download && go mod verify

# COPY All things inside the project and build
COPY ./codes .
RUN go build -o /goapp .

#FROM alpine:latest
#COPY --from=app-base /goapp /goapp

# Run the executable
EXPOSE 8080
#ENTRYPOINT [ "/goapp" ]



