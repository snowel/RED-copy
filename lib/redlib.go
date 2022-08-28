package redlib

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
