FROM golang:1.13.0-alpine

WORKDIR /go/src/zxq.co/ripple/rippleapi
COPY . .

# Set GOPATH and disable Go modules
ENV GOPATH=/go
ENV GO111MODULE=off

# Build the application
RUN CGO_ENABLED=0 go install -v ./...

FROM alpine:3.10
WORKDIR /rippleapi/
COPY --from=0 /go/bin/rippleapi ./

# Agree to License
RUN mkdir -p ~/.config && touch ~/.config/ripple_license_agreed

EXPOSE 40001

CMD ["./rippleapi"]
