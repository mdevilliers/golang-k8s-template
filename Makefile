.PHONY: check

check:
	mkdir -p ${GOPATH}/src/github.com/mdevilliers
	cookiecutter --no-input -o ${GOPATH}/src/github.com/mdevilliers .
	cd ${GOPATH}/src/github.com/mdevilliers/awesome-service; git init; git add .; git commit -m "Initial commit"
	cd ${GOPATH}/src/github.com/mdevilliers/awesome-service; make build; make image
	docker images
