GO_CMD=go

setup:
	# dev tools
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	# framework
	go get -u github.com/labstack/echo/...
	# db
	go get -u github.com/go-gorp/gorp
	go get -u github.com/go-sql-driver/mysql
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