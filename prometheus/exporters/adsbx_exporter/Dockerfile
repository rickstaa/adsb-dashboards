FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /adsbx_exporter

EXPOSE 19100

CMD [ "/adsbx_exporter" ]
