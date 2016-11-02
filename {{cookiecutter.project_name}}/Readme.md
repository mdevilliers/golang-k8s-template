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

The application is built locally before being published to your configured Docker repository.

Images are tagged with the format {git branch name}-{git short hash}


Kubernetes
----------

The development workflow is optimised around a standard Kubernetes toolset.

- minikube

```
minikube start
make image
make deploy

```

Note that the image points to the `latest` tag for developing locally.

Please remember that 'latest is not a version' and amend for your production deploy accordingly.


Develop
=======

To list makefile targets

```
make help
```

Dependancies
------------

- ???

Preflight
---------

- ???

Tests
-----

- ???


