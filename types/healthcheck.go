package types

type HealthCheck struct {
	Status string
}

func NewHealthCheck(status string) *HealthCheck {
	return &HealthCheck{
		Status: status,
	}
}
