package main

import "fmt"

//reference types ((pointers, slices, maps, functions, 	channels))

//interface Types

//func main() {
//////////////////
////poioter///////
//////////////////
/*var myInt int
	myInt = 10
  	fmt.Println(myInt)
x := 10
myFirstPointer := &x

fmt.Println("x es", x)
fmt.Println("myfirstPointer is", myFirstPointer)

*myFirstPointer = 15

fmt.Println("x is now", x)

//changeValueOFPointer(&x)

fmt.Println("after  function call, x is now", x)*/

//////////////////
////Slices////////
//////////////////

/* var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	//fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}
	fmt.Println("Element 0 is", animals[0])
	fmt.Println("Fist two elements are", animals[0:2])
	fmt.Println("the slice is", len(animals), "elemets long")
	fmt.Print("is it sorted?: ", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Print(", is it now?: ", sort.StringsAreSorted(animals))
	sort.Strings(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}
func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]

	return a
}
//////////////////
//////Maps////////
//////////////////
func main() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//delete(intMap, "four")
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is int not on the  map")
	}
	intMap["two"] = 4*/

//////////////////
//////Func////////
//////////////////

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s  says %s", a.Name, a.Sound)
	fmt.Println()
}
func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 3
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()

}

/* func main() {
	myTotal := sumMany(2, 3, 4, 5, -6, 77, -1)
	fmt.Println(myTotal)
} */
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total

}
