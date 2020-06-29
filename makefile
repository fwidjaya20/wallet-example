migration:
	@touch ./internal/databases/$(version)_$(filename).up.sql
	@touch ./internal/databases/$(version)_$(filename).down.sql

domain:
	@mkdir -p ./internal/domains/$(name)/repositories;
	@mkdir -p ./internal/domains/$(name)/endpoints;
	@mkdir -p ./internal/domains/$(name)/transports/http;
	@mkdir -p ./internal/domains/$(name)/models;
	@mkdir -p ./internal/domains/$(name)/values;
	@mkdir -p ./internal/domains/$(name)/entities;
	@touch ./internal/domains/$(name)/repositories/interface.go
	@touch ./internal/domains/$(name)/repositories/postgres.go
	@touch ./internal/domains/$(name)/usecase.go
	@touch ./internal/domains/$(name)/service.go
	@touch ./internal/domains/$(name)/endpoints/endpoint.go

run_dev:
	@go run ./cmd/app