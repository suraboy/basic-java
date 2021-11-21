FROM golang

WORKDIR /app

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main serve-rest"