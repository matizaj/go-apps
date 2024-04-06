build:
	@echo building e-com..
	@go build -o bin/e-com cmd/main.go
	@echo DONE!
test:
	@echo testing..
	@go test -v ./...
	@echo DONE!
run:
	@echo running..
	@./bin/e-com
path:
	@echo path
	@-dir cmd\migrate
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@go run cmd/migrate/main.go up
migrate-down:
	@go run cmd/migrate/main.go down