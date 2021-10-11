all: build

build:
	go build -x

test:
	go test -v

docker-run:
	docker network create zabbix-net || :
	docker run \
		-d \
		--name zabbix-db \
		--env="POSTGRES_USER=zabbix" \
		--env="POSTGRES_PASSWORD=zabbix" \
		--net=zabbix-net \
		postgres:13.4-alpine || :
	sleep 5
	docker run \
		-d \
		--name zabbix-server \
		--env="DB_SERVER_HOST=zabbix-db" \
		--env="POSTGRES_USER=zabbix" \
		--env="POSTGRES_PASSWORD=zabbix" \
		--link=zabbix-db:zabbix-db \
		--net=zabbix-net \
		zabbix/zabbix-server-pgsql:5.4.5-alpine || :
	sleep 5
	docker run \
		-d \
		--name zabbix \
		--env="ZBX_SERVER_HOST=zabbix-server" \
		--env="DB_SERVER_HOST=zabbix-db" \
		--env="POSTGRES_USER=zabbix" \
		--env="POSTGRES_PASSWORD=zabbix" \
		--link=zabbix-db:zabbix-db \
		--link=zabbix-server:zabbix-server \
		--net=zabbix-net \
		-p 8080:8080 \
		zabbix/zabbix-web-nginx-pgsql:5.4.5-alpine || :

docker-clean:
	docker rm -f zabbix-db zabbix-server zabbix
	docker network rm zabbix-net

.PHONY: all build test docker-run docker-clean
