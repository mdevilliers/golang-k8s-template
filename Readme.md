# golang-k8s-template

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


## Features

I have made no choice as to golang libraries, dependancy managment tools etc. 

I prefer to think of this as one step beyond a 'hello world' application that can be built and deployed onto K8s.

## Choices

The Makefile assumes you have a working installation of Docker and Kubernetes.

For local Kubernetes development I would recommend -

- installing kubectl
- installing [minikube](https://github.com/kubernetes/minikube)

- sharing a local docker instance with minikube

```
 eval $(minikube docker-env)
```

