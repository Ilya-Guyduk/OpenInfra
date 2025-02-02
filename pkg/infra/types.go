package infra

// ResourceType определяет возможные типы ресурсов
type ResourceType string

const (
	ResourceVirtualMachine ResourceType = "virtual_machine"
	ResourceNetwork        ResourceType = "network"
)

// Resource описывает общий интерфейс ресурса
type Resource interface {
	GetName() string
	GetType() ResourceType
	GetProperties() map[string]interface{}
}

// VirtualMachine описывает параметры виртуальной машины
type VirtualMachine struct {
	Name       string   `yaml:"name"`
	Hypervisor string   `yaml:"provider"`
	CPU        int      `yaml:"cpu"`
	Memory     string   `yaml:"memory"`
	DiskSize   string   `yaml:"disk_size"`
	OS         string   `yaml:"os"`
	Network    string   `yaml:"network"`
	Actions    []string `yaml:"actions"`
}

// GetName возвращает имя ресурса
func (vm VirtualMachine) GetName() string {
	return vm.Name
}

// GetType возвращает тип ресурса
func (vm VirtualMachine) GetType() ResourceType {
	return ResourceVirtualMachine
}

// GetProperties возвращает свойства виртуальной машины
func (vm VirtualMachine) GetProperties() map[string]interface{} {
	return map[string]interface{}{
		"provider":  vm.Hypervisor,
		"cpu":       vm.CPU,
		"memory":    vm.Memory,
		"disk_size": vm.DiskSize,
		"os":        vm.OS,
		"network":   vm.Network,
	}
}

// NetworkResource описывает параметры сети
type NetworkResource struct {
	Name       string   `yaml:"name"`
	CIDR       string   `yaml:"cidr"`
	Gateway    string   `yaml:"gateway"`
	DNSServers []string `yaml:"dns_servers"`
	Actions    []string `yaml:"actions"`
}

// GetName возвращает имя ресурса
func (n NetworkResource) GetName() string {
	return n.Name
}

// GetType возвращает тип ресурса
func (n NetworkResource) GetType() ResourceType {
	return ResourceNetwork
}

// GetProperties возвращает свойства сетевого ресурса
func (network NetworkResource) GetProperties() map[string]interface{} {
	return map[string]interface{}{
		"cidr":        network.CIDR,
		"gateway":     network.Gateway,
		"dns_servers": network.DNSServers,
	}
}
