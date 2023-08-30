# Use an official Go runtime as the base image
FROM golang

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container at /app
COPY . /app

ARG port=5000

ENV port=$port

EXPOSE $port

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]

