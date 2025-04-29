#
# Itachi (c) 2022-2025 by Michael Kondrashin (mkondrashin@gmail.com)
# Copyright under MIT Lincese. Please see LICENSE file for details
#
# Makefile
#

GOOS ?= windows
GOARCH ?= amd64
LDFLAGS := -ldflags "-s -w"
GMW_DIR := pkg/generate/gmw
SAMPLES := dropper encryptor spyware downloader antiav autorun novirus

EXE := $(patsubst %,$(GMW_DIR)/%.exe,$(SAMPLES))

PLATFORMS := linux-amd64 darwin-amd64 darwin-arm64 windows-amd64
TARGETS := $(foreach platform,$(PLATFORMS),itachi_$(platform)) $(foreach platform,$(PLATFORMS),witachi_$(platform))

all: $(TARGETS)

# Pattern rule for itachi binaries
itachi_%: main.go $(EXE)
	$(eval OS := $(word 1,$(subst -, ,$*)))
	$(eval ARCH := $(word 2,$(subst -, ,$*)))
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $@$(if $(filter windows,$(OS)),.exe,)

# Pattern rule for witachi binaries
witachi_%: main.go $(EXE)
	$(eval OS := $(word 1,$(subst -, ,$*)))
	$(eval ARCH := $(word 2,$(subst -, ,$*)))
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $@$(if $(filter windows,$(OS)),.exe,) ./cmd/witachi

# Pattern rule for building Windows executables
$(GMW_DIR)/%.exe: gmw/%/main.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $@ ./gmw/$*/*.go

gmw/antiav/AvList.txt.gz: gmw/antiav/AvList.txt
	gzip -c $< > $@

$(GMW_DIR)/antiav.exe: gmw/antiav/main.go gmw/antiav/AvList.txt.gz

.PHONY: clean
clean:
	rm -f $(TARGETS) $(EXE) gmw/antiav/AvList.txt.gz
