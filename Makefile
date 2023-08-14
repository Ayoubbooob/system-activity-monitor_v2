# file that list commands that included in Jenkinsfile

build:
	go build -o myapp-binary ./main.go


unit-tests:
	go test app/...
