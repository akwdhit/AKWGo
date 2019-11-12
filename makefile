#!/bin/bash

export BUILD_OPTS=-v
export CMD_FOLDER=cmd/AKWGo
export MAIN_FILES=${CMD_FOLDER}/*.go
export NOW=$(shell date --rfc-3339=ns)
export PKGS=$(shell go list ./... | grep -v vendor/)
export TEST_OPTS=-cover -race

test:
	@echo "${NOW} - TESTING STARTED"
	@go test ${TEST_OPTS} ${PKGS} | tee test.out
	@echo "${NOW} - TESTING FINISHED"

vet:
	@echo "${NOW} - GO VET STARTED"
	@go vet ${PKGS}
	@echo "${NOW} - GO VET FINISHED"

build:
	@make test
	@echo "" \
	
	@-make vet
	@echo "" \
	
	@echo "${NOW} - BUILD STARTED"
	@go build ${BUILD_OPTS} ${MAIN_FILES}
	@echo "${NOW} - BUILD FINISHED"