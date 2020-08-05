FROM golang:1.14
WORKDIR /app
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
EXPOSE 8000
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main