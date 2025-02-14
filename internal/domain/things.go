package domain

// Thing – base type for all items.
type Thing struct {
	name   string
	number int
}

func (t Thing) GetNumber() int {
	return t.number
}

func (t Thing) GetName() string {
	return t.name
}

type Table struct {
	Thing
}

type Computer struct {
	Thing
}

func NewThing(name string, number int) Thing {
	return Thing{name: name, number: number}
}
