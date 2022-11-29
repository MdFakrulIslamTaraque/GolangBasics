package main

import (
	// newpackage1 "HELLO/newfolder"
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"

	// "io/ioutil"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"rsc.io/quote"
)

func main() {
	// a := newpackage1.NewStruct{}
	// fmt.Println(a.X)
	// LearnVarLiterals()
	// LearnDeclaration()
	// LearnDecisonFLow()
	// LearnLoop()
	// DeferF()
	// LearnPointer()
	// LearnStruct()
	// LearnArray()
	// LearnSlices()
	// LearnMaps()
	// LearnSliceOfslice()
	// LearnFuncAsParam()
	// LearnFunctionAsClosures()
	// LearnMethods()
	// LearnInterface()
	// LearnStringer()
	// LearnErrors()
	// LearnReader()
	// LearnGenerics()
	// LearnGoRoutines()
	// LearnChannels()
	// LearnSyncMutex()
	LearnFileOperation()
}
func LearnDeclaration() {

	fmt.Print("Hello, World!")
	fmt.Println(5 / 2.0)
	fmt.Println(quote.Go())

	var (
		a1 int = 68
		b1 int = 32
		k1 string
		l1 float32 = 3.678900
	)
	fmt.Printf("a1 = %v b1 = %v  k1 = %v l1 = %g\n", a1, b1, k1, l1)

	// Variable declaration
	var i uint64
	fmt.Printf("i is of type %T\n", i)

	// Dynamic variable assignment
	y := 32.32
	var x = "feel"

	fmt.Printf("Type of y = %T\n", y)
	fmt.Printf("Type of x = %T\n", x)

	// Multivariable variable assignment without type
	var a, b, c = 3, 2.5, "foo"
	fmt.Printf("a is of %T\n", a)
	fmt.Printf("b is of %T\n", b)
	fmt.Printf("c is of %T\n", c)

	// modification of variables
	c = "go"
	fmt.Println(c)

	// const print and modify
	const PI = 22.0 / 7.0
	const (
		CA    = 0
		Ang   = 90.00
		Greet = "Hello"
	)
	fmt.Printf("Type of Ang = %T\n", Ang)
	fmt.Println(PI)
}
func LearnVarLiterals() {
	// integer literals
	il1 := 0xFee //hex
	il2 := 072   //octal
	il3 := 10e5
	fmt.Print(il1)
	fmt.Print(il2)
	fmt.Print(il3)

	// format specifier
	var i1, f1, s1 = 10, 10.3, "hello"
	fmt.Printf("\nint1 = %d & float1 = %f & string1 = %s\n", i1, f1, s1)
	fmt.Printf("Value of i1 = %v & type of i1 = %T\n", i1, i1)

	var st1 string = "Hello world"
	fmt.Printf("Value of st1(normal) = %v & value of st1(Go syntex format) = %#v & type of st1= %T \n", st1, st1, st1)

	var d1 int = 12
	fmt.Printf("in binary=%b\n", d1)
	fmt.Printf("in octal=%o\n", d1)
	fmt.Printf("in octal(starting with 0o)=%O\n", d1)
	fmt.Printf("in octal(small letter)=%x\n", d1)
	fmt.Printf("in octal(cap letter)=%X\n", d1)
	fmt.Printf("in octal(small,starting with 0x)=%#x\n", d1)
	fmt.Printf("in octal(cap,starting with 0X)=%#X\n", d1)
	fmt.Printf("left spaced=%4d\n", d1)
	fmt.Printf("right spaced=%-4d", d1)
	fmt.Printf("zero padding=%04d\n", d1)
}
func LearnDecisonFLow() {
	var chk = 30

	if chk == 20 {
		fmt.Println("Equal to 20")
	} else if chk == 25 {
		fmt.Println("equal to 25")
	} else {
		fmt.Println("not equal to 20 or 25")
	}
	fmt.Println("30")

	// The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP.
	//The difference is that it only runs the matched case so it does not need a break statement.
	var mark = 60
	var grade string

	switch mark {
	case 80:
		grade = "A+"
	case 75:
		grade = "A"
	case 70:
		grade = "A-"
	case 65:
		grade = "B+"
	case 60:
		grade = "B"
	case 55:
		grade = "B-"
	case 50:
		grade = "C"
	case 45:
		grade = "D"
	case 40:
		grade = "D-"
	default:
		grade = "F"
	}

	// একাধিক value ওয়ালা case হলে, comma দিয়েও কাজ করা যাচ্ছে
	switch grade {
	case "F":
		fmt.Println("Better Luck next time")
	case "A+", "A", "A-":
		fmt.Println("A class")
	case "B+", "B", "B-":
		fmt.Println("B class")
	case "C+", "C", "C-":
		fmt.Println("C class")
	case "D", "D-":
		fmt.Println("D class")
	}
	// বাইরে switch declare করে না দেয়াতে, case এর মাঝে একাধিক variable ও use করতে পারছি
	switch {
	case grade == "F":
		fmt.Println("Log Astese prep ne beta")
	case mark >= 40:
		fmt.Println("Khub vab barse!")
	}
}
func LearnLoop() {
	// Looping

	// single loop
	for j := 10; j >= 0; j-- {
		fmt.Println(j)
	}

	// fmt.Printf("Type of j=%T\n", j) error: j is in the scope of for loop

	// nested loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("Multiplication of %d\n", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}

	// break & continue
	for i := 0; ; i++ {
		if i == 10 {
			break
		}
		if i == 5 {
			continue
		}
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}

	var a2 int
	var b2 int = 15

	for a2 < b2 {
		a2++
		fmt.Println(a2)
	}
}
func DeferF() {
	fmt.Println("Defer function called")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Defer function finished")
}
func LearnPointer() {
	i, j := 42, 100

	p := &i

	fmt.Println("*p = ", *p)
	*p = 21
	fmt.Println("i = ", i)

	q := &j
	fmt.Println("*q = ", *q)
	*q = *q / 5
	fmt.Println("j = ", j)
}

type Vertex struct {
	a int
	b float32
}

func LearnStruct() {
	v := Vertex{10, 20}         //struct intialization
	fmt.Println("struct = ", v) //printing the struct

	v.a = 10 * 10 // struct field accessing through a dot
	fmt.Println("v.a = ", v.a)

	v.b = 10.202
	fmt.Println("v.b = ", v.b)

	structPtr := &v
	fmt.Println("*structPtr = ", *structPtr)   // printing the struct through pointer access
	fmt.Println("structPtr.a = ", structPtr.a) // printing the struct filed through pointer access
	fmt.Println("structPtr.b = ", structPtr.b) // printing the struct field through pointer access

	structPtr.a = 1000
	structPtr.b = 1000.12334
	fmt.Println("*structPtr = ", *structPtr)

	fmt.Printf("Type of v = %T  && Type of structptr = %T\n\n", v, structPtr)

	var (
		sl3 = Vertex{}
		sl1 = Vertex{1, 2}
		sl2 = Vertex{a: 10}
		sl4 = &Vertex{10, 20.00}
	)

	fmt.Println(sl1, sl2, sl3, sl4)

	sl1.a = 6868

	fmt.Println("struct after changing = ", v) //printing the struct
}
func LearnArray() {
	var a [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("array a = ", a)
	a[0] = 100
	fmt.Println("array a = ", a)
}
func LearnSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("before slice the prime array = ", primes)

	var slice1 []int = primes[1:4]
	fmt.Println("slice1 = ", slice1)

	var names = [5]string{
		"Tareq",
		"Toki",
		"Toba",
		"Rafiq",
		"Ferdousi",
	}
	fmt.Println("names = ", names)
	var childSlice []string = names[0:3]
	var parentSlice []string = names[3:5]

	var shiblingSlice []string = names[1:3]

	fmt.Println("childSlice = ", childSlice, " parentSlice = ", parentSlice, "siblings = ", shiblingSlice)

	shiblingSlice[0] = "Ruhul"

	fmt.Println("childSlice = ", childSlice, " parentSlice = ", parentSlice, "siblings = ", shiblingSlice)
	fmt.Printf("type of name array = %T\n", names)
	fmt.Printf("type of name array = %T\n", childSlice)
	fmt.Printf("type of name array = %T\n", parentSlice)

	sliceLit := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sliceLit = ", sliceLit)

	type struct1 struct {
		f1 int
		f2 string
	}

	sliceStruct := []struct1{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
	}

	fmt.Println("sliceStruct = ", sliceStruct)
	fmt.Println(sliceStruct[0].f1)

	rightOpenName := names[1:]
	fmt.Println("rightOpenName = ", rightOpenName)

	rightOpenName = names[:3]
	fmt.Println("rightOpenName = ", rightOpenName)

	s11 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("s11 = %v || len(s11) =  %v || cap(s11) = %v\n", s11, len(s11), cap(s11))

	s11 = s11[:0]
	fmt.Printf("after s11[:0]  --> s11 = %v || len(s11) =  %v || cap(s11) = %v\n", s11, len(s11), cap(s11))

	s11 = s11[:4]
	fmt.Printf("after s11[:4]--> s11 = %v || len(s11) =  %v || cap(s11) = %v\n", s11, len(s11), cap(s11))

	s11 = s11[2:]
	fmt.Printf("after s11[2:] --> s11 = %v || len(s11) =  %v || cap(s11) = %v\n", s11, len(s11), cap(s11))

	makeSlice := make([]int, 0, 5)
	fmt.Printf("makeSlice = %v || len(makeSlice) =  %v || cap(makeSlice) = %v\n", makeSlice, len(makeSlice), cap(makeSlice))

	makeSlice = makeSlice[:4]
	fmt.Printf("[:4] --> makeSlice = %v || len(makeSlice) =  %v || cap(makeSlice) = %v\n", makeSlice, len(makeSlice), cap(makeSlice))

	makeSlice = makeSlice[1:]
	fmt.Printf("[1:] --> makeSlice = %v || len(makeSlice) =  %v || cap(makeSlice) = %v\n", makeSlice, len(makeSlice), cap(makeSlice))

	var testSlice = []int{}
	fmt.Printf("testSlice = %v || len(testSlice) =  %v	|| cap(makeSlice) = %v\n", testSlice, len(testSlice), cap(testSlice))

	testSlice = append(testSlice, 1, 2, 4, 8, 16, 32, 64, 128)
	fmt.Printf("testSlice = %v || len(testSlice) =  %v	|| cap(makeSlice) = %v\n", testSlice, len(testSlice), cap(testSlice))

	for i, v := range testSlice {
		fmt.Printf("testSlice[%v] = %v\n", i, v)
	}

	fmt.Println("Only values")

	//ok
	for i := range testSlice {
		fmt.Printf("%v-th element = %v\n", i, testSlice[i])
	}

	for _, val := range testSlice {
		fmt.Printf("values = %v\n", val)
	}
}
func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	res := make(map[string]int)
	for _, word := range words {
		res[word]++
	}
	return res
}
func LearnMaps() {
	s1 := "hello the I am Learning go maps, go routines and trying to understand the concurrent programming in go"
	fmt.Println("word count of 'hello the I am Learning go maps, go routines and trying to understand the concurrent programming in go' : ", WordCount(s1))
	fmt.Println("Learning maps")
	m := make(map[string]string)

	m["js"] = "javascript"
	m["py"] = "python"
	m["rb"] = "ruby"

	fmt.Println("List of all languages: ", m)
	fmt.Println("m[js] = ", m["js"])
	fmt.Println("m[py] = ", m["py"])
	fmt.Println("m[rb] = ", m["rb"])

	for key, val := range m {
		fmt.Println("m[", key, "] = ", val)
	}

	type vertex2 struct {
		Lat, Long float64
	}

	//initializing map with make, (zerod map)
	var m2 = make(map[string]vertex2)

	//adding elements to the map
	m2["cuet"] = vertex2{-40.68433, -74.39967}
	m2["buet"] = vertex2{-41.68433, -74.39967}
	m2["ruet"] = vertex2{-42.68433, -74.39967}
	m2["kuet"] = vertex2{-42.68433, -75.39967}
	fmt.Println("detailed locations = ", m2)

	//removing elements from map with key
	delete(m2, "cuet")
	fmt.Println("detailed locations = ", m2)

	//retrive an element from map
	fmt.Println("detailed location of ruet = ", m2["ruet"])

	//element, ok of map: ok = [true, false], if ok == false, then element is the zero value for the map's element type.
	elem, ok := m2["mit"]
	if ok != true {
		fmt.Println("mit not extists")
	} else {
		fmt.Println("detailed location of mit = ", elem)
	}
}

func LearnSliceOfslice() {
	board := make([][]string, 5)

	board[0] = []string{"A", "B", "C", "D"}
	board[1] = []string{"E", "F", "G", "H"}
	board[2] = []string{"I", "J", "K", "L"}
	board[3] = []string{"M", "N", "O", "P"}
	board[4] = []string{"Q", "R", "S", "T"}

	for _, rowV := range board {
		for _, colV := range rowV {
			fmt.Printf("%v ", colV)
		}
		fmt.Println()
	}

	fmt.Println()

	for i := range board {
		for j := range board[i] {
			fmt.Printf("%v ", board[i][j])
		}
		fmt.Println()
	}
}

func UseFuncAsParam(paramFn func(float64, float64) float64, a, b float64) float64 {
	return paramFn(a, b)
}
func LearnFuncAsParam() {
	//Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	Hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("using function as function argument: ", UseFuncAsParam(Hypot, 3, 4))
}
func UseFuncAsRetArg() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Fibonacci(n int) (int, []int) {
	sum := 0
	fibNums := make([]int, 0, n)
	fibNums = append(fibNums, 0, 1)
	for i := 2; i < n; i++ {
		sum = (fibNums[i-1] + fibNums[i-2])
		fibNums = append(fibNums, sum)
	}
	return sum, fibNums
}

func LearnFunctionAsClosures() {
	// A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	//For example, the UseFuncAsRetArg() function returns a closure. Each closure is bound to its own sum variable

	//here we at first makes a copy of the funciton(posAdder, negAdder) which were savong the sum=0 at the beginning of these functions
	//then each time we were passing i values to the posAdder, negAdder funcitons, i was calculated on the basis of initialized sum = 0 value,
	//we are not initialinzing the UseFuncAsRetArg() fuinciton each time and for that sum is not becoming 0 agaion and again
	posAdder, negAdder := UseFuncAsRetArg(), UseFuncAsRetArg()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			posAdder(i),
			negAdder(-i),
		)
	}
	fmt.Println("Ficbocacii function:")
	fmt.Println(Fibonacci(10))
}

type Vertex3 struct {
	x, y float64
}

// method with struct type(pointer receiver)
func (v Vertex3) Hypot() float64 {
	//(v Vertex3) --> is the specioal receiver argument
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// method with struct type(value receiver)
func (v Vertex3) Square() {
	v.x = v.x * v.x
	v.y = v.y * v.y
}

// method with struct type(pointer receiver)
func (v *Vertex3) SquareDeep() {
	v.x = v.x * v.x
	v.y = v.y * v.y
}

func AbsFunc(v Vertex3) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

type MyFLoat float64

func (f MyFLoat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func LearnMethods() {
	//Go does not have classes. However, you can define methods on types.
	//a method is just a function; with a receiver argument, which appears in its own argument list, between the func keyword and the method name.
	// using a normal funciton using the special receiver argument as the function parameter will also work fine
	//cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

	//method with struct type(value receiver)
	v := Vertex3{5, 12}
	fmt.Println("Mehtods with struct type: ", v.Hypot())

	//method with non struct type, usage is just like method with struct type
	newFloatVar := MyFLoat(-math.Sqrt2)
	fmt.Println("Methods with non struct type: ", newFloatVar.Abs())

	//method with struct type(pointer receiver): Methods with pointer receivers can modify the value to which the receiver points (as SquareDeep does here).
	//Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
	fmt.Println("before editing with value receiver, v:", v)
	v.Square()
	fmt.Println("after editing with value receiver, v:", v) //no permanet change

	v2 := Vertex3{5, 12}
	fmt.Println("before editing with value receiver, v2:", v2)
	v2.SquareDeep()

	//Go interprets the statement v2.SquareDeep() as (&v).SquareDeep() of method as SquareDeep() is a method with pointer receiver,
	//which is a convenience/ advantage rather than using function rather than method
	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
	fmt.Println("after editing with value receiver, v2:", v2) // permanent change!

	p := &Vertex3{4, 3}
	fmt.Println(AbsFunc(*p))
}

// Interfaces are named collections of method signatures(only method declaration: "method_name type")
type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func GetAreaPerimeter(g Geometry) {
	fmt.Println("Area: ", g.Area())           // function overridden(Area() of circle/rectangle is called on the type of the argument)
	fmt.Println("Perimeter: ", g.Perimeter()) // function overridden(Perimeter() of circle/rectangle is called on the type of the argument)
}

func LearnInterface() {
	//Interfaces are named collections of method signatures(only method declaration: "method_name type")
	fmt.Println("Larning interface")
	rect1 := Rectangle{3, 5}
	circle1 := Circle{5}
	GetAreaPerimeter(&rect1)
	GetAreaPerimeter(&circle1)
	fmt.Println("Area of rect: ", rect1.Area())
	fmt.Println("Perimeter of rect: ", rect1.Perimeter())

	fmt.Println("Area of circle: ", circle1.Area())
	fmt.Println("Perimete of circle: ", circle1.Perimeter())

	var object1 Geometry = &Rectangle{3, 5}
	// object1 := &Rectangle{3, 5}
	var object2 Geometry = &Circle{3}
	fmt.Printf("Area of object1: %v && Type of object: %T\n", object1.Area(), object1)
	fmt.Printf("Area of object2: %v && Type of object: %T\nn", object2.Area(), object2)

	// The empty interface : interface type that specifies zero methods
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var i interface{}
	Describe(i)
	i = 42
	Describe(i)
	i = "Hello"
	Describe(i)

	//Type assertions of interface: provides access to an interface value's underlying concrete value
	//t := i.(T) --> This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
	//If i does not hold a T, the statement will trigger a panic.
	//To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	// t, ok := i.(T)
	// 	If i holds a T, then t will be the underlying value and ok will be true.
	// If not, ok will be false and t will be the zero value of type T, and no panic occurs.

	var ii interface{} = 12

	if s, ok := ii.(string); ok {
		ii = s
	} else {
		fmt.Println("converting to string is not possible")
	}
	// fmt.Println(s)

	if s, ok := ii.(float64); ok {
		ii = s
		fmt.Println("converted to float64")
	} else {
		fmt.Println("converting to float64 is not possible")
	}
	// f, ok := ii.(float64)
	fmt.Printf("Type of ii = %T || value of ii = %v\n", ii, ii)

	// f = ii.(float64) // panic
	// fmt.Println(f)

	//Type switches of interface: A type switch is like a regular switch statement,
	// but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
	CheckInterfaceType(20)
	CheckInterfaceType("hello")
	CheckInterfaceType(12.23)
}
func Describe(i interface{}) {
	fmt.Printf("Type of i = %T && value = %v\n", i, i)
}

func CheckInterfaceType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type of %v is int\n", v)
	case string:
		fmt.Printf("Type of '%v' is string\n", v)
	default:
		fmt.Printf("I don't know about type of %v!\n", v)
	}
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("\nname: %v \nage: %v \nid: %v\n", p.Name, p.Age, p.ID)
}

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v\n", ip[0], ip[1], ip[2], ip[3])
}

func LearnStringer() {
	//Stringer is a type in the fmt package of Golang
	//Stringer is implemented by any value that has a String method.
	//A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
	//thats why, when we are printing the method with Stringer, instgead of printing the method prperties in default way,
	//it will print the inherited struct/non-struct fields in the custom manner
	//if we are using pointer, then we should carefully pass the address of the struct/non-struct
	l := Person{ID: 1, Name: "John", Age: 1}
	fmt.Println("Person1 details: ", &l)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"new":       {1, 2, 3, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}
}

type MyError struct {
	When time.Time
	What string
}

// ERRORS
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number:"
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func LearnErrors() {

	//Go programs express error state with error values.
	// The error type is a built-in interface similar to fmt.Stringer(return string, but interface name is Error)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	if ans, err := Sqrt(2); err != nil {
		fmt.Println(err, ans)
	} else {
		fmt.Println(ans)
	}
}

func LearnReader() {
	//"io" package specifies the io.Reader interface, which represents the read end of a stream of data
	//Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

	r := strings.NewReader("Hello, Reader! Its a reader...")
	b := make([]byte, 8) //slice of bytes, kind of basket/limit of reading from the reader's buffer

	for {
		n, err := r.Read(b) // n: number of bytes read
		// err: error returned(without EOF, if err is nil)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func Getindex[T comparable](s []T, x T) int { //generic function
	//here 'comparable' is a special constraint
	//that makes it possible to use the == and != operators on values of the type, it gets
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func LearnGenerics() {

	//comparable is a useful constraint that makes it possible to use the == and != operators on values of the type.
	//In this example, we use it to compare a value to all slice elements until a match is found. This Index function works for any type that supports comparison.
	si := []int{10, 15, 20, -5, 10}
	fmt.Println("Index of 10 in si = ", Getindex(si, -15))

	ss := []string{"foo", "bar", "baz", "foo", "bar", "baz"}
	fmt.Println("Index of 'foo' in ss = ", Getindex(ss, "foo")) //returns the first matched index
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func getSum(s string) int {
	arr := []int{10, 20, 30, 40, 60, 70, 80, 90, 100}
	sum := 0
	for i := 0; i < len(arr); i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(s, "-> ")
		fmt.Println(arr[i])
		sum += arr[i]
	}
	return sum
}

func LearnGoRoutines() {
	//Goroutine --> A lightweight thread, managed by the Go runtime.
	go say("Go routines") // starts a new goroutine running
	say("Hello world!")   // happens in the current goroutine

	go getSum("goroutine")
	getSum("normal")
	//Goroutines run in the same address space, so access to shared memory must be synchronized.

}

func getSSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func LearnChannels() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Channels are a typed conduit(pipe) through which you can send and receive values with the channel operator, <-
	c := make(chan int) //Like maps and slices, channels must be created before use:

	//By default, sends and receives block until the other side is ready.
	//This allows goroutines to synchronize without explicit locks or condition variables.
	go getSSum(s[:len(s)/2], c)
	go getSSum(s[len(s)/2:], c)

	// receive from c, here untill the end of receiving sum from c to x/y, we can't assign any other sum(of slice) to c again
	//here firstly , the sum of one half (first half/second) completes from the go routing and untill this value is assigned to any variable, goroutine can't start the sum of other half of the slice
	x, y := <-c, <-c

	//and only after getting the sum from the completion of both goroutines, x+y can't be completed
	fmt.Println(x, y, x+y)

	//=====================================Buffered Channels================================
	//Provide the buffer length as the second argument to make, to initialize a buffered channel
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// Sends to a buffered channel block, only when the buffer is full. Receives block when the buffer is empty.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//fatal error: all goroutines are asleep - deadlock!
	//as the buffered channel is empty, and it has nothing to send to other variable/ print
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	ch2 := make(chan int, 2)
	go getSSum(s[:len(s)/2], ch2)
	go getSSum(s[len(s)/2:], ch2)
	x, y = <-ch2, <-ch2
	fmt.Println(x, y, x+y)

	//===================================== Close on buffered Channels================================

	//A sender can close a channel to indicate that no more values will be sent.
	//Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression

	//v, ok := <-ch  //ok is false if there are no more values to receive and the channel is closed.
	// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

	//Channels aren't like files; you don't usually need to close them.
	//Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

	//here, the close() is using like the end of sending and receiving
	//if we didn't use close here, that means, the buffer is not fixed and the channel can send and receive more values
	//and thats why when we are using buffered channel in range loop, to fixe its length, we need to close the channel
	ch3 := make(chan int, 10)
	go channeledfibonacci(cap(ch3), ch3)
	for i := range ch3 {
		fmt.Print(i, " ")
	}

	//===================================== Select =======================================================

	//The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	//Defaul case in the select statement is run in the normal flow

	ch4 := make(chan int)
	quit := make(chan int)
	fmt.Println("\nSelect opration : ")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-ch4, " ") //sends and receives block until the other side is ready, so untill the ch4 gets ready, goroutine will wait
		}
		quit <- 0
	}()
	selectFibo(ch4, quit)
}

func channeledfibonacci(n int, c chan int) {
	fmt.Println("Capasity of channel: ", n)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//here, the close is using like the end of sending and receiving
	// if we didn't use close here, that means, the buffer is not fixed and the channel can send and receive more values
	//and thats why when we are using buffered channel in range loop, to fixe its length, we need to close the channel
	close(c)
}

func selectFibo(c, quit chan int) {
	x, y := 0, 1
	for {
		//at first , select will check which of the two channels are ready to run and the ready channel will be executed
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func LearnSyncMutex() {
	//channels are great for communication among goroutines, but we don't need communication and we use only one go routine(default thread)
	//and make sure that that only goroutine can access a variable  at a time without conflict
	//then we can use mitual exclusion and the data structure for mitual exclusion is mutex
	//Go's standard library provides mutual exclusion with sync.Mutex and its two methods: Lock && Unlock
	// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to:  Lock and Unlock

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.GetValue("somekey"))
}

type SafeCounter struct {
	mu sync.Mutex //Go's standard library provides mutual exclusion with sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}
func (c *SafeCounter) GetValue(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func ReadStats(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()

	stats, errs := f.Stat() //stas() retuns : file name, size, type, mode, IsDirectory, System, ModTme()

	if errs != nil {
		fmt.Println("Errors in stat: ", errs.Error())
		return
	}
	fmt.Println("File name:", stats.Name(), "| Last modification time:", stats.ModTime().Format("01:04:03PM"), " | Size:", stats.Size())
}
func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Errors in reading whole file", err.Error())
		return
	}
	fmt.Println(string(contents))
}

func ReadLine(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func ReadByWord(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()

	buf := make([]byte, size)

	for {
		totalREad, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		fmt.Println(string(buf[:totalREad]))
	}

}

func ReadConfig(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=") // [key, value]
		fmt.Println(raw[0], "->", raw[1])
	}
}

func WriteFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	_, werr := f.WriteString("Hello New File!")
	if werr != nil {
		fmt.Println(werr.Error())
		return
	}
}
func WriteBytes(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	size, werr := f.Write(data)
	if werr != nil {
		fmt.Println(werr.Error())
		return
	}
	fmt.Println(size, " data were written int the file")
}

func WriteLine(filename string, data []string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}
	defer f.Close()
	for _, val := range data {
		_, err := fmt.Fprintln(f, val)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func AppendFile(filename, data string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_WRONLY,
		fs.ModeAppend)
	if err != nil {
		fmt.Println("Errors in opening file", err.Error())
		return
	}

	_, fErr := fmt.Fprintln(f, data)
	if fErr != nil {
		fmt.Println(err.Error())
		return
	}
}

func LearnFileOperation() {
	// filename := "newWrite.dat"
	// ReadStats(filename)
	// ReadWholeFile(filename)
	// ReadByWord(filename)
	// ReadByBytes(filename, 8)
	// ReadConfig("test.cfg")

	// WriteFile(filename)
	// WriteBytes("binary.txt", []byte("hello byte file"))

	// lines := []string{
	// 	"First Line",
	// 	"Second Line",
	// 	"Third Line",
	// }
	newLine := "new line"
	// WriteLine("lineWrite.txt", lines)
	AppendFile("lineWrite.txt", newLine)
}
