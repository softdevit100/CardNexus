FROM golang:1.22.2-alpine3.19

WORKDIR /app

# Install air for live reloading during development
RUN wget -O install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Install golangci-lint
RUN wget -O - https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2

# Set the default command to run air
CMD ["air"]
