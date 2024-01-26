FROM golang:1.21.6-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go *.json ./
RUN go build -o /golang-simple-api
EXPOSE 8080
CMD ["/golang-simple-api"]