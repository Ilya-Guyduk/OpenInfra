package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ilya-Guyduk/go-openinfra/pkg/provider"
)

// AWSProvider реализует интерфейс Provider для AWS
type AWSProvider struct {
	APIEndpoint string
	APIKey      string
}

// New создает новый AWSProvider
func New(apiEndpoint string, connection_details map[string]string) *AWSProvider {
	return &AWSProvider{
		APIEndpoint: connection_details["apiEndpoint"],
		APIKey:      connection_details["apiKey"],
	}
}

// Connect – подключение к AWS API
func (a *AWSProvider) Connect(ctx context.Context) error {
	fmt.Println("Connecting to AWS at", a.APIEndpoint)
	return nil
}

// Disconnect – закрытие сессии (если требуется)
func (a *AWSProvider) Disconnect(ctx context.Context) error {
	fmt.Println("Disconnecting from AWS")
	return nil
}

// CreateResource – создание ресурса в AWS (например, EC2)
func (a *AWSProvider) CreateResource(ctx context.Context, resource provider.Resource) (string, error) {
	if resource.Type != "virtual_machine" {
		return "", errors.New("unsupported resource type for AWS")
	}
	fmt.Println("Creating EC2 instance with properties:", resource.Properties)
	return "aws-vm-456", nil
}

// DeleteResource – удаление ресурса
func (a *AWSProvider) DeleteResource(ctx context.Context, resourceID string) error {
	fmt.Println("Deleting AWS resource", resourceID)
	return nil
}

// GetResource – получение информации о ресурсе
func (a *AWSProvider) GetResource(ctx context.Context, resourceID string) (provider.Resource, error) {
	return provider.Resource{
		ID:       resourceID,
		Type:     "virtual_machine",
		Provider: "aws",
		Properties: map[string]interface{}{
			"cpu":    4,
			"memory": "8GB",
			"disk":   "100GB",
			"status": "running",
		},
	}, nil
}

// ListResources – получение списка ресурсов
func (a *AWSProvider) ListResources(ctx context.Context, resourceType string) ([]provider.Resource, error) {
	return []provider.Resource{
		{ID: "aws-vm-456", Type: "virtual_machine", Provider: "aws", Properties: map[string]interface{}{"status": "running"}},
	}, nil
}

// PerformAction – выполнение действий (старт, стоп)
func (a *AWSProvider) PerformAction(ctx context.Context, resourceID string, action string) error {
	fmt.Printf("Performing action %s on AWS resource %s\n", action, resourceID)
	return nil
}
