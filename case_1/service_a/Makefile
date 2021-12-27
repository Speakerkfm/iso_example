CI_PROJECT_DIR?=$(shell pwd)
PROTOTOOL?=$(CI_PROJECT_DIR)/bin/prototool

.PHONY: prototool-install
prototool-install:
	@echo "это займёт время, не торопись"
	@echo "https://github.com/uber/prototool#installation"
	curl -sSL https://github.com/uber/prototool/releases/download/v1.9.0/prototool-$$(uname -s)-$$(uname -m) -o $(PROTOTOOL)
	chmod +x $(PROTOTOOL)
	$(PROTOTOOL) version