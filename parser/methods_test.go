package parser

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProviderByName(t *testing.T) {
	providers := map[string]Provider{
		"aws": {
			Name: "aws",
			Type: "cloud",
		},
	}
	infra := OpenInfraSpec{Providers: providers}

	provider, err := infra.GetProviderByName("aws")
	assert.NoError(t, err)
	assert.Equal(t, "aws", provider.Name)

	_, err = infra.GetProviderByName("nonexistent")
	assert.Error(t, err)
}

func TestGetProvidersByType(t *testing.T) {
	providers := map[string]Provider{
		"aws":  {Name: "aws", Type: "cloud"},
		"vbox": {Name: "vbox", Type: "virtualbox"},
	}
	infra := OpenInfraSpec{Providers: providers}

	cloudProviders := infra.GetProvidersByType("cloud")
	assert.Len(t, cloudProviders, 1)
	assert.Equal(t, "aws", cloudProviders[0].Name)
}

func TestHasProvider(t *testing.T) {
	providers := map[string]Provider{
		"aws": {Name: "aws", Type: "cloud"},
	}
	infra := OpenInfraSpec{Providers: providers}

	assert.True(t, infra.HasProvider("aws"))
	assert.False(t, infra.HasProvider("gcp"))
}

func TestExecuteCapability(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	provider := Provider{
		Name: "test-provider",
		Connection: Connection{
			Endpoint: ts.URL,
		},
		Capabilities: []Capability{
			{Name: "test-cap", Method: "GET", Endpoint: "/test"},
		},
	}

	result, err := provider.ExecuteCapability("test-cap", nil)
	assert.NoError(t, err)
	assert.Equal(t, "Success", result)
}

func TestExecuteCapabilityMissing(t *testing.T) {
	provider := Provider{
		Name:         "test-provider",
		Capabilities: []Capability{},
	}

	_, err := provider.ExecuteCapability("missing-cap", nil)
	assert.Error(t, err)
	assert.EqualError(t, err, "возможность missing-cap не найдена у провайдера test-provider")
}
