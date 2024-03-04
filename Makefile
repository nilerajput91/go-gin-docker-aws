run:
	go run main.go

## Install dependencies
deps:
	go mod tidy
	go mod vendor

compose:
	docker-compose up -d

test:
	hey -z 2m -q 20 -m GET -H "x-api-key:8113" http://localhost:8113/api/v1/startups/all

## Build the project (be sure set the enviroment variable)
build:
	@docker-compose -f docker-compose.yaml build 

start-local: 
	@docker-compose -f docker-compose.yaml up 
	
stop-local:
	@docker-compose -f docker-compose.yaml down 
