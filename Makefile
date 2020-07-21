build:
	set GOOARCH=386
	go build -o bin/kroll-data-stripper-x86.exe --buildmode=exe main.go
	set GOOARCH=amd64
	go build -o bin/kroll-data-stripper-amd64.exe --buildmode=exe main.go
run:
	go run main.go
package:
	make build
	copy LICENSE bin\LICENSE
	copy README.md bin\README.md
