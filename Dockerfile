# Deploy an Echo server with TLS connection and self-signed certificates.
# https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
FROM golang:alpine AS builder

# Set image characteristics => Cross-compilation can take place by modifying these parameters
# and rescuing the produced files if a volume is set.
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire code into the container
COPY . .

# Buid the application
RUN go build -o goecho-setup ./cmd/goecho-setup/ 

# Move to /dist directory as the place for the final binary
WORKDIR /dist

# COPY the binary and the necessary files.
# This recipe adds mender-artifact as a tool to interact and as an example of how extra
# files can be downloaded and included into the final image
RUN cp /build/goecho-setup .
RUN cp /build/internal/server/localhost.* .
RUN wget https://d1b0l86ne08fsf.cloudfront.net/mender-artifact/3.3.0/linux/mender-artifact -O ./mender-artifact

# Scratch is an empty image. There is no OS loaded.
FROM scratch

# This endpoints respect the endpoints of the goecho application in the source code
COPY --from=builder /dist/goecho-setup /cmd/goecho-setup/
COPY --from=builder /dist/mender-artifact /cmd/mender-artifact/
COPY --from=builder /dist/localhost.* /internal/server/

# Export the port/ This includes re-direction from 80 to 443
EXPOSE 80 443

ENTRYPOINT ["/cmd/goecho-setup/goecho-setup"]