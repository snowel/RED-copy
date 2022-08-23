package main

import (
		  "os"
		  "log"
		  "fmt"
)
func main() {

		  cmd := os.Args
		  cL := len(cmd[1:])
		  if cL < 1 {
					 fmt.Println("Args please.\nTry \"redcopy h\" for help.")
		  } else if cmd[1] == "h" || cmd[1] == "help" {
					 fmt.Println("R.E.D Copy - Redundant Error-resistant Digital Copies")
					 fmt.Println("")
					 fmt.Println("redcopy write [file to copy] {multiplicity} {header}")
					 fmt.Println("redcopy read [source file] [name of the extracted file] [multiplicity")
					 fmt.Println("")
					 fmt.Println("[] => Mandatory field, {} => Optional field")
					 fmt.Println("--- Write ---")
					 fmt.Println("")
					 fmt.Println("Multiplicity -> defaults to 8.")
					 fmt.Println("Header -> Defaults to ture. Use \"false\" to set the file with a header.")
					 fmt.Println("")
					 fmt.Println("--- Read ---")
					 fmt.Println("")
					 fmt.Println("Read automatically checks the red file for a header and ignores manual multiplicity if there is a header")
					 fmt.Println("")
					 fmt.Println("")
					 fmt.Println("")
					 
		  }else if cmd[1] == "test" {
					 // --- Opened file
					 file := OpenFile(cmd[2])
					 fmt.Println("The base file is : ", len(file))
					 for i := 0; i < len(file); i++ {
								fmt.Printf("%08b", file[i])
					 }
					 fmt.Println("")

					 // -- Base file converted to bits
					 fileBits := ByteToBit(file)
					 fmt.Println("The base file bits : ", len(fileBits))

					 //
					 file = BitToByte(MultiplyBits(fileBits, 8))
					 fmt.Println(len(file))
					 for i := 0; i < len(file); i++ {
								fmt.Printf("%08b", file[i])
					 }
					 fmt.Println("")

		  } else if cmd[1] == "write" {
					 file := OpenFile(cmd[2])
					 fileBits := ByteToBit(file)

					 var mult byte
					 if cL >= 3 {
								switch cmd[3] {
								case "weak": mult = 3
								case "normal": mult = 8
								case "strong": mult = 32
								case "super": mult = 64
								case "ludicrous": mult = 254
								default: mult = 8
								}
					 } else {
								mult = 8
					 }
					 file = BitToByte(MultiplyBits(fileBits, mult))
					 if cL >= 4 && cmd[4] == "false" {
								fmt.Println("This R.E.D Copy is being written without a header byte. This is not recomended.")
					 } else {
								file = AddHeader(file, mult)
					 }
					 WriteFile(cmd[2] + ".RED", file)
		  } else if cmd[1] == "read" {
					 file := OpenFile(os.Args[2])
					 var mult byte
					 if file[0] != 0 && file[0] != 255 {
								file, mult = StripHeader(file)
					 } else if cL >= 4 {
								switch cmd[3] {
								case "weak": mult = 3
								case "normal": mult = 8
								case "strong": mult = 32
								case "super": mult = 64
								case "ludicrous": mult = 255
								default: { 
										  fmt.Println("Unknown; Multiplicity defaulted to 8.")
										  mult = 8
								}
								}
					 } else {
								mult = 8
					 }
					 fileBits := ByteToBit(file)

					 if errArea := IsCorrupt(fileBits, mult); errArea != -1 {
								fmt.Println("There's somethign wrong. Either:\n=-=-= File is corrupt, check around: ", errArea, "\n=-=-= You have the wrong multiplicity, try weak, normal, strong, super and ludicrous")
					 }
					 WriteFile(os.Args[3], BitToByte(DevideBits(fileBits, mult)))
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


// Bit - Byte conversion

// Converts a byte slice to a bit slice
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
		  )

		  i := 0
		  for i < length {
					 for j := 0; j < 8; j++ {//Maybe I can reduce insted of double loop
								singleByte = singleByte | ((byte(data[i + j])) & Pow(2, 7-j))
					 }
					 byteSlice = append(byteSlice, singleByte)
					 i += 8
					 singleByte = 0
		  }
		  return byteSlice
}

func BreakdownByte(input byte) []bit {
		  bitMask := byte(255)
		  var redByte []bit
		  for i := 0; i < 8; i++ {
					 if 0 < (bitMask & Pow(2, (7-i) )) & input {
								redByte = append(redByte, one)
					 } else {
								redByte = append(redByte , zero)
					 }
		  }
		  return redByte
}

// --- Bit slice multiplication

func MultiplyBits(data []bit, ply byte) []bit {
		  length := len(data)
		  var redSlice []bit

		  for i := 0; i < length; i++ {
					 for j := 0; j < int(ply); j++ {
								redSlice = append(redSlice, data[i])
					 }
		  }

		  return redSlice
}


func BigMultiplyBits(data []bit, ply uint) []bit {
		  length := len(data)
		  var redSlice []bit

		  for i := 0; i < length; i++ {
					 for j := uint(0); j < ply; j++ {
								redSlice = append(redSlice, data[i])
					 }
		  }

		  return redSlice
}

// Multiplicity Header - Overflow check not included

func AddHeader(data []byte, mult byte) []byte {
		  var headedData []byte
		  headedData = append(headedData, mult)
		  headedData = append(headedData, data...)

		  return headedData
}

func StripHeader(hData []byte) ([]byte, byte) {
		  var sData []byte
		  sData = append(sData, hData[1:]...)

		  return sData, hData[0]
}

// Corruption checking

// Checks that every element in a slice is the same.
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

func IsCorrupt(data []bit, multiplicity byte) int {
		  length := len(data)

		  for i := 0; i < length; i += int(multiplicity) {
					 if !Homogeneous(data[i:i+int(multiplicity)]) {
								return i
					 }
		  }
		  return -1
}

func CorruptionAwareBitDevide(redData []bit, multiplicity byte) []bit { // better error system
		  length := len(redData)
		  var clearData []bit

		  for i := 0; i < length; i += int(multiplicity) {
					 clearData = append(clearData, DominantElem(redData[i:i + int(multiplicity)]))
		  }
		  return clearData
		  
}

// Reducess a redundant (RED) series of bits to clear data (the origianl file).
//This is an unsafe method which assumes no corruption and that you remeber the multiplicity correctly
func DevideBits(redData []bit, multiplicity byte) []bit {
		  var clearData []bit
		  length := len(redData)
		  i := 0
 
		  for i < length {
					 clearData = append(clearData, redData[i])
					 i += int(multiplicity)
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
