build:
	CGO_ENABLED=0 go test ./tests/... -v -c -o reference-addon-test-harness
