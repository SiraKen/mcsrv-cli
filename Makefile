NAME = mcsrv-init
BIN := bin/$(NAME)

LDFLAGS := -w \
		   -s

.PHONY: build
build:
	@echo "Building $(NAME)"
	@go build -ldflags "$(LDFLAGS)" -o $(BIN)

.PHONY: install
install:
	@make build
	@echo "Installing $(NAME)..."
	@sudo cp $(BIN) /usr/local/bin
	@echo "Done."

.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(NAME)..."
	@sudo rm /usr/local/bin/$(NAME)
	@echo "Done."
