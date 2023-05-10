# Start from a Debian based image with Go installed and a workspace (GOPATH) configured at /go.
FROM golang:1.17

# Copy the local package files to the container's workspace.
ADD . /go/src/wallet-recovery

# Set the current working directory inside the container
WORKDIR /go/src/wallet-recovery

# Build the application inside the container.
# If the application consists of multiple files, ensure the recovery file is specified.
RUN go build -o recovery .

# Run the binary program produced by `go build`
CMD ["/go/src/wallet-recovery/recovery"]