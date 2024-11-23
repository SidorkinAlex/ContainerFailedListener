.PHONY: build
build:
	go build -v ./cmd/CyclicCommandCheckerAndExecutive/CyclicCommandCheckerAndExecutive.go
	$(eval NEW_VER:=$(shell cat version | cut -d '_' -f 2 ))
	mkdir build
	mv CyclicCommandCheckerAndExecutive ./build/CyclicCommandCheckerAndExecutive.$(NEW_VER)