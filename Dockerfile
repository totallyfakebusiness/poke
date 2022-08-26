FROM golang:1.18-alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /serve

EXPOSE 3000

USER root

CMD [ "/serve" ]