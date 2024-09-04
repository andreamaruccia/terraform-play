package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsInstance(t *testing.T) {
	// Configure Terraform options
	terraformOptions := &terraform.Options{
		TerraformDir: "../tf",
	}

	// Clean up resources with "terraform destroy" at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Get the output variable
	message := terraform.Output(t, terraformOptions, "hello_world")

	// Verify we're getting a valid instance ID
	assert.Equal(t, message, "Hello, World!")
}
