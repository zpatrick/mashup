run:
	go run main.go

generate:
	$(MAKE) -C generator run

.PHONY: run generate
