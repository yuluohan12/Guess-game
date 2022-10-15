//------------------------------------------------------------------lv2---------------------------------
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	var str2 string
	fmt.Print("input the ufo name,the people name:")
	fmt.Scanf("%s,%s", &str1, &str2)
	a := SwitchNumber(strings.ToUpper(str1))
	b := SwitchNumber(strings.ToUpper(str2))
	if a == b {
		fmt.Println("GO")

	} else {
		fmt.Println("STAY")
	}

}
func SwitchNumber(str string) int {
	var as []int
	var sum int
	sum = 1
	for _, v := range str {
		as = append(as, int(v))
	}
	for _, i := range as {
		sum *= i
	}
	number := sum % 47
	return number

}
