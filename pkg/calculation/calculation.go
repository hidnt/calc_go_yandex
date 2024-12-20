package calculation

import (
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	var parts []string
	var curPart string

	for _, char := range expression {
		if char == ' ' {
			continue
		}
		if char == '-' && (curPart == "" || curPart == "(") {
			curPart += string(char)
			continue
		}
		if strings.ContainsAny(string(char), "0123456789.") {
			curPart += string(char)
			continue
		}
		if curPart != "" {
			parts = append(parts, curPart)
			curPart = ""
		}
		if strings.ContainsAny(string(char), "+*/()-") {
			parts = append(parts, string(char))
			continue
		}
		return 0, ErrUnknownOp
	}

	if curPart != "" {
		parts = append(parts, curPart)
	}

	var nums []float64
	var operators []string

	priority := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"(": 0,
	}

	calculate := func() error {
		if len(operators) == 0 {
			return ErrNotEnoughtOp
		}
		if len(nums) < 2 && len(operators) >= 1 {
			err := ErrNotEnoughtNums
			return err
		}

		right := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		left := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		var result float64
		switch operator {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				return ErrDivByZero
			}
			result = left / right
		default:
			return ErrUnknownOp
		}
		nums = append(nums, result)
		return nil
	}

	for _, part := range parts {
		if num, err := strconv.ParseFloat(part, 64); err == nil {
			nums = append(nums, num)
			continue
		}
		if part == "(" {
			operators = append(operators, part)
		} else if part == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				if err := calculate(); err != nil {
					return 0, err
				}
			}
			if len(operators) == 0 {
				return 0, ErrIncorrectPriorOp
			}
			operators = operators[:len(operators)-1]
		} else {
			for len(operators) > 0 && priority[operators[len(operators)-1]] >= priority[part] {
				if err := calculate(); err != nil {
					return 0, err
				}
			}
			operators = append(operators, part)
		}
	}

	for len(operators) > 0 {
		if err := calculate(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, ErrCalc
	}

	return nums[0], nil
}
