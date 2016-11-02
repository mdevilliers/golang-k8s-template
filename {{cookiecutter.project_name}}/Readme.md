{{ cookiecutter.project_name }}
-------------------------------

TODO: add a description for your project


Build
=====

Golang
-----

To build the application locally -

```
make build
```

The built image is outputted to the /bin folder

Docker
------

To build the Docker image -

```
make image
```

The application is built in the `build image` e.g. golang:{some version} before being published to your configured Docker repository.

Images are tagged with the format {git branch name}-{git short hash}


Kubernetes
----------

The development workflow is optimised around the standard Kubernetes toolset

- minikube
- helm

```
minikube start


```



Develop
=======

Dependancies
------------

- glide

Preflight
---------

- ???

Tests
-----

- ???


