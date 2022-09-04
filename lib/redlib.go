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
					 bitSlice = append(bitSlice, breakdownByte(data[i])...)
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
								singleByte = singleByte | ((byte(data[i + j])) & pow(2, 7-j))
					 }
					 byteSlice = append(byteSlice, singleByte)
					 i += 8
					 singleByte = 0
		  }
		  return byteSlice
}

// Conversion helpers

func breakdownByte(input byte) []bit {
		  bitMask := byte(255)
		  var redByte []bit
		  for i := 0; i < 8; i++ {
					 if 0 < (bitMask & pow(2, (7-i) )) & input {
								redByte = append(redByte, one)
					 } else {
								redByte = append(redByte , zero)
					 }
		  }
		  return redByte
}

// --- Bit slice multiplication

// Multiplies a bit slice.
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


// Multiplies a bit slice without size restrition.
// *** Not compatible with red header functions.
// TODO int rather than Uint to be compatible with qualifying multiplicity?
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

func CorruptionAwareBitDevide(redData []bit, multiplicity byte) []bit { 
		  length := len(redData)
		  var clearData []bit

		  for i := 0; i < length; i += int(multiplicity) {
					 clearData = append(clearData, DominantElem(redData[i:i + int(multiplicity)]))
		  }
		  return clearData
		  
}

// Bit slice devide - RED files to clear files

// Reducess a redundant (RED) series of bits to clear data (the origianl file).
// *** This is an unsafe method which assumes no corruption and that you remeber the multiplicity correctly.
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


// Multiplicity guessing

// Determin the multiplicity of a RED file.
// Return val of -1 indicates no conffident guess can be made.
// TODO - add multiple negative returns to specify which test fails.
func QualifyMulitplicity(data []bit) int {
		  length := len(data)
		  
		  // collect the lengths of all the sequences of same bits
		  sequ := sequenceLens(data, length)
		  
		  // take the shortest as the possible mult
		  candidate := shortestSequ(sequ)

		  // test the various required properties
		  passed := true
		  
		  if length % candidate != 0 {passed = false}
		  if !sequMultiple(sequ, candidate) {passed = false}

		  // return
		  if passed {
					 return candidate
		  } else {
					 return -1
		  }

}

// Discreet helpers


// Collect the lengths of all the sequences of same bits.
func sequenceLens(data []bit, length int) []int {
		  sequ := make([]int, 0)
		  travel := data[0]
		  counter := 1
		  for i := 1; i < length; i++ {
					 if data[i] == travel {
								counter++
					 } else {
								sequ = append(sequ, counter)

								counter = 1
								travel = data[i]
					 }
		  }
		  return sequ
}

func shortestSequ(sequ []int) int {
		  travel := sequ[0]

		  for _, val := range sequ {
					 if travel > val {
								travel = val
					 }
		  }

		  return travel

}

// Checks if every element of a sequence is a multiple of the "unit".
func sequMultiple(sequ []int, unit int) bool {
		  for _, val := range sequ {
					 if val % unit != 0 {
								return false
					 }
		  }
		  return true
}

//SEQU TEST - TODO refac

func SqueSumTest(data []bit) bool {
		  length := len(data)
		  sequ := sequenceLens(data, length)
		  sequSum := 0
		  for _, val := range sequ {
					 sequSum += val
		  }

		  if sequSum == length {
					 return true
		  } else {
					 return false
		  }
}

// Numerical helpers

func divisible(length int) []int {
		  var absLength int
		  if length < 0 {
					 absLength = length * -1
		  } else {
					 absLength = length
		  }
		  allDiv := make([]int, 0)
		  counter := 1
		  for counter < absLength {
					 if absLength % counter == 0 {
								allDiv = append(allDiv, counter)
					 }
					 
					 counter++
		  }
		  allDiv = append(allDiv, absLength)

		  return allDiv
}

func pow(num byte, expo int) byte {
		  if expo == 0 {return 1}
		  if expo == 1 {return num}
		  if expo == 2 {return num * num}
		  return pow(num, expo - 1) * num 
}
