package examples

import (
	"fmt"

	"github.com/mech-pig/bob"
)

type User struct {
	Name string
	Age  int
}

var userBuilder = bob.New(func() User {
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

func ExampleBuilder_Build() {
	fmt.Println(userBuilder.Build())
	// Output: {test 18}
}

func ExampleBuilder_Build_withOverrides() {
	fmt.Println(userBuilder.Build(nameIsBob, ageIs15))
	// Output: {bob 15}
}

func ExampleBuilder_BuildMany() {
	fmt.Println(userBuilder.BuildMany(3))
	// Output: [{test 18} {test 18} {test 18}]
}

func ExampleBuilder_BuildMany_withOverrides() {
	fmt.Println(userBuilder.BuildMany(3, func(i int, u User) User {
		u.Name = fmt.Sprint("test-", i)
		return u
	}))
	// Output: [{test-0 18} {test-1 18} {test-2 18}]
}

func ExampleBuilder_Override() {
	bobUserBuilder := userBuilder.Override(nameIsBob)
	fmt.Println(bobUserBuilder.Build())
	fmt.Println(bobUserBuilder.Build(ageIs15))
	fmt.Println(bobUserBuilder.BuildMany(2, func(i int, u User) User { u.Age = i; return u }))
	// Output:
	// {bob 18}
	// {bob 15}
	// [{bob 0} {bob 1}]
}
