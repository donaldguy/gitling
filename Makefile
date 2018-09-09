dev:
	cd frontend/ && parcel index.html &
	ibazel run "//server" -- $(PWD)


test:
	bazel test //...


.PHONY: frontend-dev test
