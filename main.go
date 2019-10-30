package main

import (
	"fmt"
	"go101/g02_captcha/captcha"
)

func main() {
	fmt.Println("Input1: 0, 1, 0, 1")
	fmt.Println("Captcha1:", captcha.Gen(0, 1, 0, 1))
	fmt.Println("Input2: 1, 1, 0, 1")
	fmt.Println("Captcha2:", captcha.Gen(1, 1, 0, 1))

	fmt.Println("Input3: 0, 1, 0, 1")
	fmt.Println("Captcha3:", captcha.Gen(0, 1, 1, 1))
	fmt.Println("Input4: 1, 1, 1, 1")
	fmt.Println("Captcha4:", captcha.Gen(1, 1, 1, 1))

	fmt.Println("Input5: 0, 1, 2, 3")
	fmt.Println("Captcha5:", captcha.Gen(0, 1, 2, 3))
	fmt.Println("Input6: 1, 3, 2, 1")
	fmt.Println("Captcha6:", captcha.Gen(1, 3, 2, 1))

}
