.PHONY: check

check:
	cookiecutter --no-input -o ../ .
	cd ../awesome-service; git init; git add .; git commit -m "Initial commit"
	cd ../awesome-service; make build; make image
	# sudo rm -rf ../awesome-service

