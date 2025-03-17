# Variables
COVERAGE_OUT = coverage.out
FILTERED_COVERAGE_OUT = filtered_coverage.out

# This command will run the tests for the project
.PHONY: tests
tests:
	@echo "Running tests..."
	go test -v ./...

# This command will generate a coverage report for the project, excluding files with "mock" in the name
.PHONY: coverage
coverage:
	@echo "Generating coverage report..."
	go test -cover -coverprofile=$(COVERAGE_OUT) ./...
	# Filtrar solo las líneas que deben ser procesadas
	awk '!/mock/ && /:/' $(COVERAGE_OUT) > $(FILTERED_COVERAGE_OUT)

# This command will display the coverage report in an HTML file
.PHONY: coverage-html
coverage-html: coverage
	@echo "Displaying coverage report in HTML format..."
	go tool cover -html=$(FILTERED_COVERAGE_OUT)

# This command will display the coverage report in the terminal
.PHONY: coverage-total
coverage-total: coverage
	@echo "Displaying total coverage in the terminal..."
	go tool cover -func=$(COVERAGE_OUT)

# This command will run the linter for the project
.PHONY: linter
linter:
	@echo "Running linters..."
	golangci-lint run --config .code_quality/.golangci.yml

# This Command will to start Colima with low configuration
.PHONY: start-colima-low
start-colima-low:
	@echo "Iniciando Colima con recursos limitados (CPU: 1, Memoria: 1G)..."
	colima start --cpu 1 --memory 1G

# This command will to start Colima with high configuration
.PHONY: start-colima-high
start-colima-high:
	@echo "Iniciando Colima con recursos optimizados (CPU: 4, Memoria: 4G)..."
	colima start --cpu 4 --memory 4G

# This command will to run testing with Fury
.PHONY: fury-test
fury-test:
	@echo "Ejecutando las pruebas con Fury..."
	fury test

# This command will to stop Colima
.PHONY: stop-colima
stop-colima:
	@echo "Deteniendo Colima..."
	colima stop

# This command can execute the complete workflow. 
.PHONY: test-flow
test-flow: start-colima-low fury-test stop-colima start-colima-high fury-test stop-colima

# This command will to run the pre-commit hooks for the project
.PHONY: pre-commit
pre-commit:
	@echo "Ejecutando pre-commit..."
	pre-commit run --all-files

# This command will to install the pre-commit hooks for the project
.PHONY: pre-commit-upd
pre-commit-upd:
	@echo "Instalando pre-commit..."
	pre-commit clean && pre-commit autoupdate
