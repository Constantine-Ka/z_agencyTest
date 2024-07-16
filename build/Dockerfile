FROM golang:alpine
LABEL authors="Constantine"
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o demoservice cmd/main/main.go
CMD ["/app/demoservice"]
