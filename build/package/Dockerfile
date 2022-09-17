# Builder
FROM golang:1.17

WORKDIR /usr/src/app

COPY . .

RUN go mod download
RUN go build -o tilt-demo ./cmd/app/main.go
  
EXPOSE 8080

ENTRYPOINT ["/usr/src/app/tilt-demo"]