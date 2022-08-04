package main

import (
		  //"RED-copy/redread"
		  //"RED-copy/redwrite"
		  "os"
		  "log"
		  "fmt"
)
func main() {

		  if len(os.Args[1:]) < 1 {
					 fmt.Println("Args please. One of either 'read' &  filename to red copy or 'write' 2 file names.")
		  } else if os.Args[1] == "write" {
					 file := OpenFile(os.Args[2])
					 fmt.Println(file)
					 //fmt.Println(len(file))
					 fileBits := ByteToBit(file)
					 //fmt.Println(fileBits)
					 //fmt.Println(len(fileBits))
					 WriteFile(os.Args[2] + ".RED", BitToByte(MultiplyBits(fileBits, 8)))
		  } else if os.Args[1] == "read" {
					 file := OpenFile(os.Args[2])
					 fileBits := ByteToBit(file)
					 WriteFile(os.Args[3], BitToByte(DevideBits(fileBits, 8)))
		  }	else {
					 fmt.Println("Not those Args please. One of either 'read' &  filename to red copy or 'write' 2 file names.")
		  }
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
		  var bitSlice []bit

		  for i := 0; i < length; i++ {
					 bitSlice = append(bitSlice, BreakdownByte(data[i])...)
		  }
		  return bitSlice
}

// Converts a bit slice to a byte slice, placing each bit in a bite of width 8 (currenyl hard coded)
func BitToByte(data []bit) []byte {
		  length := len(data)
		  var (
					 byteSlice []byte
					 singleByte byte
					 i int
		  )

		  for i < length {
					 for j := 0; j < 8; j++ {//Maybe I can reduce insted of double loop
								singleByte = singleByte | ((byte(data[i + j])) & byte(Pow(2, 7-j)))
					 }
					 byteSlice = append(byteSlice, singleByte)
					 i += 8
					 singleByte = 0
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
		  var redByte []bit
		  for i := 0; i < 8; i++ {
					 if 0 < (bitMask & byte(2^(7-i))) & input {
								redByte = append(redByte, one)
					 } else {
								redByte = append(redByte , zero)
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
func REDCopy(filename string, multiplicity uint) {
		  WriteFile(filename + ".red", append(RawBitToByte(MultiplyBits(ByteToBit(OpenFile(filename)), multiplicity)), byte(multiplicity * 8))) 
		  // Wanted to try composing it... is there some other syntax for this?...
}

func Homogeneous[E comparable](slice []E) bool {
		  length := len(slice)
		  ref := slice[0]
		  for i := 0; i < length; i++{
					 if ref != slice[i] {
								return false
					 }
		  }

		  return true
}

func DominantElem[E comparable](slice []E) E {
		  length := len(slice)
		  collect := make(map[E]int)
		  for i := 0; i < length; i++{
					 collect[slice[i]]++
		  }

		  var (
					 domElem E
					 domOccs int
		  )

		  for k, v := range collect {
					 if v > domOccs {
								domElem = k
					 }
		  }

		  return domElem
}

func IsCorrupt(data []bit, multiplicity uint) int {
		  length := uint(len(data))

		  for i := uint(0); i < length; i += multiplicity {
					 if !Homogeneous(data[i:i+multiplicity]) {
								return int(i)
					 }
		  }
		  return -1
}

func CorruptionAwareBitDevide(redData []bit, multiplicity uint) []bit { // better error system
		  length := uint(len(redData))
		  var clearData []bit

		  for i := uint(0); i < length; i += multiplicity {
					 clearData = append(clearData, DominantElem(redData[i:i+multiplicity]))
		  }
		  return clearData
		  
}

// Reducess a redundant (RED) series of bits to clear data (the origianl file).
//This is an unsafe method which assumes no corruption and that you remeber the multiplicity correctly
func DevideBits(redData []bit, multiplicity uint) []bit {
		  var clearData []bit
		  length := uint(len(redData))
		  i := uint(0)

		  for i < length {
					 clearData = append(clearData, redData[i])
					 i += multiplicity
		  }

		  return clearData
}
/*
// Determin the multiplicity of a RED file
func QualifyMulitplicity(data []bit) int {
		  passed := false
		  var mult int

		  for passed == false {
					 // figure out the smallest
		  }

}

*/

func Pow(num byte, pow int) byte {
		  if pow == 0 {return 1}
		  if pow == 1 {return num}
		  if pow == 2 {return num * num}
		  return Pow(num, pow - 1) * num 
}
