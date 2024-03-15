all: clean build
clean:
	@echo "Cleaning up..."
	rm -rfv dist
	rm -rfv tmp
	rm -rfv data
	rm -fv deploydock
build: clean
	@echo "Building..."
	go build -o dist/deploydock
test: build
	@echo "Testing..."
	./dist/deploydock