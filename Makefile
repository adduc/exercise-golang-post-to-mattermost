# take dir name, strip out "exercise-" prefix
NAME := $(shell basename $(PWD) | sed 's/^exercise-//')

run:
	go run main.go

build:
	go build -o bin/$(NAME) main.go