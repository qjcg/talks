OUTDIR := _build

.PHONY: build
build:
	mkdir -p $(OUTDIR)
	CGO_ENABLED=0 go build -o $(OUTDIR) ./cmd/...


clean:
	rm -rf $(OUTDIR)
