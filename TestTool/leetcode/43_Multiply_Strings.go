package main

func main() {
	
}

/**
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var result []byte
	temp, carry,length := 0, 0, len(num1) + len(num2) - 2

	for i := len(num1) -1; i>=0; i-- {
		for j := len(num2)-1;j>=0;j-- {
			temp = int(num1[i] - '0') * int(num2[j] - '0') + carry

			if len(result) > length - i - j {
				result[]
			} else {
				carry = temp/10
				result = append(result, byte(temp%10 - '0'))
			}
		}
	}
}