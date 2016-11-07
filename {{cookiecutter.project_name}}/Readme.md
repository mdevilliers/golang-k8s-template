# {{ cookiecutter.project_name }}

TODO: add a description for your project


## Develop

To list makefile targets

```
make help
```

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

To deploy to Kubernetes -

```
make image
make deploy
```

Once deployed your service is accessible via a [NodePort](http://kubernetes.io/docs/user-guide/services/#type-nodeport).

To query the port run -

```
kubectl describe service {{ cookiecutter.project_name }}-svc-v1

```

Note that the image points to the `latest` tag for developing locally.

Please remember that 'latest is not a version' and amend for your production deploy accordingly.

