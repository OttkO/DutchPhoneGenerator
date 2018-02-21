package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateDutchPhoneNumber(answer *string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	firstDigit := 0
	secondDigit := 6
	thirdDigit := r1.Intn(10)
	fourthDigit := r1.Intn(10)
	fifthDigit := r1.Intn(10)
	sixthDigit := r1.Intn(10)
	seventhDigit := r1.Intn(10)
	eightDigit := r1.Intn(10)
	ninthDigit := r1.Intn(10)
	tenthDigit := r1.Intn(10)
	*answer = fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d", firstDigit, secondDigit, thirdDigit, fourthDigit, fifthDigit, sixthDigit, seventhDigit, eightDigit, ninthDigit, tenthDigit)
}

func generateDutchPhoneNumberWithForLoop(answer *string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	*answer = "06"
	for i := 0; i < 8; i++ {
		*answer = fmt.Sprintf("%s%d", *answer, r1.Intn(10))
	}
}

func main() {
	result := "empty"
	//generateDutchPhoneNumber(&result)
	generateDutchPhoneNumberWithForLoop(&result)
	fmt.Println(result)
}
