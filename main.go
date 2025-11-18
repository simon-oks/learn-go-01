package main

import "fmt"

func main() {
	var (
		b   bool
		s   string
		i   int  // par défaut int64 ou int32 suivant si votre machine est en 64 bits ou 23 bits
		u   uint // par défaut uint64 ou uint32
		u8  uint8
		i8  int8
		i16 int16
		u16 uint16
		f   float32
	)

	b = true
	s = "Simon"
	i = -15
	u = 15
	u8 = 254 // 0 - 255
	i8 = 127 // -128 - 127
	i16 = -21500
	u16 = 40000
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f)
}
