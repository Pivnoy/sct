
run:
	go run ./cmd/main/main.go

bench_test:
	cd internal
	go test -bench=. ./internal -count=1