GO_CMD=go

setup:
	# dev tools
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	# db
	go get -u github.com/go-gorp/gorp
	go get -u github.com/go-sql-driver/mysql
	# scraping
	go get -u github.com/PuerkitoBio/goquery
	# openapi
	# go get -u github.com/getkin/kin-openapi
	# go get -u github.com/go-swagger/go-swagger/cmd/swagger

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