.PHONY: check

check:
	mkdir -p ./check
	cookiecutter --no-input -o ./check .
	cd ./check/awesome_service; make image; make lint;
	docker images
