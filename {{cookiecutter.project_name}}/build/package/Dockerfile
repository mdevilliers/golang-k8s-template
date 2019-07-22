# Stage 1 - Binary Build
# BUILD_X args should be passed at build time as docker build args
FROM golang:{{cookiecutter.golang_version}}-alpine AS builder
ARG BIN_VERSION
ARG GIT_COMMIT
ARG GIT_SHA
ARG GIT_TAG
ARG GIT_DIRTY
ENV BIN_OUTDIR=./
ENV BIN_NAME={{cookiecutter.project_name}}
RUN apk update && apk add build-base git libressl-dev
WORKDIR /usr/src/{{cookiecutter.project_name}}
# install dependencies in separate docker layer
COPY go.mod .
COPY go.sum .
RUN go mod download
# copy application source and build
COPY ./ .
RUN make static

# Stage 2 - Final Image
# The application should be statically linked
FROM alpine:latest
RUN apk update \
	&& apk add --no-cache ca-certificates \
	&& rm -rf /var/cache/apk/* \
    && addgroup {{cookiecutter.project_name}} \
	&& adduser -D -H -G {{cookiecutter.project_name}} {{cookiecutter.project_name}}
COPY --from=builder /usr/src/{{cookiecutter.project_name}}/{{cookiecutter.project_name}} /usr/bin/{{cookiecutter.project_name}}
VOLUME ["/etc/{{cookiecutter.project_name}}"]
ENTRYPOINT ["{{cookiecutter.project_name}}"]
EXPOSE 3000
USER {{cookiecutter.project_name}}
