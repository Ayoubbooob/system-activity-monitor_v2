# Use an official Golang runtime as the base image
FROM  golang:latest as builder
# FROM  golang:1.21rc4-alpine3.18 as builder

# git is required to fetch go dependencies
# RUN apk add --no-cache ca-certificates git

# Copy the predefined netrc file into the location that git depends on
# COPY ./.netrc /root/.netrc
# RUN chmod 600 /root/.netrc


# Set the working directory
WORKDIR /app


COPY app/go.mod app/go.sum ./


# Run go mod to download dependencies
RUN go mod download

# Copy the entire source code into the container
# COPY . .
# this isn't good practice, you have to copy just *.go file not all folders and files 
COPY app/. ./





# To compile our application
# The result of that command will be a static application binary named system-activity-monitor

RUN CGO_ENABLED=0 GOOS=linux go build -o /system-activity-monitor

#RUN  go build -o /system-activity-monitor

# Use a smaller base image for the final runtimeDocker Compose for Prometheus + Grafana

FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder system-activity-monitor .

# Expose the desired metrics port
EXPOSE 9091

# Command to run the app
CMD ["./system-activity-monitor"]
