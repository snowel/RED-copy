package main

import (
		  //"RED-copy/redread"
		  //"RED-copy/redwrite"
		  "os"
		  "log"
		  "fmt"
)
func main() {

		  file := OpenFile("redcopy.go")
		  fileBits := ByteToBit(file)
		  WriteFile("Red.txt", RawBitToByte(MultiplyBits(fileBits, 20)))
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
					 for j := uint(0); j < 8*ply; j++ {
								redSlice = append(redSlice, data[i])
					 }
		  }

		  return redSlice
}

// Multiply a single byte into an array of 8 bytes, each correponding to the 
func OctoplyByte(input byte) []byte {
		  bitMask := byte(1)
		  redByte := make([]byte, 8)
		  for i := 0; i < 8; i++ {
					 if 1 == bitMask & (input << (i+1)) {
								redByte[i] = byte(one)
					 } else {
								redByte[i] = byte(zero)
					 }
		  }
		  return redByte
}

// Multiply a file's bytes by 8.
func OctoplyFile(file []byte) []byte {
		  var redFile []byte
		  length := len(file)

		  for i := 0; i < length; i++ {
					 redByte := OctoplyByte(file[i])
					 redFile = append(redFile, redByte...)
		  }
		  return redFile
}
/*
// Eval byte
func EvalByte(val byte) int {
		  avg := byte(0)
		  bitMask := byte(1)

		  for i := 0; i < 8; i++ {
					 avg += bitMask & (val << (i+1))
		  }

		  if avg > 4 {
					 return 1
		  } else {
					 return 0
		  }
}

// 
func QuickColimator(redFile []byte, ply int) []byte{
		  var rawFile []byte
		  length := len(redFile)

		  if length % ply != 0 {
					 fmt.Println("Something is wrong with the RED file. Most likely the wrong number of ply.")
		  }

		  bitMask := byte(1)
		  for i := 0; i < length; i + ply {
					 singleByte := byte(0)
					 bitShift := ply
					 bigEnd := 0
					 
					 for bitShift > 0 {
								bitVal := byte(refFile[i + bigEnd] & bitMask)
								singleByte = (singleByte | bitVal) << bitShift
								bitShift--
								bitEnd++
					 } 
					 rawFile = append(rawFile, singleByte)
		  }
		  return rawFile
}

func CorruptionAwareCollimator(redFile []byte, ply int) []byte {
		  var rawFile []byte
		  length := len(redFile)

		  if length % ply != 0 {
					 fmt.Println("Something is wrong with the RED file. Most likely the wrong number of ply.")
		  }

		  bitMask := byte(1)
		  for i := 0; i < length; i + ply {
					 singleByte := byte(0)
					 bitShift := ply
					 bigEnd := 0
					 
					 for bitShift > 0 {
								bitVal := byte(refFile[i + bigEnd] & bitMask)
								singleByte = (singleByte | bitVal) << bitShift
								bitShift--
								bitEnd++
					 } 
					 rawFile = append(rawFile, singleByte)
		  }
		  return rawFile
}
*/
