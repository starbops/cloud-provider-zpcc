.PHONY: build package clean

build:
	@echo "Building..."
	@mkdir -p bin/
	@env GOOS=linux GOARCH=amd64 go build -o bin/cloud-provider-zpcc-amd64 .
	@env GOOS=darwin GOARCH=arm64 go build -o bin/cloud-provider-zpcc-arm64 .

package:
	@echo "Packaging..."
	@docker buildx build --load -t starbops/cloud-provider-zpcc:latest --platform linux/amd64 .

clean:
	@echo "Cleaning..."
	@\rm -rf bin/
