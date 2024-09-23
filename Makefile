build:
ifeq ($(OS),Windows_NT)
	CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build $(BUILD_FLAGS) -o build/main.exe main.go
else
	CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build $(BUILD_FLAGS) -o build/main main.go
endif

genabi:
	./scripts/abigen --abi ./contracts/stakehub.abi --pkg contracts --out ./contracts/stakehub.go --type StakeHub
	./scripts/abigen --abi ./contracts/stakecredit.abi --pkg contracts --out ./contracts/stakecredit.go --type StakeCredit

.PHONY: build
