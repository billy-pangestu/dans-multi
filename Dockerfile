# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Set git config credential
ARG USER_GITLAB
ENV USER_GITLAB=$USER_GITLAB
ARG TOKEN_GITLAB
ENV TOKEN_GITLAB=$TOKEN_GITLAB

# Run credential here
RUN git config --global url."https://${USER_GITLAB}:${TOKEN_GITLAB}@gitlab.mncbank.co.id".insteadOf "https://gitlab.mncbank.co.id"

# Copy go mod and sum files
ENV GOSUMDB=off
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add tzdata
ENV TZ Asia/Jakarta
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"] 
