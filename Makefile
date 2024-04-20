run:
	@templ generate
	@npm run dev 
	@go run cmd/main.go

build:
	@templ generate
	@npm run build
	@go build cmd/main.go
