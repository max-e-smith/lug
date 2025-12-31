main_package_path = ./cmd
binary_package_path = ./bin
binary_name = nodd-lug
test_package_path = ./test

# ====
# TEST
# ====

.PHONY: test
test:
	echo "TODO test"

.PHONY: it
it:
	test
	echo "TODO integration test"

.PHONY: smoke
smoke:
	go run main.go get cruise AT43-02 .test/smoke/testd -mv -p=8


# ====
# BILD
# ====

.PHONY: tidy
tidy:
	echo "TODO tidy"

# clean the workspace
.PHONY: clean
clean:
	echo "TODO clean"

# build the binary
.PHONY: build
build:
	go build -o bin/nodd-lug

# build and install the nodd-lug for local use
.PHONY: install
install:
	go install


# ====
# DIST
# ====

# check dependencies for vulnerabilities
.PHONY: dep-check
dep-check:
	echo "TODO dep-check"

# build project and publish
.PHONY: publish
publish: it
	echo "TODO publish"