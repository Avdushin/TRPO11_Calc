# GUI Калькулейтор

## Как арбайтен?

Если есть [make](https://www.gnu.org/software/make/#download)

```
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
	go build -o bin/$(appname) main.go
```