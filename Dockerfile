# Specifies a parent image
FROM golang:1.22
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Download and install the dependencies
RUN go get -d -v ./...
 
# Builds your app with optional configuration
RUN go build -o /build
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD [ “/build” ]