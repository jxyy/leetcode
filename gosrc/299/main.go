package main

import "fmt"

func getHint(secret string, guess string) string {
	var a, b = 0, 0
	var mem1, mem2 = make(map[byte]int), make(map[byte]int)
	// var mem = make(map[byte]bool)
	for i := 0; i < len(secret); i++ {
		var cs, cg = secret[i], guess[i]
		if cs == cg {
			a++
		} else {
			if v, ok := mem1[cg]; ok && v > 0 {
				b++
				mem1[cg]--
			} else {
				mem2[cg]++
			}

			if v, ok := mem2[cs]; ok && v > 0 {
				b++
				mem2[cs]--
			} else {
				mem1[cs]++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1234", "0111"))
	fmt.Println(getHint("1122", "2211"))
}
