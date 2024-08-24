package functions

import (
	"os"
	"strings"
)

func FS(InputFile, text string) string {
	D.Result = ""
	var sep string

	if InputFile == "standard" || InputFile == "shadow" {
		sep = "\n"
	} else {
		sep = "\r\n"
	}

	InputFile = "../static/" + InputFile + ".txt"

	data, err := os.ReadFile(InputFile)
	if err != nil {
		return "ERORR"
	}
	// slice hold the splited banner file
	slice := RemoveEmptyStrings(strings.Split(string(data), sep))
	// sliceArgs holds the splited text
	slicedArgs := strings.Split(text, "\n")
	// ranging over runes or the input text
	for _, word := range slicedArgs {
		if word != "" {
			// printing the words ascii
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						return "ERORR"
					} else {
						start := int(char-32)*8 + i
						D.Result += slice[start]
					}
				}
				D.Result += "\n"
			}
		} else {
			D.Result += "\n"
		}
	}

	if IsAllNewLines(D.Result) {
		D.Result = D.Result[1:]
	}
	return D.Result
}

func RemoveEmptyStrings(slice []string) []string {
	var temp []string
	for i := range slice {
		if slice[i] != "" {
			temp = append(temp, slice[i])
		}
	}
	return temp
}

func IsAllNewLines(str string) bool {
	for _, char := range str {
		if char != '\n' {
			return false
		}
	}
	return true
}
