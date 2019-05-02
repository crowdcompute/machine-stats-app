BINARY_NAME		=	machine-stats-app
VERSION         :=	$(shell cat ./VERSION)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o "./build/bin/$(BINARY_NAME)" main.go

run:
	./build/bin/$(BINARY_NAME)

run-docker:
	docker run -p 3000:3000 --name cc-machine-stats -d cc-machine-stats
	echo "Open: http://localhost:3000"

build-docker:
	docker build --tag cc-machine-stats ./build/

stop-docker:
	docker container stop cc-machine-stats
	docker container rm -f cc-machine-stats
