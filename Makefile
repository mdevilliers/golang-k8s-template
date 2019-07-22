.PHONY: check

check:
	mkdir -p ./check
	cookiecutter --no-input -o ./check .
	cd ./check/awesome-service; make image;
	docker images
