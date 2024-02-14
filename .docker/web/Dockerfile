FROM golang:1.21.4-alpine

# air install
RUN go install github.com/cosmtrek/air@latest

RUN echo "alias air='$(go env GOPATH)/bin/air'" >> /root/.bashrc

# nodejs install
RUN apk add nodejs npm