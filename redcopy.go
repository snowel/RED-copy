package main

import (
		  //"RED-copy/redread"
		  //"RED-copy/redwrite"
		  "os"
		  "log"
		  "fmt"
)
func main() {

		  file := OpenFile("roof")
		  fileBits := ByteToBit(file)
		  WriteFile("superroof", RawBitToByte(MultiplyBits(fileBits, 20)))
}

func OpenFile(filename string) []byte {
		  f, ok := os.ReadFile(filename)
		  if ok != nil {
					 log.Fatal(ok)
		  }
		  return f
}

func WriteFile(filename string, data []byte) {
		  ok := os.WriteFile(filename, data, 0666)
		  if ok != nil {
					 log.Fatal(ok)
		  }
}

type bit byte
const (
		  zero bit = 0
		  one bit = 255
)

func ByteToBit(data []byte) []bit {
		  length := len(data)
		  bitSlice := make([]bit, length)

		  for i := 0; i < length; i++ {
					 bitSlice = append(bitSlice, BreakdownByte(data[i])...)
					 fmt.Println(BreakdownByte(data[i]))
		  }
		  return bitSlice
}

// Converts a bit slice to a byte slice, placing each bit in a bite of width 8 (currenyl hard coded)
func BitToByte(data []bit) []byte {
		  length := len(data)
		  var byteSlice []byte

		  for i := 0; i < length; i += 8 {
					 var singleByte byte
					 for j := 0; j < 8; j++ {//Maybe I can reduce insted of double loop
								// For each innner loop, the byte is ored to the val of the bit masked to the position in the inner loop. 
								singleByte = singleByte | (byte(data[i + j]) & byte(2^(7-j)))
					 }
					 byteSlice = append(byteSlice, singleByte)
		  }
		  return byteSlice
}

// Converts a bit slice to a byte slice, but treats each bit as a whole byte. This is 8x redundancy.
func RawBitToByte(data []bit) []byte {
		  length := len(data)
		  var byteSlice []byte

		  for i := 0; i < length; i++ {
					 byteSlice = append(byteSlice, byte(data[i]))
		  }
		  return byteSlice
}

func BreakdownByte(input byte) []bit {
		  bitMask := byte(255)
		  redByte := make([]bit, 8)
		  for i := 0; i < 8; i++ {
					 if 0 < (bitMask & byte(2^(7-i))) & input {
								redByte[i] = one
					 } else {
								redByte[i] = zero
					 }
		  }
		  return redByte
}

// --- Bit slice multiply

func MultiplyBits(data []bit, ply uint) []bit {
		  length := len(data)
		  var redSlice []bit

		  for i := 0; i < length; i++ {
					 for j := uint(0); j < ply; j++ {
								redSlice = append(redSlice, data[i])
					 }
		  }

		  return redSlice
}

// Pre-built operations

// Appends a filan byte which shows the multiplicity
func REDCopy(filename string, mutiplicity uint) {
		  WriteFile(append(RawBitToByte(MultiplyBits(ByteToBit(OpenFile(filename)), multiplicity)), (multiplicity * 8))) 
		  // Wanted to try composing it... is there some other syntax for this?...
}

func Homogeneous[E comparable](slice []E) bool {
		  length := len(slice)
		  ref := slice[0]
		  for i := 0; i < length; i++{
					 if i != slice[i] {
								return false
					 }
		  }

		  return true
}

func IsCorrupt(data []bit, multiplicity uint) int {
		  length := len(data)

		  for i := 0; i < length; i += mulitplicity {
					 if !Homogenenous(data[i:i+multiplicity]) {
								return i
					 }
		  }
		  return -1
}

// Reducess a redundant (RED) series of bits to clear data (the origianl file).
//This is an unsafe method which assumes no corruption and that you remeber the multiplicity correctly
func DevideBits(redData []Bits, multiplicity uint) []bits {
		  var clearData []bit
		  length := len(redData)
		  i := uint(0)

		  for i < length {
					 clearData = append(clearData, redData[i])
					 i += multiplicity
		  }

		  return clearData
}

// Determin the multiplicity of a RED file
func QualifyMulitplicity(data []bit) int {

}
