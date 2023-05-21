package credhelpers

import "github.com/nextlinux/stereoscope/pkg/image"

type CredentialHelper interface {
	GetRegistryCredentials() (*image.RegistryCredentials, error)
}

type internalHelper interface {
	Get(serverURL string) (string, string, error)
}
