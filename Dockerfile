FROM golang:latest

WORKDIR /bonds

RUN go version

COPY server ./

RUN ls -la
RUN pwd

RUN go mod download
RUN go build -o bonds main.go
EXPOSE 8080
CMD ["./bonds"]

#DOCKER_BUILDKIT=0 docker build -t bonds .
#-t name tag
#docker run -p 8080:8080 bonds