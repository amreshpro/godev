Go interfaces define a set of method signatures. Any type that implements those methods satisfies the interface—no explicit declaration needed.

```go
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

```

🔹 Kyu chahiye?
Interface allows polymorphism.

It helps in mocking dependencies in tests.

Decouples your code for flexibility.

🔹 Pros:
Loosely coupled design.

Easy to mock for testing.

Useful in DI (Dependency Injection).

🔹 Cons:
Can make debugging harder (interface wrapping).

No generics before Go 1.18, so often overused.

🔹 Industry Use:
Clean architecture → Repositories, Services, etc. use interfaces.

Logging, Storage, APIs, Testing.