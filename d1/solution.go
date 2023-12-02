package d1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solution(fileName string) int {
	val := 0
	fl, err := os.OpenFile(fmt.Sprintf("./%s", fileName), os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		val += findVal(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return val
}
func checkIsNum(str byte) (int, bool) {
	numbers := strings.Split("123456789", "")
	for k, v := range numbers {
		if v == string(str) {
			return k + 1, true
		}
	}
	return 0, false
}
func findVal(str string) int {
	var l, h, lVal, hVal = 0, len(str) - 1, 0, 0
	for l <= h {
		if val, ok := checkIsNum(str[l]); ok {
			lVal = val
		} else {
			l++
		}

		if val, ok := checkIsNum(str[h]); ok {
			hVal = val
		} else {
			h--
		}
		if lVal > 0 && hVal > 0 {
			break
		}
	}
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for k, v := range nums {
		if i := strings.Index(str, v); i >= 0 && i < l {
			lVal = k + 1
			l = i
		}
		if i := strings.LastIndex(str, v); i >= 0 && i > h {
			hVal = k + 1
			h = i
		}

	}
	return lVal*10 + hVal

	// for (lv == 0 || hv == 0) && l <= h {
	// 	if v, err := strconv.Atoi(string(str[l])); err == nil {
	// 		lv = v
	// 	} else {
	// 		l++
	// 	}
	// 	if v, err := strconv.Atoi(string(str[h])); err == nil {
	// 		hv = v
	// 	} else {
	// 		h--
	// 	}
	// }
	// fmt.Println(str, len(str)-1, l, h, lv, hv)

	// arr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// for i := 0; i < len(arr); i++ {
	// 	newL := strings.Index(str, arr[i])
	// 	if newL > -1 && newL < l {
	// 		l = newL
	// 		lv = i + 1
	// 	}
	// 	newH := strings.LastIndex(str, arr[i])
	// 	if newH > -1 && newH > h {

	// 		h = newH
	// 		hv = i + 1
	// 	}
	// }
	// return 10*lv + hv
}
