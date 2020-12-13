build: frontend
	go build -o build/ocluso main.go 

build_dirs:
	mkdir -p build/frontend

frontend: build_dirs
	echo "Hello World from frontend!" > build/frontend/index.html

clean:
	rm -r build

run: build
	cd build && ./ocluso