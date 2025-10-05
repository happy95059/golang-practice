package exercises

// ## 7. struct（結構體）基本用法
type Person struct {
	Name string
	Age  int
}

func CreatePerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
