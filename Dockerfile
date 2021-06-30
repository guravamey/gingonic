# Golang Base Image
FROM golang
# Create app folder 
WORKDIR /go/src/gingonic
# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
# Download dependency
RUN go mod download
# Docker Build
RUN go build
# Expose the port 5000
EXPOSE 5000:5000
# Run the app binarry file 
CMD ["cd /go/src/gingonic"]
CMD ["./gingonic"]