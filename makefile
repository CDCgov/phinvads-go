DB_SERVICE = db  # Name of the PostgreSQL service in docker-compose
DUMP_FILE = /app/phinvads.dump  # Name of your SQL dump file (within db container)

.PHONY: startdb stop startapp starttempl startair load refreshdb

startdb:
	@echo "Starting database..."
	docker compose up -d

stop:
	@echo "Stopping database..."
	docker compose down

startapp:
	@echo "Starting application locally..."
	$(MAKE) -j starttempl startair

starttempl:
	templ generate --watch --proxy="http://localhost:4000"

startair:
	air -c .air-with-proxy.toml

load:
	@echo "Inserting data from $(DUMP_FILE) into $(DB_NAME)..."
	docker compose exec -T $(DB_SERVICE) pg_restore -U $(DB_USER) -x --no-owner -d $(DB_NAME) $(DUMP_FILE)

refreshdb:
	@echo "Running database refresh..."
	docker compose exec -T $(DB_SERVICE) psql -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME) WITH (FORCE)"
	docker compose exec -T $(DB_SERVICE) psql -U $(DB_USER) -c 'CREATE DATABASE $(DB_NAME)'
	$(MAKE) load