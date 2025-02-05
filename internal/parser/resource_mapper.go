package parser

import "github.com/Ilya-Guyduk/go-openinfra/pkg/infra"

// ConvertResources преобразует ресурсы в конкретные структуры
func ConvertResources(resources []ResourceDefinition) []ResourceDefinition {
	for i, resource := range resources {
		switch resource.Type {
		case infra.ResourceVirtualMachine:
			// Преобразуем в VirtualMachine
			var vm infra.VirtualMachine
			vm.Name = resource.Name
			vm.Hypervisor = resource.Provider
			vm.CPU = convertToInt(resource.Properties["cpu"])
			vm.Memory = resource.Properties["memory"].(string)
			vm.DiskSize = resource.Properties["disk_size"].(string)
			vm.OS = resource.Properties["os"].(string)
			vm.Network = resource.Properties["network"].(string)
			vm.Actions = resource.Actions

			// Преобразуем обратно в ResourceDefinition
			resources[i] = ResourceDefinition{
				Type:       infra.ResourceVirtualMachine,
				Name:       vm.Name,
				Properties: map[string]interface{}{"hypervisor": vm.Hypervisor, "cpu": vm.CPU, "memory": vm.Memory, "disk_size": vm.DiskSize, "os": vm.OS, "network": vm.Network},
				Actions:    vm.Actions,
			}

		case infra.ResourceNetwork:
			// Преобразуем в NetworkResource
			var network infra.NetworkResource
			network.Name = resource.Name
			network.CIDR = resource.Properties["cidr"].(string)
			network.Gateway = resource.Properties["gateway"].(string)

			// Обрабатываем массив DNS-серверов
			network.DNSServers = make([]string, 0)
			if dns, ok := resource.Properties["dns_servers"].([]interface{}); ok {
				for _, d := range dns {
					network.DNSServers = append(network.DNSServers, d.(string))
				}
			}

			network.Actions = resource.Actions

			// Преобразуем обратно в ResourceDefinition
			resources[i] = ResourceDefinition{
				Type:       infra.ResourceNetwork,
				Name:       network.Name,
				Properties: map[string]interface{}{"cidr": network.CIDR, "gateway": network.Gateway, "dns_servers": network.DNSServers},
				Actions:    network.Actions,
			}
		}
	}
	return resources
}
