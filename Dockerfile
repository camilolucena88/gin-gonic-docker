FROM golang:1.17 as builder

ENV CGO_ENABLED 0
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
# Compile the app WITHOUT optimization flags, allows Delve to
# provide a better debug experience. This creates an executable `server`
# and looks under `go-remote-debug-tutorial/example-app` for the Go files.
RUN go build -gcflags="all=-N -l" -o /server
# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

## Create Instance Container, we use Alpine to reduce size
FROM alpine:3.7

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat

# Expose debug port and application port
EXPOSE 40000 8080

# Set current working directory
WORKDIR /

# Copy the compiled executable to root
COPY --from=builder /server /
# Copy the delve executable to root
COPY --from=builder /go/bin/dlv /

# Run Delve on port 40000 on
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]