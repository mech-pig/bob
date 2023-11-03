# bob 🔨

Generate fixtures easily.

Inspired by classical `factory_bot`-like fixtures replacement libraries, `bob` uses simple functions as the primary mechanism to override default values instead of relying on annotations or a builder-pattern API.

`bob` requires go 1.18+.

## Installation

```shell
go get github.com/mech-pig/bob
```

## Usage

### Create a factory
A factory is created by providing a function that is used to generate a default instance.

```go
type User struct {
    Name string
    Age int
}

userFactory := bob.New(func () User {
    return User{
        Name: "test",
        Age: 18,
    }
})
```

### Build an instance

Instances are created with the `Build` method.

```go
userFactory.Build()
```

The build method accepts an optional list of functions that are used to customize the instance. An override function takes an instance as input and returns a modified one.

```go
func nameIsBob(u User) User {
    u.Name = "bob"
    return u
}

func is15(u User) User {
    u.Age = 15
    return u
}

userFactory.Build(nameIsBob, is15)
```

### Derive factory from an existing one

The `Override` method can be used to derive a new factory from an existing one:

```go

bobUserFactory := userFactory.Override(nameIsBob)
bobUserFactory.Build()
```

### Build many instances

The `BuildMany` method is used to build multiple instances at once. It accepts an `int` to indicate the number of instances that will be generated.

```go
userFactor.BuildMany(5)
```

Like `Build`, it's possible to customise the generated instances by providing one or more overriding functions. Each of these function takes as input an index and the default instance

```go
userFactory.BuildMany(3, func (i int, u User) User {
    u.Name = fmt.Sprint("user-", i)
    return u
})
```


 