package exercise3

type Maintenance struct {
	name  string
	price float64
}

type MaintenanceCollection []Maintenance

func (maintenances MaintenanceCollection) sum() float64 {
	accumulator := 0.0

	for _, maintenance := range maintenances {
		accumulator += maintenance.price
	}

	return accumulator
}
