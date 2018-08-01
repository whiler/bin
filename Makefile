PKGS := $(shell go list ./... | grep -v vendor | grep -v mock | grep ^bin)

default:test

test:
	go test -v .

test-coverage:
	rm -f *.cover.out coverage.out
	rm -f coverage.html ut_coverage_report.html
	$(foreach i, $(PKGS), go test -covermode=count -coverprofile=./`basename ${i}`.cover.out ${i} || exit;)
	echo "mode: count" > coverage.out && cat *.cover.out | grep -v mode: | sort -r | awk '{if($$1 != last) {print $$0;last=$$1}}' | grep -v "\.pb\.go:" >> coverage.out
	go tool cover -html=./coverage.out -o=./coverage.html
	gocov convert coverage.out | gocov-html > ut_coverage_report.html
	rm *.cover.out coverage.out