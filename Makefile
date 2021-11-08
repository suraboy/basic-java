GO_BINARY_NAME=go-fiber-api
VERSION=$(shell git describe --tags || git rev-parse --short HEAD || echo "unknown version")
LDFLAGS+= -X "github.com/suraboy/go-fiber-api/cmd.Version=$(VERSION)"
LDFLAGS+= -X "github.com/suraboy/go-fiber-api/cmd.GoVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"

# Environment injection.
inject-env:
ifneq ("$(wildcard .env)","")
	$(eval -include .env) \
	$(eval export $(shell sed 's/=.*//' .env))
else
 	$(warning ".env file not found.")
endif

# Always turn on go module when use `go` command.
GO := GO111MODULE=on go

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
