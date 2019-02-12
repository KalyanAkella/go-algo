package main

import (
	"fmt"
	"strconv"
	"strings"
)

const NO = "Neither"

func validIPAddress(IP string) string {
	has_dots := strings.ContainsAny(IP, ".")
	has_colons := strings.ContainsAny(IP, ":")

	if has_dots && has_colons {
		return NO
	} else if has_dots {
		return validateIPV4(IP)
	} else if has_colons {
		return validateIPV6(IP)
	} else {
		return NO
	}
}

func validateIPV4(ip string) string {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return NO
	}
	for _, part := range parts {
		if len(part) == 0 {
			return NO
		}
		if len(part) > 1 && strings.HasPrefix(part, "0") {
			return NO
		}
		if num, err := strconv.Atoi(part); err != nil {
			return NO
		} else {
			if num < 0 || num > 255 {
				return NO
			}
		}
	}
	return "IPV4"
}

func validateIPV6(ip string) string {
	return NO
}

func numWays(stairs int) int {
	if stairs == 1 {
		return 1
	}

	if stairs == 2 {
		return 2
	}

	return numWays(stairs-1) + numWays(stairs-2)
}

/*
T(N) = T(N-1) + T(N-2) + C
T(3) = T(1) + T(2) + C
T(3) = C + C + C = 3(C)

so, O(N) solution
*/

func toLower(str string) string {
	result := ""
	for i := range str {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}
		result = result + string(c)
	}
	return result
}

func compress(input string) string {
	output := ""
	counter := 1
	prevChar := byte(0)

	for i := range input {
		currChar := input[i]
		if currChar == prevChar {
			counter++
		} else {
			if prevChar > byte(0) {
				output = fmt.Sprintf("%s%s%d", output, string(prevChar), counter)
			}
			prevChar = currChar
			counter = 1
		}
	}
	output = fmt.Sprintf("%s%s%d", output, string(prevChar), counter)
	return output
}

func isSubstring(short, long string) bool {
	return strings.Contains(long, short)
}

func areRotationsBad(s1, s2 string) bool {
	index := strings.IndexByte(s2, s1[0])
	if index < 0 {
		return false
	}
	prefix := s2[0:index]
	if isSubstring(prefix, s1) {
		temp := s2[index:len(s2)] + prefix
		fmt.Println(temp)
		return temp == s1
	}
	return false
}

func areRotations(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1s1 := s1 + s1
	return isSubstring(s2, s1s1)
}

func main() {
	/*
		fmt.Println(numWays(4))
		fmt.Println(numWays(3))
		fmt.Println(toLower("aBcDeF"))
		fmt.Println(compress("aab"))
		fmt.Println(compress("a"))
		fmt.Println(compress("aabcccccaaa"))
	*/

	fmt.Println(areRotations("water", "rewat"))
}
