GODEP=godep

init:
	$(GODEP) restore

build:
	go build main.go

run:
	go run main.go

test:build
	#go test -v sender/redis_test.go
	#go test -v models/message_test.go
	go test -v tunnel/redis_test.go
	#go test -v sender/mail_test.go
