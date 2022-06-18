FROM golang:latest

WORKDIR /bonds

RUN go version

COPY sendMail ./

RUN ls -la
RUN pwd

RUN go mod download
RUN go build -o sendMail .
EXPOSE 8080
CMD ["./sendMail"]
