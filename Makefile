# Compile & run app
run:
	@go run cmd/api/main.go $(args)

# Generate easyjson marshallers
json:
	@easyjson pkg/tools/response.go \
		internal/service/users/delivery/rest \
		internal/service/coffee/delivery/rest \
		internal/service/orders/delivery/rest

# Generate documentation
swagger:
	@docker run --rm -v .:/source -w /source quay.io/goswagger/swagger:0.30.5 generate spec -m -o ./docs/swagger.json

# Generate custom errors
errors:
	@go run cmd/errgen/errgen.go