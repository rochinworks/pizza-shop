default: build

build:
	docker build -t pizza-shop .

test: 
	go vet ./... &&\
	go test ./...

# starts the application in watch mode
start:
	docker-compose up -d --build
	docker-compose exec app gin -i --path=src --bin=app

start-dev:
	docker-compose docker-compose.yml up -d
	HOST="localhost" gin -i run src/*.go
