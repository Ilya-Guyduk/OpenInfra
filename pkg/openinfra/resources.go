package openinfra

// GetResourceByName возвращает ресурс по его имени.
func (o *OpenInfra) GetResourceByName(name string) (*Resource, error) {
	for _, resource := range o.Resources {
		if resource.Name == name {
			return &resource, nil
		}
	}
	return nil, fmt.Errorf("resource '%s' not found", name)
}

// GetDependencies возвращает зависимости для указанного ресурса.
func (o *OpenInfra) GetDependencies(resourceName string) ([]string, error) {
	for _, dep := range o.Dependencies {
		if dep.Resource == resourceName {
			return dep.DependsOn, nil
		}
	}
	return nil, fmt.Errorf("no dependencies found for resource '%s'", resourceName)
}
