package captcha

import "fmt"

var (
	mOperand = map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eigth",
		9: "nine",
	}
	mOperator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}
	gender = []Gender{
		format0{},
		format1{},
	}
)

type Gender interface {
	make(format, left, operator, right int) string
}

type format0 struct {
}

type format1 struct {
}

func (f format0) make(format, left, operator, right int) string {
	return fmt.Sprintf("%d %s %s", left, mOperator[operator], mOperand[right])
}
func (f format1) make(format, left, operator, right int) string {
	return fmt.Sprintf("%s %s %d", mOperand[left], mOperator[operator], right)
}

func Gen(format, left, operator, right int) string {
	return gender[format].make(format, left, operator, right)
}
