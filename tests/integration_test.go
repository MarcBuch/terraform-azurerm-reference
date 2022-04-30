package test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type ResourceGroup struct {
	Name string `yaml:"name"`
}

func TestResourceGroup(t *testing.T) {
	configureEnv()
	rgName := readYAML("/config/input.yaml").Name

	opts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "/tf",
		BackendConfig: map[string]interface{}{
			"storage_account_name": os.Getenv("STORAGE_ACCOUNT_NAME"),
			"access_key":           os.Getenv("STORAGE_ACCOUNT_KEY"),
			"container_name":       os.Getenv("CONTAINER_NAME"),
			"key":                  os.Getenv("STATE_FILE_NAME"),
		},
	})

	t.Run("Resource group does not exist", func(t *testing.T) {
		result, _ := azure.ResourceGroupExistsE(rgName, os.Getenv("ARM_SUBSCRIPTION_ID"))

		assert.Equal(t, false, result)
	})

	terraform.InitAndApply(t, opts)
	defer terraform.Destroy(t, opts)

	t.Run("Resource group exists", func(t *testing.T) {
		assert.Equal(t, true, azure.ResourceGroupExists(t, rgName, os.Getenv("ARM_SUBSCRIPTION_ID")))
	})
}

func readYAML(filepath string) ResourceGroup {
	resourceGroup := ResourceGroup{}

	yfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(yfile), &resourceGroup)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return resourceGroup
}

func setAzureCliEnv(azureCliVarName string, armVarName string) {
	_, ok := os.LookupEnv(azureCliVarName)

	if !ok {
		armTenantId := os.Getenv(armVarName)
		os.Setenv(azureCliVarName, armTenantId)
	}
}

func configureEnv() {
	setAzureCliEnv("AZURE_TENANT_ID", "ARM_TENANT_ID")
	setAzureCliEnv("AZURE_CLIENT_ID", "ARM_CLIENT_ID")
	setAzureCliEnv("AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID")
	setAzureCliEnv("AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET")
}
