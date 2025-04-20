package main

// Go interfaces define a set of method signatures. Any type that implements those methods satisfies the interface—no explicit declaration needed.


type Animal interface{
	Speaks() string
} 

type Dog struct {} // any type


func (dog Dog) Speaks() string{
	return "bho-bho"
}
func main(){

}