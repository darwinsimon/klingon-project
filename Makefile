build:
	@echo ">> Building..."
	@-rm -f char.txt
	@go get
	@go build
	@echo ">> Finished"