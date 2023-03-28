.PHONY: buf-gen buf-lint buf-format

buf-lint: _buf-exists
	buf lint --config buf.yaml

buf-format: _buf-exists
	buf format -w --config buf.yaml

buf-gen: buf-format
	buf generate --template buf.gen.yaml --config buf.yaml

.SILENT:
_buf-exists:
ifeq ($(shell which buf),)
	@echo 'need buf. see https://docs.buf.build/installation'
	exit 1
endif
