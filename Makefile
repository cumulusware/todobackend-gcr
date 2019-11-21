help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check         Format, lint, vet, and test Go code"
	@echo "  cover         Show test coverage in html"
	@echo "  deploy        Deploy to IBM Cloud Foundry"
	@echo "  dev           Build and run for local development OS"
	@echo "  local         Build for local development OS"

check:
	@echo 'Formatting, linting, vetting, and testing Go code'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./... -cover

cover:
	@echo 'Test coverage in html'
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

#  Compile the project to run locally on your machine
local:
	go build -o dist/todobackend .

dev: local
	dist/todobackend

deploy:
	go mod tidy
	gcloud builds submit --tag gcr.io/todobackendgcr/todobackend-gcr
	gcloud run deploy --image gcr.io/todobackendgcr/todobackend-gcr --platform managed
