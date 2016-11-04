# golang-k8s-template

[![CircleCI](https://circleci.com/gh/mdevilliers/golang-k8s-template.svg?style=svg)](https://circleci.com/gh/mdevilliers/golang-k8s-template)

This project exists to provide a template to kick-start a golang project that will 

- build a minimal Docker container
- deploy to Kubernetes


## Usage

```
$ pip install --user cookiecutter
$ cookiecutter gh:mdevilliers/golang-k8s-template
```

Parameters

Parameter              | Example          | Notes
-----------------------|------------------|-----------------------------------
`project_name`         | awesome-service  | The name of your service
`docker_repository`    | mdevilliers      | The name of the docker repository
`github_account`       | mdevilliers      | Your github repository or user name


## Choices

I have made no choice as to golang libraries, dependancy managment tools etc.

This project generates a 'hello world' application with the configuration required to be built and deployed into K8s.


## FAQ

Q - I've made my project - what next?

A - Run `make help` to discover the available targets

The Makefile assumes you have a working installation of Docker and Kubernetes.

Q Anything else?

A - For local Kubernetes development I would recommend -

- installing kubectl
- installing [minikube](https://github.com/kubernetes/minikube)

- sharing a local docker instance with minikube

```
 eval $(minikube docker-env)
```

## Thanks

[jml](https://github.com/jml)

