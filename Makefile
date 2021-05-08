# ~/github.com/stnnnghm/go-design-patterns/Makefile

APP=go-design-patterns

build:
	clean
	go build -o ${APP} cmd/pk-strategy/strategy.go

build-strategy:
	${MAKE} -C strategy build

clean:
	go clean

run:
	go run -race cmd/pk-strategy/strategy.go

run-strategy:
	${MAKE} -C strategy run

test:
	go test -v -cover ./...

test-strategy:
	${MAKE} -C strategy test