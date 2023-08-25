.PHONY: build test clean

build:
	go build

test:
	go test -v

ifeq ($(OS),Windows_NT)
clean:
	del /f ProjectName.exe
else
clean:
	rm -f ProjectName
endif
	
