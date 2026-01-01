.PHONY: build clean

build:
	mkdir -p ../build
	go build -o ../build/serpent .
	GOOS=windows GOARCH=amd64 go build -o ../build/serpent.exe .
	cp -r ../test/* ../build/

clean:
	rm -rf ../build
