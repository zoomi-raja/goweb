FROM golang:1.14
WORKDIR /app

COPY go.mod . 
COPY go.sum .
# befor copy commany so it get caches otherwise it will try to fetch everytime
RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
EXPOSE 8000
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main