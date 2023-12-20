DBUSER=root
DBSECRET=helloworld
DBNAME=blog_service
DB_URL=mysql://$(DBUSER):$(DBSECRET)@localhost:3306/$(DBNAME)?sslmode=disable

# Start mysql container
mysql_install:
	docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=$(DBSECRET) mysql

# Create database
db_create:
	docker exec -it mysql-test mysql -p$(DBSECRET) -e "create database $(DBNAME)"

# Generate new migration
migration_new:
	migrate create -ext sql -dir db/migration -seq $(name)

# lint source file
lint:
	golangci-lint run

.PHONY: mysql_install db_create lint
