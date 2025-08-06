package main

// import "fmt"

// // Base "class"
// type Animal struct {
// 	Name string
// }

// // Method on Animal
// func (a Animal) Speak() {
// 	fmt.Println(a.Name, "makes a sound")
// }

// // Derived "class" using composition
// type Dog struct {
// 	Animal // Embedded struct (inherit behavior)
// 	Breed  string
// }

// // Override method using receiver on Dog
// func (d Dog) Speak() {
// 	fmt.Println(d.Name, "barks! Breed:", d.Breed)
// }

// type Speaker interface {
// 	Speak()
// }

// func main() {
// 	a := Animal{Name: "GenericAnimal"}
// 	d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Labrador"}

// 	// Interface polymorphism
// 	var s Speaker

// 	s = a
// 	s.Speak() // Output: GenericAnimal makes a sound

// 	s = d
// 	s.Speak() // Output: Buddy barks! Breed: Labrador
// }
