FROM golang

WORKDIR /app

COPY . .

RUN go mod download -x

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main serve-rest"