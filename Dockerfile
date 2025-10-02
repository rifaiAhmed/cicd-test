FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /belajar-cicd-pemula

EXPOSE 8081

CMD ["/belajar-cicd-pemula"]