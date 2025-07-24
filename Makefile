all: install

install:
	@sudo bash install.sh

uninstall:
	@sudo bash uninstall.sh

.PHONY: install
