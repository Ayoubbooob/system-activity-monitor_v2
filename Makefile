build:
	'go build -o myapp-binary ./main.go'


unit-tests:
	go test ./app/...
