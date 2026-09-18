tw:
	tailwindcss -i ./static/input.css -o ./static/output.css -w -m

tw-cmd:
	tailwindcss -i ./cmd/static/input.css -o ./cmd/static/output.css -w -m

dev:
	air

test:
	go test ./...
