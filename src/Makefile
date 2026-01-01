.PHONY: build clean

build:
	mkdir -p ../build
	go build -o ../build/serpent .
	cp -r ../test/* ../build/

clean:
	rm -rf ../build