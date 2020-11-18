FROM golang:alpine AS build

WORKDIR /opt/gateway

COPY . .
RUN go get

RUN go install

CMD ["hivengw"]
