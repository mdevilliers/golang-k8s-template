.PHONY: check

check:
	cookiecutter --no-input .
	cd ../awesome-service; git init; git add .; git commit -m "Initial commit"
	cd ../awesome-service; make build; make image
	rm -rf ../awesome-service

