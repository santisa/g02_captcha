package captcha

import "fmt"

var (
	mOperand = map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eigth",
		9:  "nine",
		10: "ten",
	}
	mOperator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}
)

func Gen(format, left, operator, right int) string {
	if format == 0 {
		return fmt.Sprintf("%d %s %s", left, mOperator[operator], mOperand[right])
	} else if format == 1 {
		return fmt.Sprintf("%s %s %d", mOperand[left], mOperator[operator], right)
	} else {
		return "Invalid 'format' (should be 0,1)"
	}
}
