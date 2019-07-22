.PHONY: check

check:
	mkdir -p ./check
	cookiecutter --no-input -o ./check .
	cd ./check/awesome-service; make build; make image; make test;
	docker images
