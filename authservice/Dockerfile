FROM golang:1.22.5 AS builder

WORKDIR /app

# Copy and download dependencies
COPY . .
RUN go mod download

COPY .env .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .

# Expose port 8080
EXPOSE 8081
EXPOSE 8075

# Command to run the executable
CMD ["./myapp"]
