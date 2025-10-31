.PHONY: help backend frontend run clean

help:
	@echo "Available targets:"
	@echo "  make backend    - Run the Go backend server"
	@echo "  make frontend   - Run the React frontend dev server"
	@echo "  make run        - Run both backend and frontend (requires two terminals)"
	@echo "  make clean      - Clean build artifacts"

backend:
	@echo "Starting backend server..."
	cd backend && go run ./cmd/server

frontend:
	@echo "Starting frontend dev server..."
	cd frontend && npm run dev

clean:
	@echo "Cleaning build artifacts..."
	cd frontend && rm -rf dist node_modules
	cd backend && rm -f server *.exe

# Note: Running both requires separate terminals
# Use: make backend (in one terminal) and make frontend (in another)
run:
	@echo "To run both services, use separate terminals:"
	@echo "  Terminal 1: make backend"
	@echo "  Terminal 2: make frontend"
