download_lib:
	go mod download

install_go_bind:
	go get -u github.com/jteeuwen/go-bindata/...

generate_schema:
	cd gql/schema && rm bindata.go && go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

test_all:
	go test -v ./...

test_unit:
	go test -v -short ./...

up:
	docker-compose up -d

run:
	go run cmd/server.go

lint:
	golangci-lint run -v -c ./.golangci.yml