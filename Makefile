# Variables
appname = caluleitor

# Actions
Default:
	go run main.go
start:
	./bin/$(appname)
run:
	go build -o bin/caluleitor.exe main.go
win:
	./bin/$(appname).exe
build:
	go build -o bin/$(appname).exe main.go