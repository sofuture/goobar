
GO_OPTIONS ?=
ifeq ($(VERBOSE), 1)
GO_OPTIONS += -v
endif

SRC_DIR := src
BIN_DIR := bin
GOOBAR_BIN := $(BIN_DIR)/goobar

.PHONY: all clean test

all: $(GOOBAR_BIN)

$(GOOBAR_BIN): $(BIN_DIR)
	@(cd $(SRC_DIR); go build $(GO_OPTIONS) -o ../$@)
	@echo $(GOOBAR_BIN) is created.

$(BIN_DIR):
	@mkdir -p $@

clean:
	@rm -rf $(dir $(GOOBAR_BIN))
ifeq ($(GOPATH), $(BUILD_DIR))
	@rm -rf $(BUILD_DIR)
else ifneq ($(GOOBAR_DIR), $(realpath $(GOOBAR_DIR)))
	@rm -f $(GOOBAR_DIR)
endif

test: all
	@(cd $(SRC_DIR); go test $(GO_OPTIONS))
