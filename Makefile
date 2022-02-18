VERSION = $(shell git branch --show-current)

.PHONY: run
run: ## run it will instance server 
	VERSION=$(VERSION) go run main.go -port=3005 -debug=true
