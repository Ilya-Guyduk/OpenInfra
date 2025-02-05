package virtualbox

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
)

// VirtualBoxProvider реализует интерфейс Provider
type VirtualBoxProvider struct {
	Address  string
	Username string
	Password string
}

// New создает новый инстанс VirtualBoxProvider
func New(address string, connection_details map[string]string) *VirtualBoxProvider {
	return &VirtualBoxProvider{
		Address:  connection_details["address"],
		Username: connection_details["username"],
		Password: connection_details["password"],
	}
}

// Connect – устанавливает соединение с VirtualBox
func (v *VirtualBoxProvider) Connect(ctx context.Context) error {
	fmt.Println("Connecting to VirtualBox at", v.Address)
	return nil // Здесь можно добавить реальную проверку подключения
}

// Disconnect – закрывает соединение (если требуется)
func (v *VirtualBoxProvider) Disconnect(ctx context.Context) error {
	fmt.Println("Disconnecting from VirtualBox")
	return nil
}

// CreateResource – создает ресурс (например, виртуальную машину)
func (v *VirtualBoxProvider) CreateResource(ctx context.Context, resource provider.Resource) (string, error) {
	if resource.Type != "virtual_machine" {
		return "", errors.New("unsupported resource type")
	}
	fmt.Println("Creating Virtual Machine with properties:", resource.Properties)
	return "vm-123", nil
}

// DeleteResource – удаляет ресурс
func (v *VirtualBoxProvider) DeleteResource(ctx context.Context, resourceID string) error {
	fmt.Println("Deleting resource", resourceID)
	return nil
}

// GetResource – получает информацию о ресурсе
func (v *VirtualBoxProvider) GetResource(ctx context.Context, resourceID string) (provider.Resource, error) {
	return provider.Resource{
		ID:       resourceID,
		Type:     "virtual_machine",
		Provider: "virtualbox",
		Properties: map[string]interface{}{
			"cpu":    2,
			"memory": "4GB",
			"disk":   "50GB",
			"status": "running",
		},
	}, nil
}

// ListResources – список всех ресурсов определенного типа
func (v *VirtualBoxProvider) ListResources(ctx context.Context, resourceType string) ([]provider.Resource, error) {
	return []provider.Resource{
		{ID: "vm-123", Type: "virtual_machine", Provider: "virtualbox", Properties: map[string]interface{}{"status": "running"}},
	}, nil
}

// PerformAction – выполняет действие (start, stop, restart)
func (v *VirtualBoxProvider) PerformAction(ctx context.Context, resourceID string, action string) error {
	fmt.Printf("Performing action %s on resource %s\n", action, resourceID)
	return nil
}
