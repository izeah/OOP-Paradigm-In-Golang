package person

type person struct {
	FirstName  string
	MiddleName string
	LastName   string

	// encapsulation
	age int
}

func NewPerson(firstName, middleName, lastName string, age int) person {
	return person{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		age:        age,
	}
}

func (p person) GetAge() int {
	return p.age
}

func (p person) GetFullName() string {
	return p.FirstName + " " + p.MiddleName + " " + p.LastName
}

func (p person) UlangTahun() {
	p.age++
}

func (p *person) UpdateUsia(age int) {
	p.age = age
}
