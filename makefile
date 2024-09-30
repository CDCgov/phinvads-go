DB_SERVICE = db  # Name of the PostgreSQL service in docker-compose
DUMP_FILE = /app/phinvads.dump  # Name of your SQL dump file (within db container)

.PHONY: start stop logs load refresh

start:
	@echo "Starting application..."
	docker compose up -d

stop:
	@echo "Stopping application..."
	docker compose down

logs:
	docker compose logs --follow app

load:
	@echo "Inserting data from $(DUMP_FILE) into $(DB_NAME)..."
	docker compose exec -T $(DB_SERVICE) pg_restore -U $(DB_USER) -x --no-owner -d $(DB_NAME) $(DUMP_FILE)

refresh:
	@echo "Running database refresh..."
	docker compose exec -T $(DB_SERVICE) psql -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME) WITH (FORCE)"
	docker compose exec -T $(DB_SERVICE) psql -U $(DB_USER) -c 'CREATE DATABASE $(DB_NAME)'
	$(MAKE) load