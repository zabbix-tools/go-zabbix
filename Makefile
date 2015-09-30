all: build

build:
	go build -x

test:
	go test -v

docker-run:
	docker run \
		-d \
		--name zabbix-db \
		--env="MARIADB_USER=zabbix" \
		--env="MARIADB_PASS=zabbix" \
		million12/mariadb || :
	sleep 5
	docker run \
		-d \
		--name zabbix \
		--env="DB_ADDRESS=zabbix-db" \
		--env="DB_USER=zabbix" \
		--env="DB_PASS=zabbix" \
		--link=zabbix-db:zabbix-db \
		-p 8080:80 \
		zabbix/zabbix-server-2.4 || :

docker-clean:
	docker rm -f zabbix-db zabbix

.PHONY: all build test docker-run docker-clean
