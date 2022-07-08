OUTDIR := _build

.PHONY: build
build:
	mkdir -p $(OUTDIR)
	CGO_ENABLED=0 go build -o $(OUTDIR) ./cmd/...

.PHONY: slides
slides: main.pdf

main.pdf: main.tex
	context $<

.PHONY: clean
clean:
	rm -rf $(OUTDIR) main.pdf
