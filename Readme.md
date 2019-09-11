# golang-k8s-template

[![CircleCI](https://circleci.com/gh/mdevilliers/golang-k8s-template.svg?style=svg)](https://circleci.com/gh/mdevilliers/golang-k8s-template)

This project exists to provide a template to kick-start a golang project that will 

- build a minimal Docker container
- deploy to a Kubernetes instance

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
`golang_version`       | 1.13             | Golang version

## Thanks

[jml](https://github.com/jml)

