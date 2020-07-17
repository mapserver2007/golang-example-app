PROJECT_ROOT=$(GOPATH)/src/github.com/mapserver2007/golang-example-app

setup:
	go mod tidy
	# dev tools
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	# framework
	go get -u github.com/labstack/echo/...
	# db
	go get -u gopkg.in/gorp.v1
	go get -u github.com/go-sql-driver/mysql
	# grpc
	brew install protobuf
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	# others
	go get -u gopkg.in/yaml.v2
	go get -u github.com/PuerkitoBio/goquery

lint:
	golangci-lint run --golint.min-confidence 1.1
fmt:
	goimports -w web
gen:
	sh web/openapi/generator.sh
	goimports -w web
	cp -r web/openapi/out/go web/openapi
	rm -rf web/openapi/out
run-server:
	go run web/main.go

proto:
	protoc --go_out=plugins=grpc:$(PROJECT_ROOT)/grpc-web/proto grpc-web/proto/user.proto