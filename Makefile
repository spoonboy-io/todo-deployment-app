run:
	go run -race ./cmd/todo/*.go

build:
	go build ./cmd/todo

release:
	@echo "Enter the release version (format vx.x.x).."; \
	read VERSION; \
	git tag -a $$VERSION -m "Releasing "$$VERSION; \
	git push origin $$VERSION