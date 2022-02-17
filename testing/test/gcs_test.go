package test

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const input_yaml = `
---
project_id: "seed-334620"
prefix: "storage"
names: ["anto","general"]
folders:
  anto: ["/documents","/private/anto"]
  general: ["/docs","/public/general"]
bucket_policy_only:
  anto: true
  general: false
force_destroy: true
lifecycle_rules:
  - action:
      type: "SetStorageClass"
      storage_class: "NEARLINE"
    condition:
      age: "10"
      matches_storage_class: "MULTI_REGIONAL,STANDARD,DURABLE_REDUCED_AVAILABILITY"
`

func writeInput(content string) {
	d1 := []byte(content)
	e := os.WriteFile("/config/input.yaml", d1, 0644)
	if e != nil {
		panic(e)
	}
}

func TestTerraformGCS(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "/tf",
		BackendConfig: map[string]interface{}{
			"prefix": "test",
			"bucket": "tf-state-outofdevops",
		},
	})

	writeInput(input_yaml)
	defer terraform.Destroy(t, terraformOptions)

	assert.Equal(t, false, bucketExists("storage-eu-anto"))
	assert.Equal(t, false, bucketExists("storage-eu-general"))
	terraform.InitAndApply(t, terraformOptions)

	assert.Equal(t, true, bucketExists("storage-eu-anto"))
	assert.Equal(t, true, bucketExists("storage-eu-general"))
}

func bucketExists(bucketName string) bool {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return false
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	_, err = bucket.Attrs(ctx)

	return err == nil
}
