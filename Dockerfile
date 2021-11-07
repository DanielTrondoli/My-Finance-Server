FROM golang:latest

WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

EXPOSE 8080

CMD ["./finance"]