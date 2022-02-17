# BUILD and RUN TESTS

`docker build . -t terraform_test`

```shell
docker run -it -e GOOGLE_APPLICATION_CREDENTIALS=/config/sa.json \
    -v ~/sa.json:/config/sa.json:ro \
    terraform_test /working_dir/gcs_test
```