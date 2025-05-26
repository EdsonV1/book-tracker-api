FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /app

# Download all go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code to /app
COPY . .

# Run all test
RUN go test ./...

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o book-tracker-api .

RUN apt-get update && apt-get install -y ca-certificates sqlite3 libsqlite3-0 && rm -rf /var/lib/apt/lists/*

EXPOSE 3000

# Run it
CMD [ "/app/book-tracker-api" ]
