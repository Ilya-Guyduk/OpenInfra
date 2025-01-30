package openinfra

// ExecuteAction выполняет действие над ресурсом.
func (r *Resource) ExecuteAction(action string) error {
	for _, a := range r.Actions {
		if a == action {
			log.Printf("Executing action '%s' on resource '%s'", action, r.Name)
			return nil
		}
	}
	return fmt.Errorf("action '%s' not supported for resource '%s'", action, r.Name)
}
