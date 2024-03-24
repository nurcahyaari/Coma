start:
	go run github.com/cosmtrek/air

dev: generate
	go run github.com/cosmtrek/air

generate:
	go generate ./...

build: generate
	mkdir -p build && go build -o build/app

clean:
	@find . -name **fakes -delete
	@rm -rf coma/coma.cfg docs/docs.go docs/swagger.json docs/swagger.yaml docs