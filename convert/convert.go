package convert

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var mapsValue = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

var aliasData = map[string]string{}

var creditData = map[string]float64{}

// CheckData - check if string in slice of string
func CheckData(word string, list []string) bool {
	for _, b := range list {
		if b == word {
			return true
		}
	}
	return false
}

// FindIdx - function to find index of string from slice of string
func FindIdx(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// ValidateInputData - function for validate input from user when user add new data
func ValidateInputData(arr []string) bool {
	ok := false
	lenArr := len(arr)
	if CheckData("is", arr) && CheckData("credits", arr) {
		roman := ""
		idxIs := FindIdx(arr, "is")
		if idxIs != 0 && lenArr == idxIs+3 {
			romanSlice := arr[:idxIs]
			sliceWORoman := []string{}
			prev := true
			for _, elm := range romanSlice {
				val, _ := getData(elm)
				if val != "" {
					if prev {
						roman += val
					} else {
						roman = ""
						break
					}
				} else {
					sliceWORoman = append(sliceWORoman, elm)
					prev = false
				}
			}
			convertRoman := ConvertData(roman)
			if len(sliceWORoman) == 1 && convertRoman != 0 {
				credit, err := strconv.Atoi(arr[idxIs+1])
				if err != nil {
					fmt.Println("I have no idea what you are talking about")
					return ok
				}
				creditValue := float64(credit) / float64(convertRoman)
				creditValue = math.Round(creditValue*100) / 100
				saveCreditData(sliceWORoman[0], creditValue)
				fmt.Println("Success Add Data")
				return true
			} else {
				fmt.Println("I have no idea what you are talking about")
			}
		} else {
			fmt.Println("I have no idea what you are talking about")
		}
	} else if CheckData("is", arr) && !CheckData("credits", arr) {
		idxIs := FindIdx(arr, "is")
		romanSlice := arr[:idxIs]
		if idxIs != 0 && lenArr == idxIs+2 && len(romanSlice) == 1 && saveData(arr[0], arr[len(arr)-1]) {
			fmt.Println("Success Add Data")
			return true
		} else {
			fmt.Println("I have no idea what you are talking about")
		}
	} else {
		fmt.Println("I have no idea what you are talking about")
	}
	return ok
}

// ValidateReadData - function for validate input from user when user read data
func ValidateReadData(arr []string) bool {
	ok := false
	lenArr := len(arr)
	if CheckData("many", arr) {
		idxMany := FindIdx(arr, "many")
		if idxMany == 1 && lenArr >= 6 && arr[0] == "how" && arr[2] == "credits" && arr[3] == "is" && arr[lenArr-1] == "?" {
			romanSlice := arr[4 : lenArr-1]
			roman := ""
			sliceWORoman := []string{}
			prev := true
			for _, elm := range romanSlice {
				val, _ := getData(elm)
				if val != "" {
					if prev {
						roman += val
					} else {
						roman = ""
						break
					}
				} else {
					sliceWORoman = append(sliceWORoman, elm)
					prev = false
				}
			}
			convertRoman := ConvertData(roman)
			lenSliceWORoman := len(sliceWORoman)
			// fmt.Println(convertRoman, lenSliceWORoman)
			if lenSliceWORoman >= 1 && convertRoman != 0 {
				dValues := []float64{}
				value := 1.0
				for _, elm := range sliceWORoman {
					dValue, _ := getCreditData(elm)
					if dValue == 0.0 {
						// fmt.Println("I have no idea what you are talking about")
						dValues = []float64{}
						break
					} else {
						dValues = append(dValues, dValue)
					}
				}
				if len(dValues) == 0 {
					fmt.Println("I have no idea what you are talking about")
				} else {
					for _, x := range dValues {
						value = value * x
					}
				}
				value = float64(convertRoman) * value
				result := ""
				for _, x := range romanSlice {
					result = result + x + " "
				}
				// valToStr := fmt.Sprintf(%2f, value)
				result += "is " + fmt.Sprintf("%.2f", value)
				fmt.Println(result)
				return true
			} else {
				fmt.Println("I have no idea what you are talking about")
			}
		} else {
			fmt.Println("I have no idea what you are talking about")
		}
	} else if CheckData("much", arr) {
		idxMuch := FindIdx(arr, "much")
		if idxMuch == 1 && lenArr >= 5 && arr[0] == "how" && arr[2] == "is" && arr[lenArr-1] == "?" {
			romanSlice := arr[3 : lenArr-1]
			roman := ""
			for _, elm := range romanSlice {
				val, _ := getData(elm)
				if val != "" {
					roman += val
				} else {
					// fmt.Println("I have no idea what you are talking about")
					roman = ""
					break
				}
			}
			ConvertData := ConvertData(roman)
			if ConvertData != 0 {
				result := ""
				for _, x := range romanSlice {
					result = result + x + " "
				}
				result += "is " + strconv.Itoa(ConvertData)
				fmt.Println(result)
				return true
			} else {
				fmt.Println("I have no idea what you are talking about")
			}
			// fmt.Println(romanSlice)
		} else {
			fmt.Println("I have no idea what you are talking about")
		}
	} else {
		fmt.Println("I have no idea what you are talking about")
	}
	return ok
}

func saveData(word string, roman string) bool {
	res := ConvertData(roman)
	if res != 0 {
		aliasData[word] = roman
		return true
	} else {
		return false
	}
}

func getData(word string) (string, bool) {
	res := ""
	if val, ok := aliasData[word]; ok {
		res = val
		return res, true
	} else {
		return res, false
	}
}

func saveCreditData(credit string, value float64) bool {
	creditData[credit] = value
	return true
}

func getCreditData(credit string) (float64, string) {
	res := 0.0
	message := ""
	if val, ok := creditData[credit]; ok {
		res = val
		message = "Success get credit data"
		return res, message
	} else {
		message = "credit data not found"
		return res, message
	}
}

// ConvertData - function for convert roman number to integer
func ConvertData(word string) int {
	lenWord := len(word)
	if lenWord == 0 {
		return 0
	}
	for _, x := range word {
		count := strings.Count(word, string(x))
		if count > 3 {
			return 0
		} else if count == 3 {
			if string(x) == "I" || string(x) == "X" || string(x) == "C" {

			} else {
				return 0
			}
		}
	}
	sum := mapsValue[string(word[lenWord-1])]
	for x := lenWord - 1; x > 0; x-- {
		data := mapsValue[string(word[x])]
		prev := mapsValue[string(word[x-1])]
		if prev >= data {
			sum += prev
		} else {
			sum -= prev
		}
	}
	return sum
}
