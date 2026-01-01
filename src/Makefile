.PHONY: build clean

build:
	mkdir -p ../build
	go build -o ../build/ .
	cp -r ../test/* ../build/

clean:
	rm -rf ../build