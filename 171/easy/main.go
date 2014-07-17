package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// go run main.go "FF 81 BD A5 A5 BD 81 FF"
// go run main.go "AA 55 AA 55 AA 55 AA 55"
// go run main.go "3E 7F FC F8 F8 FC 7F 3E"
// go run main.go "93 93 93 F3 F3 93 93 93"

func main() {
	draw_char := "x"
	directive := strings.Fields(os.Args[1])

	for _, v := range directive {
		mask := uint8(0x80)
		binary, err := hex.DecodeString(v)
		if err != nil {
			return
		}
		int_p := uint8(binary[0])
		for i := 0; i <= 8; i++ {
			p := int_p & mask
			if p == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf(draw_char)
			}
			mask = mask >> 1
		}
		fmt.Printf("\n")

	}
}
