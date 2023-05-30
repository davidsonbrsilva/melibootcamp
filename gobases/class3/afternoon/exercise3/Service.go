package exercise3

type Service struct {
	name          string
	price         float64
	workedMinutes float64
}

type ServiceCollection []Service

func (services ServiceCollection) sum() float64 {
	accumulator := 0.0

	for _, service := range services {
		if service.workedMinutes < 30 {
			service.workedMinutes = 30
		}

		accumulator += service.price * (service.workedMinutes / 60)
	}

	return accumulator
}
