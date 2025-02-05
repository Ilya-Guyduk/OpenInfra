package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProviderList(t *testing.T) {
	// Создаем тестовые данные
	providers := map[string]*Provider{
		"provider1": &Provider{Name: "provider1", Type: "virtualbox"},
		"provider2": &Provider{Name: "provider2", Type: "aws"},
	}

	spec := &OpenInfraSpec{
		Providers: providers,
	}

	// Тестируем метод GetProviderList
	providerList := spec.GetProviderList()

	// Проверяем, что список содержит правильные элементы
	assert.Equal(t, 2, len(providerList))
	assert.Contains(t, providerList, providers["provider1"])
	assert.Contains(t, providerList, providers["provider2"])
}

func TestGetProviderMap(t *testing.T) {
	// Создаем тестовые данные
	providers := map[string]*Provider{
		"provider1": &Provider{Name: "provider1", Type: "virtualbox"},
		"provider2": &Provider{Name: "provider2", Type: "aws"},
	}

	spec := &OpenInfraSpec{
		Providers: providers,
	}

	// Тестируем метод GetProviderMap
	providerMap := spec.GetProviderMap()

	// Проверяем, что карта содержит правильные данные
	assert.Equal(t, 2, len(providerMap))
	assert.Equal(t, providers["provider1"], providerMap["provider1"])
	assert.Equal(t, providers["provider2"], providerMap["provider2"])
}

func TestGetProviderByName(t *testing.T) {
	// Создаем тестовые данные
	providers := map[string]*Provider{
		"provider1": &Provider{Name: "provider1", Type: "virtualbox"},
		"provider2": &Provider{Name: "provider2", Type: "aws"},
	}

	spec := &OpenInfraSpec{
		Providers: providers,
	}

	// Тестируем существующего провайдера
	provider, err := spec.GetProviderByName("provider1")
	assert.NoError(t, err)
	assert.NotNil(t, provider)
	assert.Equal(t, "provider1", provider.Name)
	assert.Equal(t, "virtualbox", provider.Type)

	// Тестируем отсутствующего провайдера
	provider, err = spec.GetProviderByName("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, provider)
}
