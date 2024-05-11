auction-service:
	@echo "Running Auction service."
	@go run ./auction-service


bidding-service:
	@echo "Running Bidding service"
	@go run ./bidding-service

run:
	docker-compose up

demo:
	echo "Make a demo request"
	curl --location 'http://localhost:8080' --header 'Content-Type: application/json' --data '{"AdPlacementId": "11111111"}' -s
