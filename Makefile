##################################
#######       Setup       ########
##################################
.PHONY: setup

setup:
	@go mod download

start:
	@sudo ./mac-network-switch toggleAirport --switch start

stop:
	@ sudo ./mac-network-switch toggleAirport --switch stop

##################################
#######        Tool       ########
##################################
.PHONY: fmt lint clean
fmt:
	@golangci-lint run --fix

lint:
	@golangci-lint run ./...

clean:
	@git clean -fdx ${COVERAGE_FILE}
