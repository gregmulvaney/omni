.PHONY: run build

BIN := ./bin/omni
PKG := ./cmd/omni

run:
	go run $(PKG)

dev:
	go run $(PKG) --log.dev --log.format console

build:
	go build -o $(BIN) $(PKG)



