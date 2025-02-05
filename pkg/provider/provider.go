package provider

import "context"

// Provider – общий интерфейс для всех провайдеров (локальных и облачных).
type Provider interface {
	Connect(ctx context.Context) error                                          // Установить соединение с провайдером
	Disconnect(ctx context.Context) error                                       // Закрыть соединение (если требуется)
	CreateResource(ctx context.Context, resource Resource) (string, error)      // Создать ресурс
	DeleteResource(ctx context.Context, resourceID string) error                // Удалить ресурс
	GetResource(ctx context.Context, resourceID string) (Resource, error)       // Получить информацию о ресурсе
	ListResources(ctx context.Context, resourceType string) ([]Resource, error) // Список ресурсов определенного типа
	PerformAction(ctx context.Context, resourceID string, action string) error  // Выполнить действие (start, stop и т. д.)
}

// Resource – универсальная структура ресурса.
type Resource struct {
	ID         string                 // Уникальный идентификатор ресурса
	Type       string                 // Тип (virtual_machine, network и т. д.)
	Provider   string                 // Провайдер (virtualbox, aws, gcp и т. д.)
	Properties map[string]interface{} // Свойства ресурса (CPU, RAM, CIDR и т. д.)
}
