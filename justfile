# List all available recipies
default:
    just -l -u

# Run the program
run *OPTS:
    go run . {{OPTS}}

# Build the project for Linux and Windows
build:
	GOOS=linux go build -o omp_fs_api .
	GOOS=windows go build -o omp_fs_api.exe .

# Build the project for Linux and Windows and put the binaries in the bin directory
build_bin:
	GOOS=linux go build -o bin/omp_fs_api .
	GOOS=windows go build -o bin/omp_fs_api.exe .

# Run all tests
test:
    go test -v ./...
