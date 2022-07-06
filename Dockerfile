FROM golang:latest

WORKDIR /bonds

RUN go version

COPY cmd/server ./

RUN ls -la
RUN pwd

RUN go mod download
RUN go build -o bonds .
EXPOSE 8080
CMD ["./bonds"]

#DOCKER_BUILDKIT=0 docker build -t bonds .
#DOCKER_BUILDKIT=0 docker build -f client/Dockerfile -t bonds-client .
#-t name tag
#docker run -p 8080:8080 bonds
#docker-compose up --build
