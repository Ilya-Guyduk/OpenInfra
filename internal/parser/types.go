package parser

// OpenInfraSpec описывает структуру корневого документа OpenInfra
type OpenInfraSpec struct {
	Version      string               `yaml:"openinfra"`
	Info         Info                 `yaml:"info"`
	Providers    map[string]Provider  `yaml:"providers"`
	Resources    []ResourceDefinition `yaml:"components"`
	Dependencies []Dependency         `yaml:"dependencies"`
}

// Info содержит общую информацию о спецификации
type Info struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Contact     struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"contact"`
	License struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"license"`
}

type Provider struct {
	Name         string       `yaml:"name"`
	Type         string       `yaml:"type"`
	Connection   Connection   `yaml:"connection"`
	Capabilities []Capability `yaml:"capabilities"`
}

type Connection struct {
	Protocol       string         `yaml:"protocol"`
	Host           string         `yaml:"host,omitempty"`
	Port           int            `yaml:"port,omitempty"`
	Endpoint       string         `yaml:"endpoint,omitempty"`
	Authentication Authentication `yaml:"authentication"`
}

type Authentication struct {
	Method   string `yaml:"method"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	APIKey   string `yaml:"api_key,omitempty"`
}

type Capability struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Method      string      `yaml:"method"`
	Endpoint    string      `yaml:"endpoint"`
	Parameters  []Parameter `yaml:"parameters"`
}

type Parameter struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
}

// Component — описание компонента инфраструктуры
type Component struct {
	Type       string                 `yaml:"type"`
	Name       string                 `yaml:"name"`
	Provider   string                 `yaml:"provider"`
	Properties map[string]interface{} `yaml:"properties"`
	Actions    []string               `yaml:"actions"`
}

// ResourceDefinition описывает конкретный ресурс
type ResourceDefinition struct {
	Type         string                 `yaml:"type"`
	Provider     string                 `yaml:"provider"`
	Name         string                 `yaml:"name"`
	Properties   map[string]interface{} `yaml:"properties"`
	Actions      []string               `yaml:"actions"`
	Dependencies []Dependency           `yaml:"dependencies"`
}

// Dependency описывает зависимости между ресурсами
type Dependency struct {
	Resource  string   `yaml:"resource"`
	DependsOn []string `yaml:"depends_on"`
}
