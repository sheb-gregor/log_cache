# Compile stage
FROM golang:alpine

#ENV GOPROXY=direct
ENV CGO_ENABLED 0
ENV GO111MODULE=on
#ENV GOPRIVATE=gitlab.com

ARG CONFIG="master"

WORKDIR /service

ADD . .
RUN apk add --no-cache git libc6-compat ca-certificates bash \
 && go get github.com/derekparker/delve/cmd/dlv \
 && go build -gcflags "all=-N -l" -o /app . \
 && ./env/${CONFIG}.config.yaml /config.yaml

# Port 5000, 9102 belongs to our application, 40000 belongs to Delve
EXPOSE 5000 9102 40000

# Run delve
CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/app", "serve"]
