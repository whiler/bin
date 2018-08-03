PKGS := $(shell go list ./... | grep -v vendor | grep -v mock | grep ^bin)

default:test

test:
ifndef COVERALLS_TOKEN
	sed -i '' -e 's/github.com\/whiler\/bin/bin/' example*_test.go
	go test -v -covermode=count .
	sed -i '' -e 's/"bin"/"github.com\/whiler\/bin"/' example*_test.go
else
	go test -v -covermode=count .
endif

test-coverage:
	rm -f *.cover.out coverage.out
	rm -f coverage.html ut_coverage_report.html

ifndef COVERALLS_TOKEN
	sed -i '' -e 's/github.com\/whiler\/bin/bin/' example*_test.go
	$(foreach i, $(PKGS), go test -covermode=count -coverprofile=./`basename ${i}`.cover.out ${i} || exit;)
	sed -i '' -e 's/"bin"/"github.com\/whiler\/bin"/' example*_test.go
else
	$(foreach i, $(PKGS), go test -covermode=count -coverprofile=./`basename ${i}`.cover.out ${i} || exit;)
endif

	echo "mode: count" > coverage.out && cat *.cover.out | grep -v mode: | sort -r | awk '{if($$1 != last) {print $$0;last=$$1}}' | grep -v "\.pb\.go:" >> coverage.out
	go tool cover -html=./coverage.out -o=./coverage.html
	gocov convert coverage.out | gocov-html > ut_coverage_report.html
	rm *.cover.out coverage.out
