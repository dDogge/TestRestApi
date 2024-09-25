run:
	@echo "Running..."
	cd frontend && npm run build
	go run main.go

result:
	@echo "Getting results from db..."
	sqlite3 temp_db.db "SELECT * FROM persons;"

drop:
	@echo "Dropping table..."
	sqlite3 temp_db.db "DROP TABLE IF EXISTS persons"

.PHONY: run result drop
