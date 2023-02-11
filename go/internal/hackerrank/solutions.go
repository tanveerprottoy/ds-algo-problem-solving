package hackerrank

import "math/big"

func ExtraLongFactorials(n int32) *big.Int {
	res := big.NewInt(1)
	res = longFactorials(n, res)
	return res
}

func longFactorials(n int32, res *big.Int) *big.Int {
	if n < 2 {
		return res
	}
	res = res.Mul(res, big.NewInt(int64(n)))
	return longFactorials(n-1, res)
}

func ExtraLongFactorialsAlt(n int32) *big.Int {
	res := big.NewInt(1)
	for n > 0 {
		res = res.Mul(res, big.NewInt(int64(n)))
		n--
	}
	return res
}

/*
	Complete the timeInWords function in the editor below.

timeInWords has the following parameter(s):

int h: the hour of the day
int m: the minutes after the hour
Returns

string: a time string as described
Input Format

The first line contains , the hours portion The second line contains , the minutes portion
*/
func TimeInWords(h, m int32) string {
	res := ""
	hours := []string{
		"0", "one", "two", "three", "four", "five", "six", "seven",
		"eight", "nine", "ten", "eleven", "twelve",
	}
	minutes := []string{
		"0", "one", "two",
		"three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
		"fourteen", "quarter", "sixteen",
		"seventeen", "eighteen",
		"nineteen", "twenty",
		"twenty one", "twenty two",
		"twnenty three", "twenty four",
		"twenty five", "twenty six",
		"twenty seven", "twenty eight",
		"twenty nine", "half",
	}
	if m == 0 {
		res = hours[h] + " o' clock"
	} else if m == 1 {
		res = minutes[m] + " minute past " +
			hours[h]
	} else if m == 15 {
		res = "quarter past " + hours[h]
	} else if m <= 29 {
		res = minutes[m] + " minutes past " +
			hours[h]
	} else if m == 30 {
		res = "half past " + hours[h]
	} else if m == 45 {
		res = "quarter to " + hours[h+1]
	} else if m <= 59 {
		res = minutes[60-m] + " minutes to " +
			hours[h+1]
	}
	return res
}

/*
	Given an array of strings of digits, try to find the occurrence of
	a given pattern of digits. In the grid and pattern arrays, each string represents a row in the grid. For example, consider the following grid:

1234567890
0987654321
1111111111
1111111111
2222222222
The pattern array is:
876543
111111
111111
*/
func GridSearch(G []string, P []string) string {
	// Write your code here
	return ""
}
