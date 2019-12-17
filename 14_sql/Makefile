db-start:
	docker-compose -f docker-compose.postgres.yml up -d

db-end:
	docker-compose -f docker-compose.postgres.yml down -v

db-use:
	docker exec -it postgres-one psql -d app_db -U app_user
