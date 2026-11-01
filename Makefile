# Compile & run app
run:
	@go run cmd/coffeeshop/main.go $(args)

# Generate easyjson marshallers
json:
	@easyjson pkg/tools/response.go \
		internal/services/users/model \
		internal/services/coffee/model \
		internal/services/orders/model

# Generate documentation
swagger:
	@docker run --rm -v .:/source -w /source quay.io/goswagger/swagger:latest generate spec -m -o ./docs/swagger.json

# Generate custom errors
errors:
	@go run cmd/errgen/errgen.go

# Generate mocks
mocks:
	go generate ./...

# Run tests
test:
	@go test ./...

# Add migration
%:
	@:
migration:
	@goose -dir migrations create $(filter-out $@,$(MAKECMDGOALS)) sql