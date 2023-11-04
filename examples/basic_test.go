package examples

import (
	"fmt"

	"github.com/mech-pig/bob"
)

type User struct {
	Name string
	Age  int
}

var userFactory = bob.New(func() User {
	return User{
		Name: "test",
		Age:  18,
	}
})

func nameIsBob(u User) User {
	u.Name = "bob"
	return u
}

func ageIs15(u User) User {
	u.Age = 15
	return u
}

func ExampleFactory_Build() {
	fmt.Println(userFactory.Build())
	// Output: {test 18}
}

func ExampleFactory_Build_withOverrides() {
	fmt.Println(userFactory.Build(nameIsBob, ageIs15))
	// Output: {bob 15}
}

func ExampleFactory_BuildMany() {
	fmt.Println(userFactory.BuildMany(3))
	// Output: [{test 18} {test 18} {test 18}]
}

func ExampleFactory_BuildMany_withOverrides() {
	fmt.Println(userFactory.BuildMany(3, func(i int, u User) User {
		u.Name = fmt.Sprint("test-", i)
		return u
	}))
	// Output: [{test-0 18} {test-1 18} {test-2 18}]
}

func ExampleFactory_Override() {
	bobUserFactory := userFactory.Override(nameIsBob)
	fmt.Println(bobUserFactory.Build())
	fmt.Println(bobUserFactory.Build(ageIs15))
	fmt.Println(bobUserFactory.BuildMany(2, func(i int, u User) User { u.Age = i; return u }))
	// Output:
	// {bob 18}
	// {bob 15}
	// [{bob 0} {bob 1}]
}
