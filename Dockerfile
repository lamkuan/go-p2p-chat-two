FROM golang

WORKDIR /app
ADD . /app/

CMD ["cd", "/app/server", && go run server.go]
