FROM golang:latest
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o webserver
ENTRYPOINT [ "/app/webserver" ]
