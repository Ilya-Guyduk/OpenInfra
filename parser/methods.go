package parser

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (ois *OpenInfraSpec) GetProviderList() []Provider {
	var providerList []Provider
	for _, provider := range ois.Providers {
		providerList = append(providerList, provider)
	}
	return providerList
}

func (ois *OpenInfraSpec) GetProviderMap() map[string]Provider {
	return ois.Providers
}

func (ois *OpenInfraSpec) GetProviderByName(name string) (Provider, error) {
	if provider, exists := ois.Providers[name]; exists {
		return provider, nil
	} else {
		return Provider{}, fmt.Errorf("Provider not found!")
	}
}

func (ois *OpenInfraSpec) GetProvidersByType(providerType string) []Provider {
	var providers []Provider
	for _, provider := range ois.Providers {
		if provider.Type == providerType {
			providers = append(providers, provider)
		}
	}
	return providers
}

func (ois *OpenInfraSpec) HasProvider(name string) bool {
	_, exists := ois.Providers[name]
	return exists
}

func (ois *OpenInfraSpec) GetProviderCapability(name, capabilityName string) (*Capability, error) {
	if provider, exists := ois.Providers[name]; exists {
		for _, cap := range provider.Capabilities {
			if cap.Name == capabilityName {
				return &cap, nil
			}
		}
		return nil, fmt.Errorf("capability %s not found for provider %s", capabilityName, name)
	}
	return nil, fmt.Errorf("provider %s not found", name)
}

func (ois *OpenInfraSpec) ProviderCapabilityList(name string) []Capability {
	pr := ois.Providers[name]
	return pr.Capabilities
}

func (ois *OpenInfraSpec) GetAllCapabilities() []Capability {
	var capabilities []Capability
	for _, provider := range ois.Providers {
		capabilities = append(capabilities, provider.Capabilities...)
	}
	return capabilities
}

func (ois *OpenInfraSpec) GetProvidersWithCapability(capabilityName string) []Provider {
	var providers []Provider
	for _, provider := range ois.Providers {
		for _, cap := range provider.Capabilities {
			if cap.Name == capabilityName {
				providers = append(providers, provider)
				break
			}
		}
	}
	return providers
}

func (p *Provider) ExecuteCapability(name string, params map[string]interface{}) (string, error) {
	for _, capability := range p.Capabilities {
		if capability.Name == name {
			// Формируем URL
			url := p.Connection.Host
			if p.Connection.Endpoint != "" {
				url = p.Connection.Endpoint
			}
			url += capability.Endpoint

			// Подставляем параметры в URL (если есть в пути)
			for _, param := range capability.Parameters {
				if param.Required {
					value, exists := params[param.Name]
					if !exists {
						return "", fmt.Errorf("отсутствует обязательный параметр: %s", param.Name)
					}
					url = strings.Replace(url, "{"+param.Name+"}", fmt.Sprintf("%v", value), -1)
				}
			}

			// Делаем HTTP-запрос
			req, err := http.NewRequest(capability.Method, url, nil)
			if err != nil {
				return "", fmt.Errorf("ошибка при создании запроса: %w", err)
			}

			// Аутентификация
			if p.Connection.Authentication.Method == "api_key" {
				req.Header.Set("Authorization", "Bearer "+p.Connection.Authentication.APIKey)
			} else if p.Connection.Authentication.Method == "password" {
				req.SetBasicAuth(p.Connection.Authentication.Username, p.Connection.Authentication.Password)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return "", fmt.Errorf("ошибка при выполнении запроса: %w", err)
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			return string(body), nil
		}
	}
	return "", fmt.Errorf("возможность %s не найдена у провайдера %s", name, p.Name)
}
