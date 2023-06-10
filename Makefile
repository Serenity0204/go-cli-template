.PHONY: build run

build:
	go build -o build main.go

## without argument
run:
	./build

## with argument
test:
	./build test $(filter-out $@,$(MAKECMDGOALS))
%:
	@: