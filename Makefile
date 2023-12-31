NAME=go-mono

format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...

fumpt:
	${call colored,fumpt is running...}
	gofumpt -l -w .