GO := $(shell go version 2> /dev/null)

check:
ifndef GO
	$(error "go not found")
endif

install: check
	go install
	cp .http-beat.yaml ~

monitor: check
	go run main.go monitor