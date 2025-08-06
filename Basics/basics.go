package main

//printing a string
// func main() {
// 	fmt.Println("Hello World!")
// }

// typed and inferred
// func main() {
// 	var student1 string = "John" //type is string
// 	var student2 = "Jane"        //type is inferred
// 	x := 2                       //type is inferred

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)
// }

// declaring vars without initialing
// func main() {
// 	var a string
// 	var b int
// 	var c bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// use of := here, ":=" can only be used within scope/method
// func main() {
// 	name := "nehal"
// 	fmt.Println(name)
// }

// "type" keyword, only one type is used per line
// func main() {
// 	var a, b = 6, "Hello"
// 	c, d := 7, "World!"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// "const" keyword
// const PI = 3.14

// func main() {
// 	fmt.Println(PI)
// }

// "typed constants" are declared with defined type
// const A int = 1

// func main() {
// 	fmt.Println(A)
// }

// "untyped constants" are declared without defined type
// const A = 1

// func main() {
// 	fmt.Println(A)
// }

// Multiple Constants Declaration
// const (
// 	A int = 1
// 	B     = 3.14
// 	C     = "nehal learning GOOOO!"
// )

// func main() {
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(C)
// }

// Println() Function
// func main() {
// 	var i, j string = "Nehal", "Salman"

// 	fmt.Println(i, j)
// }

/* Printf() Function
 Printf() function first formats its argument based on the given formatting verb and
 then prints them.
 Here we will use two formatting verbs:
	%v is used to print the value of the arguments
	%T is used to print the type of the arguments
*/

// func main() {
// 	var i string = "Syed"
// 	var j int = 22
// 	k := true

// 	fmt.Printf("i has value: %v and type: %T\n", i, i)
// 	fmt.Printf("j has value: %v and type: %T\n", j, j)
// 	fmt.Printf("k has value: %v and type: %T", k, k)
// }

//General Formatting Verbs for printf()
/* The following verbs can be used with all data types:

Verb	Description
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign

*/

// func main() {
// 	var i = 15.5
// 	var txt = "NIN HAO!"

// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v\n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)

// 	fmt.Printf("%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }

// Integer Formatting Verbs
/* Verb	Description
%b		Base 2
%d		Base 10
%+d		Base 10 and always show sign
%o		Base 8
%O		Base 8, with leading 0o
%x		Base 16, lowercase
%X		Base 16, uppercase
%#x		Base 16, with leading 0x
%4d		Pad with spaces (width 4, right justified)
%-4d	Pad with spaces (width 4, left justified)
%04d	Pad with zeroes (width 4 )  */

// func main() {
// 	var i = 15

// 	fmt.Printf("%b\n", i)
// 	fmt.Printf("%d\n", i)
// 	fmt.Printf("%+d\n", i)
// 	fmt.Printf("%o\n", i)
// 	fmt.Printf("%O\n", i)
// 	fmt.Printf("%x\n", i)
// 	fmt.Printf("%X\n", i)
// 	fmt.Printf("%#x\n", i)
// 	fmt.Printf("%4d\n", i)
// 	fmt.Printf("%-4d\n", i)
// 	fmt.Printf("%04d\n", i)
// }

// String Formatting Verbs
/* Verb	Description
%s	Prints the value as plain string
%q	Prints the value as a double-quoted string
%8s	Prints the value as plain string (width 8, right justified)
%-8s	Prints the value as plain string (width 8, left justified)
%x	Prints the value as hex dump of byte values
% x	Prints the value as hex dump with spaces */

// func main() {
// 	var txt = "Nehal"

// 	fmt.Printf("%s\n", txt)
// 	fmt.Printf("%q\n", txt)
// 	fmt.Printf("%8s\n", txt)
// 	fmt.Printf("%-8s\n", txt)
// 	fmt.Printf("%x\n", txt)
// 	fmt.Printf("% x\n", txt)
// }

//Boolean Formatting Verbs
// Verb	Description
// %t	Value of the boolean operator in true or false format (same as using %v)

// func main() {
// 	var i = true
// 	var j = false

// 	fmt.Printf("%t\n", i)
// 	fmt.Printf("%t\n", j)
// }

// Float Formatting Verbs
/*
Verb	Description
%e	Scientific notation with 'e' as exponent
%f	Decimal point, no exponent
%.2f	Default width, precision 2
%6.2f	Width 6, precision 2
%g	Exponent as needed, only necessary digits */

// func main() {
// 	var i = 3.141

// 	fmt.Printf("%e\n", i)
// 	fmt.Printf("%f\n", i)
// 	fmt.Printf("%.2f\n", i)
// 	fmt.Printf("%6.2f\n", i)
// 	fmt.Printf("%g\n", i)
// }
