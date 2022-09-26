FROM golang

WORKDIR /app

COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download

COPY ./server/ ./

RUN go build -o /server

EXPOSE 9999

CMD ["/server"]