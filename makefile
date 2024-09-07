#DO NOT use startdb and stopdb if you use mariadb outside this shit project


DB_NAME := temp_db
DB_USER := root
DB_PASSWORD :=

.PHONY: startdb
startdb:
	@echo "Starting MariaDB..."
	sudo systemctl start mariadb
	@echo "Creating database $(DB_NAME) and table persons..."
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -e "USE $(DB_NAME); CREATE TABLE IF NOT EXISTS persons (id INT AUTO_INCREMENT PRIMARY KEY, firstname VARCHAR(100), lastname VARCHAR(100));"

.PHONY: run
run:
	@echo "Running..."
	go run main.go

.PHONY: resultdb
resultdb:
	@echo "Getting results from db..."
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -e "USE $(DB_NAME); SELECT * FROM persons;"

.PHONY: stopdb
stopdb:
	@echo "Dropping database $(DB_NAME)..."
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -e "DROP DATABASE IF EXISTS $(DB_NAME);"
	@echo "Stopping MariaDB..."
	sudo systemctl stop mariadb
