# STAGE 1 CA Certificates
FROM alpine:latest as certs

RUN apk --update add ca-certificates

# STAGE 2 Dependencies
FROM hashicorp/terraform:1.1.8 as tf

ENV TF_LOG=DEBUG
COPY ./providers.tf /tf/providers.tf
COPY ./versions.tf /tf/versions.tf
RUN mkdir /mirrors /empty_dir
RUN cd /tf && terraform providers mirror /mirrors
# Import modules & main after terraform providers mirror
COPY ./modules/tf-azure-resource-group /tf/modules/tf-azure-resource-group
COPY ./modules/wrapper /tf/modules/wrapper
COPY ./main.tf /tf/main.tf

# STAGE 3 Tests
FROM golang:1.18 as test

RUN mkdir /tests
COPY ./tests/go.mod /tests
COPY ./tests/go.sum /tests
RUN cd /tests && go mod download && go mod verify
COPY ./tests/integration_test.go /tests
ENV CGO_ENABLED=0
RUN cd /tests && go test -c -o integration_test

# STAGE 4 Packaging
FROM scratch

ENV PATH=/bin
ENV TF_CLI_CONFIG_FILE=/terraform.rc
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=tf /tf /tf
COPY --from=tf /mirrors /mirrors
COPY --from=tf /empty_dir /tmp
COPY terraform.rc /terraform.rc
COPY --from=tf /bin/terraform /bin/terraform
COPY --from=test /tests /tests
WORKDIR /working_dir