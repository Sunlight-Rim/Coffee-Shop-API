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
	@docker run --rm -v .:/source -w /source quay.io/goswagger/swagger:0.30.5 generate spec -m -o ./docs/swagger.json

# Generate custom errors
errors:
	@go run cmd/errgen/errgen.go