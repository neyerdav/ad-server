auction-service:
	@echo "Running Auction service."
	@go run ./auction-service


bidding-service:
	@echo "Running Bidding service"
	@go run ./bidding-service

run:
	docker-compose up

gojq:
	command -v gojq > /dev/null 2>&1 || go install github.com/itchyny/gojq/cmd/gojq@latest


demo: gojq
	echo "Getting record without filters"
	curl --location 'http://localhost:8080' --header 'Content-Type: application/json'     --data '{"AdPlacementId": "11111111"}' -s | gojq
