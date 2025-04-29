oapi-gen:
	echo "Generating OpenAPI code..."
	go tool oapi-codegen -generate types,fiber -o ./oapi/openapi.gen.go -package=api ./openapi_spec.yaml
	@echo "OpenAPI code generated successfully."

jet:
	echo "Generating SQL models..."
	jet -source=mysql -dsn="tatar:tatar@tcp(localhost:3306)/mydb" -path=./.gen

full: oapi-gen jet
	echo "All code generation tasks completed successfully."
	go mod tidy
	echo "Go modules tidied up."
	go run main.go
	echo "Running the application..."