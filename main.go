package main // * เรียกใช้ Package

import (
	"fmt" // * เรียกใช้ Package Libarary
)

// --------- Scope --------------------
var a1 = 10
var a2 = 20

//--------- Scope --------------------

// * ตัว Main
func main() {
	Simple()
	fmt.Println("----------------------------------------------------------------")

	//--------- Operator --------------------
	Operator()
	fmt.Println("----------------------------------------------------------------")

	//--------- Operator --------------------

	//--------- Scope -----------------------
	fmt.Println("Example : Scope")
	fmt.Println("a = ", a1)
	Processor() // * เรียกใช้ Functions
	fmt.Println("----------------------------------------------------------------")
	//--------- Scope -----------------------

	//--------- Defining Multiple Variables -----------------------
	DefiningMultipleVariables()
	fmt.Println("----------------------------------------------------------------")
	//--------- Defining Multiple Variables -----------------------

	//--------- ForLoop -----------------------
	Forloop()
	fmt.Println("----------------------------------------------------------------")
	//--------- ForLoop -----------------------

	//--------- If Condition -----------------------
	IfCondition()
	fmt.Println("----------------------------------------------------------------")

	//--------- If Condition -----------------------

	//--------- ElseIf Condition -----------------------
	ElseIfCondition()
	fmt.Println("----------------------------------------------------------------")
	//--------- ElseIf Condition -----------------------

	//--------- Switch -----------------------
	Switch()
	fmt.Println("----------------------------------------------------------------")
	//--------- Switch -----------------------

	//--------- Array -----------------------
	Arrays()
	fmt.Println("----------------------------------------------------------------")
	//--------- Array -----------------------

	//--------- Array -----------------------
	Slices()
	fmt.Println("----------------------------------------------------------------")
	//--------- Array -----------------------

	Paramid()

}

func Simple() {
	// x := "inet" // * เป็นการกำหนดตัวแปรที่เหมือนของ JS
	var x string  // * เป็นการกำหนดตัวแปรด้วยระบุชนิดข้อมูลชัดเจน (ข้อความ)
	var y int     // * เป็นการกำหนดตัวแปรด้วยระบุชนิดข้อมูลชัดเจน (number)
	var z float32 // * เป็นการกำหนดตัวแปรด้วยระบุชนิดข้อมูลชัดเจน (float)
	//
	x = "inter bkk"
	y = 20
	z = 20.2
	// print ข้อความ คล้ายกับ C
	fmt.Printf("string : %s \n", x)
	fmt.Printf("int : %d \n", y)
	fmt.Printf("float : %.2f \n", z)
}

func Operator() {
	a := 5
	b := 2

	fmt.Println("Example -: Operators")
	fmt.Println("Symbol\tMeaning\t\tExample")
	fmt.Println("+\tAddition\ta + b = ", a+b)
	fmt.Println("-\tSubstraction\ta - b = ", a-b)
	fmt.Println("*\tMultiplication\ta * b = ", a*b)
	fmt.Println("/\tDivision\ta / b = ", a/b)
	fmt.Println("%\tRemainder\ta % b =", a%b)
}

func Processor() {
	fmt.Println("b == ", a2)
	fmt.Println("a+b == ", a1+a2)
}

func DefiningMultipleVariables() {
	var (
		a = 1
		b = 1.2
		c = "Hello , Sir"
	)

	fmt.Println("Example-: Defining multiple variables")
	fmt.Printf("Variable a = %d and type is %T \n", a, a)
	fmt.Printf("Variable b = %g and type is %T \n", b, b)
	fmt.Printf("Variable c = %s and type is %T \n", c, c)
}

func Forloop() {
	fmt.Println("Example-: For loop")
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("\n================================================")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("\n================================================")

	var myWords = "AWSOME GO!"
	for i, v := range myWords {
		fmt.Println("Index = ", i, ", value = ", string(v))
	}
}

// for ซ้อน if
func IfCondition() {
	fmt.Println("Example-: If Condition")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("i = %d , Odd (เลขคี่) \n", i)
		} else {
			fmt.Printf("i = %d , Even (เลขคู่) \n", i)
		}
	}
}

func ElseIfCondition() {
	fmt.Println("Example-: ElseIF Condition.")
	for i := 1; i <= 10; i++ {
		if i == 2 {
			fmt.Printf("i = %d \n", i)
		} else if i == 4 {
			fmt.Printf("i = %d \n", i)
		} else if i == 10 {
			fmt.Printf("i = %d \n", i)
		}
	}
}

func Switch() {
	i := 2

	fmt.Println("Example-: Switch case condition.")

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Your i not in case.")
	}
}

func Arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("Example-: Arrays")
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func Slices() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{99, 98, 97, 96}
	combindedSlices := append(slice1, slice2...)
	slice3 := make([]int, 9)

	fmt.Println("Example-: Slices")
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("combindedSlices = ", combindedSlices)

	copy(slice3, slice2)
	fmt.Println("coppy slice2 to slice3 = ", slice3)
}

func Paramid() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// func Map() {
// 	obj := make(map[string]int)
// 	obj["pen"] = 10
// 	obj["pencil"] = 20

// 	obj2 := make(map[string]string)
// 	obj2["car_01"] = "Truck"
// 	obj2["car_02"] = "Sedan"

// 	fmt.Println("Example :- Maps")
// 	fmt.Println(obj)
// 	fmt.Println("Pren price %dus and Pencil price %dus.\n" , obj["pen"], obj["pencil"])

// }
