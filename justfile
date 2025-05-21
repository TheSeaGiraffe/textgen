# List all available recipies
default:
    just -l -u

# Run the program
run *OPTS:
    go run . {{OPTS}}

# Build the project for Linux and Windows
build:
	GOOS=linux go build -o textgen .
	GOOS=windows go build -o textgen.exe .

# Build the project for Linux and Windows and put the binaries in the bin directory
build_bin:
	GOOS=linux go build -o bin/textgen .
	GOOS=windows go build -o bin/textgen.exe .

# Run all tests
test:
    go test -v ./...
