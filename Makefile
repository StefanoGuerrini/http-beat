GO := $(shell go version 2> /dev/null)

check:
ifndef GO
	$(error "go not found")
endif

monitor: check
	go run main.go monitor