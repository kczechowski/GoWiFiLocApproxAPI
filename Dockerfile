FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

## Build the Go app
#RUN go build -o app .
#
## Expose port 8080 to the outside world
#EXPOSE 8080
#
## Run the executable
#CMD ["./app"]

ENTRYPOINT CompileDaemon --build="go build -o build/app ." --command=./build/app