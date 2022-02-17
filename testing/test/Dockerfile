#STAGE 1 Build your test
FROM golang:1.16 as build
WORKDIR /work_dir
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY gcs_test.go .
ENV CGO_ENABLED=0
RUN go test -c -o gcs_test

#STAGE 2 Run Test
FROM terraform-gcs as test

COPY --from=build /work_dir/gcs_test /working_dir/gcs_test

