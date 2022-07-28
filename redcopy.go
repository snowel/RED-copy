package main
/*
import (
		  "RED-copy/redread"
		  "RED-copy/redwrite"
)
*/
func main() {

}

// Multiply a single byte into an array of 8 bytes, each correponding to the 
func OctoplyByte(input byte) [8]byte {
		  bitMask := byte(1)
		  var redByte [8]byte
		  for i := 0; i < 8; i++ {
					 redByte[i] = bitMask & (input << (i+1))
		  }
		  return redByte
}

// Multiply a s
func OctoplyFile(file []byte) []byte {
		  var redFile []byte
		  length := len(file)

		  for i := 0; i < length; i++ {
					 redByte := OctoplyByte(file[i])
					 redFile = append(redfile, redByte...)
		  }
		  return redFile
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

func REDByteEval(Red []byte, ply int) {

}

func CorruptionAwareCollimator(redFile []byte, ply int) []byte
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
