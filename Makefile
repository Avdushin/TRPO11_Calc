# Variables
appname = caluleitor

# Actions
Default:
	go run main.go
start:
	./bin/$(appname)
win:
	./bin/$(appname).exe
build:
	go build -o bin/$(appname).exe main.go