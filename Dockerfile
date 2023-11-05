FROM golang:1.21.3-bookworm
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 3001
CMD ["go", "run", "/app/server/server.go"]