current_dir=$(shell pwd)
pkg=github.com/trashhalo/tinygo-import-report

build:
	docker run --rm -v $(current_dir)/tests/$(target):/go/src/$(pkg) tinygo/tinygo \
	build -o /go/src/$(pkg)/wasm.wasm -target wasm $(pkg)

build-init:
	docker run --rm -v $(current_dir)/tests/$(target):/go/src/$(pkg) tinygo/tinygo \
	build -initinterp -o /go/src/$(pkg)/wasm.wasm -target wasm $(pkg)