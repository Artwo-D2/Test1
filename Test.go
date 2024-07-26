package main
import "fmt"
import "strconv"

func ToRoman(num int) string {
	numerals := []struct {
		value int
		sym   string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	result := ""
	for _, numeral := range numerals {
		for num >= numeral.value {
			result += numeral.sym
			num -= numeral.value
		}
	}
	return result
}

func ToArabic (input string) (sum int){
        romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	sum = 0
	prevValue := 0
	for i := len(input) - 1; i >= 0; i-- {
		value := romanValues[rune(input[i])]
		if value < prevValue {
			sum -= value
		} else {
			sum += value
		}
		prevValue = value
	}
	return sum
}

func TypeOfNum(input string) (isRoman, isArabic bool){
    	validRomans := map[rune]bool{
		'I': true, 'V': true, 'X': true	}
	validArabics := map[rune]bool{
		'1': true, '2': true, '3': true,'4': true, '5': true, '6': true,'7': true, '8': true, '9': true,'0': true	}
	for _, char := range input {
		if validRomans[char] {
			isRoman=true
		}
	}
		for _, char := range input {
		if validArabics[char] {
			isArabic=true
		}
	}
	return isRoman, isArabic
}

func GetValue (input string) (value1 string, value2 string, operand string) {
        for i:=0;i<len(input);i++{
        if string(input[i])=="+"||string(input[i])=="-"||string(input[i])=="*"||string(input[i])=="/"{
            operand=string(input[i])
            i++
        }
        if operand==""{
        value1+=string(input[i])
        } else {
            value2+=string(input[i])
        }
    }
    return
}

func main() {
    for ;;{
    var input string
    var intVar1, intVar2 int
    fmt.Scan(&input)
    isRoman, isArabic:=TypeOfNum(input)
    value1, value2, operand:= GetValue(input)
    if isRoman==isArabic||operand==""{
        panic("panic")
    }
    if isRoman{
        intVar1=ToArabic(value1)
        intVar2=ToArabic(value2)
        if intVar1>=1&&intVar2>=1&&intVar1<=10&&intVar2<=10{
    switch operand {
	case "+":
		fmt.Println(ToRoman(intVar1+intVar2))
	case "-":
	if intVar1-intVar2>0{
		fmt.Println(ToRoman(intVar1-intVar2))
	}else {
	    panic("panic")
	}
	case "*":
	    fmt.Println(ToRoman(intVar1*intVar2))
	case "/":
		fmt.Println(ToRoman(intVar1/intVar2))
        }
        }else {
            panic("panic")
        }
    } else{
    intVar1, _=strconv.Atoi(value1)
    intVar2, _=strconv.Atoi(value2)
    if intVar1<1||intVar2<1||intVar1>10||intVar2>10{
    panic("panic")
    }
    switch operand {
	case "+":
		fmt.Println(intVar1+intVar2)
	case "-":
		fmt.Println(intVar1-intVar2)
	case "*":
	    fmt.Println(intVar1*intVar2)
	case "/":
		fmt.Println(ToRoman(intVar1/intVar2))
	}
    }
    }
}