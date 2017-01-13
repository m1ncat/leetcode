package main

import (
	"fmt"
	"reflect"
)

func isPalindrome(arr []byte) bool {
	length := len(arr)
	var firstPart []byte
	var lastPart []byte

	if length%2 == 0 {
		for i := 0; i < length/2; i++ {
			firstPart = append(firstPart, arr[i])
		}

		for j := length - 1; j >= length/2; j-- {
			lastPart = append(lastPart, arr[j])
		}
	} else {
		for i := 0; i < length/2; i++ {
			firstPart = append(firstPart, arr[i])
		}

		for j := length - 1; j > length/2; j-- {
			lastPart = append(lastPart, arr[j])
		}
	}

	var flag bool
	if reflect.DeepEqual(firstPart, lastPart) {
		flag = true
	} else {
		flag = false
	}

	return flag
}

func longestPalindrome(s string) string {
	var retString string

	if len(s) == 1 || len(s) == 0 {
		retString = s
	}

	arr := []byte(s)
	length := len(arr)

	start := 0
	end := 0
	refLength := 0
	refCount := 0
	finalStart := 0
	finalEnd := 0

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			flag := isPalindrome(arr[i : j+1])
			if flag && refCount == 0 {
				refLength = j - i + 1
				start = i
				end = j
			}

			if flag && refCount != 0 {
				compLength := j - i + 1
				if compLength > refLength {
					refLength = compLength
					finalStart = i
					finalEnd = j
					start = i
					end = j
				} else {
					finalStart = start
					finalEnd = end
				}
			}
		}
		if refLength != 0 {
			refCount += 1
		}
	}

	if finalStart == 0 && finalEnd == 0 {
		finalStart = start
		finalEnd = end
	}

	retString = string(arr[finalStart : finalEnd+1])
	return retString
}

func main() {
	test := "jrjnbctoqgzimtoklkxcknwmhiztomaofwwzjnhrijwkgmwwuazcowskjhitejnvtblqyepxispasrgvgzqlvrmvhxusiqqzzibcyhpnruhrgbzsmlsuacwptmzxuewnjzmwxbdzqyvsjzxiecsnkdibudtvthzlizralpaowsbakzconeuwwpsqynaxqmgngzpovauxsqgypinywwtmekzhhlzaeatbzryreuttgwfqmmpeywtvpssznkwhzuqewuqtfuflttjcxrhwexvtxjihunpywerkktbvlsyomkxuwrqqmbmzjbfytdddnkasmdyukawrzrnhdmaefzltddipcrhuchvdcoegamlfifzistnplqabtazunlelslicrkuuhosoyduhootlwsbtxautewkvnvlbtixkmxhngidxecehslqjpcdrtlqswmyghmwlttjecvbueswsixoxmymcepbmuwtzanmvujmalyghzkvtoxynyusbpzpolaplsgrunpfgdbbtvtkahqmmlbxzcfznvhxsiytlsxmmtqiudyjlnbkzvtbqdsknsrknsykqzucevgmmcoanilsyyklpbxqosoquolvytefhvozwtwcrmbnyijbammlzrgalrymyfpysbqpjwzirsfknnyseiujadovngogvptphuyzkrwgjqwdhtvgxnmxuheofplizpxijfytfabx"
	ret := longestPalindrome(test)
	fmt.Println(ret)
}
