run:
	sudo go run app/main.go

unit-tests:
	go test app/cpu/cpu_metrics_test.go
