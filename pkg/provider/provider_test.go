package provider_test

import (
	"context"
	"testing"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
	"github.com/stretchr/testify/assert"
)

// MockProvider – фейковая реализация для тестов
type MockProvider struct{}

func (m *MockProvider) Connect(ctx context.Context) error    { return nil }
func (m *MockProvider) Disconnect(ctx context.Context) error { return nil }
func (m *MockProvider) CreateResource(ctx context.Context, r provider.Resource) (string, error) {
	return "mock-123", nil
}
func (m *MockProvider) DeleteResource(ctx context.Context, resourceID string) error { return nil }
func (m *MockProvider) GetResource(ctx context.Context, resourceID string) (provider.Resource, error) {
	return provider.Resource{ID: resourceID, Type: "virtual_machine", Provider: "mock"}, nil
}
func (m *MockProvider) ListResources(ctx context.Context, resourceType string) ([]provider.Resource, error) {
	return []provider.Resource{{ID: "mock-123", Type: "virtual_machine", Provider: "mock"}}, nil
}
func (m *MockProvider) PerformAction(ctx context.Context, resourceID string, action string) error {
	return nil
}

func TestProviderInterface(t *testing.T) {
	p := &MockProvider{}
	ctx := context.Background()

	err := p.Connect(ctx)
	assert.NoError(t, err, "Connect() должно работать без ошибок")

	resourceID, err := p.CreateResource(ctx, provider.Resource{Type: "virtual_machine"})
	assert.NoError(t, err, "CreateResource() должно работать без ошибок")
	assert.Equal(t, "mock-123", resourceID, "ID ресурса должен совпадать")

	res, err := p.GetResource(ctx, resourceID)
	assert.NoError(t, err, "GetResource() должно работать без ошибок")
	assert.Equal(t, "mock-123", res.ID, "ID ресурса должен совпадать")

	err = p.DeleteResource(ctx, resourceID)
	assert.NoError(t, err, "DeleteResource() должно работать без ошибок")

	err = p.PerformAction(ctx, resourceID, "start")
	assert.NoError(t, err, "PerformAction() должно работать без ошибок")

	err = p.Disconnect(ctx)
	assert.NoError(t, err, "Disconnect() должно работать без ошибок")
}
