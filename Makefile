tw:
	tailwindcss -i ./cmd/static/input.css -o ./cmd/static/output.css -w -m

dev:
	air

test:
	go test ./...
