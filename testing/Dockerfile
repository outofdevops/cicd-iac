#STAGE 1 CA Certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates

#STAGE 2 Dependencies
FROM hashicorp/terraform:1.1.4 as tf

ENV TF_LOG=DEBUG
RUN git clone -b v3.1.0 --depth 1 \
  https://github.com/terraform-google-modules/terraform-google-cloud-storage \
  /tf/modules/terraform-google-cloud-storage
COPY ./providers.tf /tf/providers.tf
RUN mkdir /mirrors /empty_dir
RUN cd /tf && terraform providers mirror /mirrors
COPY ./modules/wrapper /tf/modules/wrapper
COPY ./main.tf /tf/main.tf

#STAGE 3 Packaging
FROM scratch
ENV PATH=/bin
ENV TF_CLI_CONFIG_FILE=/terraform.rc
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=tf /tf /tf
COPY --from=tf /mirrors /mirrors
COPY --from=tf /empty_dir /tmp
COPY terraform.rc /terraform.rc
COPY --from=tf /bin/terraform /bin/terraform
WORKDIR /working_dir
