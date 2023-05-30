package exercise2

type Id struct {
	nextValue int
}

func (id *Id) New() int {
	if id.nextValue < 0 {
		panic("apenas valores positivos são permitidos")
	}

	id.nextValue++
	return id.nextValue
}
