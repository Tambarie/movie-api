
run:
	gofmt -w .
	go run movie-service.go


mock-redis-db:
	mockgen -source=internal/ports/movieService.go -destination=internal/ports/mock_redis_db.go -package=ports


mock-postgres-db:
	mockgen -source=internal/ports/movieRepository.go -destination=internal/ports/mock_postgres_db.go -package=ports