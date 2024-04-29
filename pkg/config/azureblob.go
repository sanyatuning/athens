package config

// AzureBlobConfig specifies the properties required to use Azure as the storage backend.
type AzureBlobConfig struct {
	AccountName   string `envconfig:"ATHENS_AZURE_ACCOUNT_NAME"   validate:"required"`
	AccountKey    string `envconfig:"ATHENS_AZURE_ACCOUNT_KEY"    validate:"required"`
	ContainerName string `envconfig:"ATHENS_AZURE_CONTAINER_NAME" validate:"required"`
}
