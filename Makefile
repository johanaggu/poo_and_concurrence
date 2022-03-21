test: 
	@go mod tidy
	@echo "Running test..."
	@go test ./...

test_logs:
	@go mod tidy
	@echo "Running test..."
	@go test ./... -v

test_report:
	@go mod tidy
	@echo "Running test..."
	@echo "Creating report..."
	@go test ./...  -coverprofile=coverage.out
	@go tool cover -func=coverage.out      

profiling:
	@go mod tidy
	@go test ./testing/unit_testing/... -cpuprofile=cpu.out
	@go tool pprof cpu.out