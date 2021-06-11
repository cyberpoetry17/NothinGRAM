FROM golang:1.15-alpine3.13
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o user-service
EXPOSE 8080
CMD ["./user-service"]