package main

// TODO: implement the following method **only** by using defer. The **first** operation is provided
// Step 1: if the number is divisible by 13 multiply it by 7
// Step 2: if the number is negative then take its absolute value
// Step 3: convert the number to a string and strip off the left-most character. If the resulting
// string is the empty string then make it "0". Convert the result back to an integer.
// Step 4: set the 0th, 1st, and 6th bit to 1
//
// For example: Defer(-104) -> (step 1) -728 -> (step 2) 728 -> (step 3) 28 -> (step 4) 95
func Defer(n int) (ret int) {
	// Step 1
	defer func() {
		if ret%13 == 0 {
			ret = ret * 7
		}
	}()
	ret = n
	return
}

// TODO: implement the NumberChecker interface for the `MyInt` type.
type MyInt int

func (n MyInt) IsEven() bool {
	return false
}

func (n MyInt) IsOdd() bool {
	return false
}

// Hints:
// - a number is prime if its only even divisors are 1 and itself
// - this function will only be tested with very small numbers
// - any number less than 2 is not considered prime
func (n MyInt) IsPrime() bool {
	return false
}

// TODO: implement the following function which takes both `int` and `MyInt` using approximation constraints (`~T`)
// The function should create a string in the following format:
// "==== ${msg} ===="
// where each "====" bar is has width characters
func MakeHeaderLine[T ~int](width T, msg string) string {
	return ""
}

// TODO: make an interface that includes the Read(filePath string) method found in util.
// These two concrete types already exist are are implemented.
type FileReader interface {
}

// TODO: construct a message that will be added to the log in the following format:
// "${msg} - ${obj1} - ${obj2}"
// If `obj1` or `obj2` are `int`, print them in hexadecimal (use "0x%X"), otherwise use the "%v" print
// format verb to convert `obj1` and `obj2` to strings
func FormatLogMessage(msg string, obj1 any, obj2 any) string {
	return ""
}

// TODO: make a struct that embeds the `Lion`, `Goat`, and `Snake` structs, which are all defined in
// util
type Chimera struct{}

// TODO: make an interface that embeds the three interfaces: `Roarer`, `Headbutter`, and `Hisser`.
// These interfaces and structs that implement the interfaces are defined in util.
type ChimeraI interface{}

// TODO: implement the `Priceable` interface (in util) on both `Sedan` and `Motorcycle`.
// Use pointer receivers.
// The price for motorcycles should be: (numWheels * 500) + (weight * 2) + (coolnessFactor * 1000)
// The price for sedans should be: (numWheels * 500) + (weight * 2) + (engineHorsepower * 10)
func (moto Motorcycle) CalcPrice() uint {
	return 0
}
func (sedan Sedan) CalcPrice() uint {
	return 0
}

// TODO: implement the function which adjusts price based on country
// Add the following import tax to the price of the vehicle based on the country:
// - USA: 150%
// - EU: 10%
// - CANADA: 7%
// Hint: convert the vehicle price to a float, the adjust price, then cast to `uint` (rounding down)
func AdjustPrice(p Priceable, territory Territory) uint {
	return 0
}

// TODO: double the elements in arr. This is an attempt to modify the caller's `arr`. This will be
// unsuccessful though, as golang will pass a copy of the array (pass by value)
func DoubleArr(arr [3]int) {}

// TODO: double the elements in arr. This should modify the caller's `arr` successfully.
func DoubleArrByPtr(arr *[3]int) {}

// TODO: double the elements in the slice. Slices are passed by reference automatically unlike
// arrays, which is why we don't explicitly need the a pointer to the slice.
func DoubleSlice(slice []int) {}

// TODO: implement a method on the receiver type (p Point2D) that calculates the area of a rectangle
// bounded by the points (0, 0) and (p.X, p.Y)
// Hint: area should always be non-negative
func (p Point2D) CalcArea() float64 {
	return 0.0
}

// TODO: implement a method on receiver type (p *Point2D) that swaps the X and Y components
func (p *Point2D) Transpose() {}

// TODO: calculate the sum of squares in the range 1 to num (inclusive). Do this using a gotos and
// labels **instead** of a for loop. Branching statements like if else statements are allowed.
func GotoLabelSumSquares(num int) int {
	return 0
}

// TODO: modify the following function to continue execution at the outermost loop when the
// innermost loop has an element that is negative. Perform this action before the element is added
// to the sum.
// Hint: use a break and label
func BreakLabel3D(matrix [][][]int) int {
	sum := 0
	for _, plane := range matrix {
		for _, row := range plane {
			for _, elem := range row {
				// TODO add condition here
				sum += elem
			}
		}
	}
	return sum
}

// TODO: modify the following function that calculates the sum of a linked list (in a funky way) to
// stop execution once a node with a negative value is encountered.
// The only modification needed is to:
// add a new case in the switch with a labeled break **inside** the switch statement.
func FunkySumLinkedList(node SLL) int {
	sum := 0
	var trav *SLL = &node
	for {
		switch {
		case trav.Val%2 == 0:
			sum += trav.Val * 10
		default:
			sum += trav.Val
		}
		if trav.Next == nil {
			break
		}
		trav = trav.Next
	}
	return sum
}
