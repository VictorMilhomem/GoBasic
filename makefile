run:
	go run ./cmd/main.go

build:
	go build ./cmd/

antlr:
	java -jar .\antlr-4.13.1-complete.jar -Dlanguage=Go -o ./cmd/parser/ ./cmd/parser/Basic.g4 -visitor -no-listener