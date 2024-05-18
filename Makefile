# Default target
all: run

# Run the Go application
run:
	go run main.go -query="$(QUERY)" -output="$(OUTPUT_FILE)"

# Clean up generated files
clean:
	rm -f $(OUTPUT_FILE)

.PHONY: all run clean
