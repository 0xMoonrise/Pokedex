#Possibles flags here
all: 
	go run .

test_cache:
	go test -v ./tests/cache_test.go
