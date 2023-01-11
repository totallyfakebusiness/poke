FROM golang:1.18-alpine

LABEL org.opencontainers.image.authors="Patrick Easters"

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /serve

EXPOSE 3000

USER root

CMD [ "/serve" ]