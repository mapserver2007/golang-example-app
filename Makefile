PROJECT_ROOT=$(GOPATH)/src/github.com/mapserver2007/golang-example-app

setup-init:
	brew install protobuf

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
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
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
	protoc -I ./grpc-web/proto ./grpc-web/proto/*.proto \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:./grpc-web/gen/go \
		--go_out=plugins=grpc,paths=source_relative:./grpc-web/gen/go
run-grpc-server:
	go run grpc-web/server/grpc/main.go
	go run grpc-web/server/gateway/main.go
