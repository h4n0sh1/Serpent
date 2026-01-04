
.PHONY: build clean

build:
	mkdir -p build
	cd src && go build -o ../build/serpent .
	cd src && GOOS=windows GOARCH=amd64 go build -o ../build/serpent.exe .
	cp -r test/* build/

clean:
	rm -rf build
