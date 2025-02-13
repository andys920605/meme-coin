.PHONY: redis-cluster mysql gen-env gen-mock test build clean

redis-cluster:
	docker run -d --name redis-cluster -e 'IP=0.0.0.0' -e CLUSTER_ONLY=true -p 7000-7005:7000-7005 grokzen/redis-cluster:7.0.10

mysql:
	docker run --name mysql-local -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=hr -p 3306:3306 -d mysql:8.4.4

gen-env:
	cp ./config/.env.default .env
	
gen-mock:
	go generate ./...

test:
	go test ./... -cover

build:
	go mod tidy
	go build -o build/api_server cmd/api/main.go
	
clean:
	rm -rf build