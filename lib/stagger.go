package redlib

import (
		  "crypto/sha512"
		  "fmt"
)

// Redundance by headered copy append

// Header

/*
 Head structure:

 n bytes - arbitary header fingerprint
 1 byte - multiplicity
 64 bytes -  sha512 sum of the original file
*/

// The finger print used by default to find the header
var (
		  DefaultFingerprint = []byte{1,2,3}
)


// Construct a byte slice for the header
func StaggerHeaderGen(fingerprint []byte, mult byte, hashsum [sha512.Size]byte) []byte {
		  var header []byte
		  header = append(header, fingerprint...)
		  header = append(header, mult)
		  hashSlice := hashSlice(hashsum)
		  header = append(header, hashSlice...)

		  return header
}

// Staggered multiply
func StaggerFile(file []byte, mult byte, fingerprint []byte) []byte {
		  hashsum := sha512.Sum512(file)
		  header := StaggerHeaderGen(fingerprint, mult, hashsum)
		  headedFile := append(header, file...)
		  var stag []byte
		  for i := byte(0); i < mult; i++ {
					 stag = append(stag, headedFile...)
		  }

		  return stag
}

// Unstagger a file with a known fingerprint
func Unstagger(stagFile []byte, fingerprint []byte) ([]byte, int) {

		  nextSeg := nextSubslice(stagFile, fingerprint, 1)// The recursions slices the file, so we always search after the start.
		  fpLen := len(fingerprint)
		  if sliceHeadMatch(fingerprint, stagFile){//TODO redundancy here as this will be cheched twice for files with corrupt first segments. This isn't an issue most of the time as it wont be
					 hashSlice := stagFile[fpLen+1: fpLen + 65]// +1 for the mult, +64 for the length of the hashsum
					 hashsum := hashArr(hashSlice)
					 if nextSeg == -1 {// Special case, last segment
								file := stagFile[fpLen+66:]
								if hashsumVerify(file, hashsum) {
								fmt.Println(file)
										  return file, 0
								} else {
										  return nil, -1
								}
					 }

					 file := stagFile[fpLen+66:nextSeg-1]
					 if hashsumVerify(file, hashsum) {
								fmt.Println(file)
								return file, 0
					 } else {
								return Unstagger(stagFile[nextSeg:], fingerprint)
					 }
		  }

		  if nextSeg == -1 {
					 return nil, -1
		  } else {
					 return Unstagger(stagFile[nextSeg:], fingerprint)
		  } 
}

// Hashsum check the file
func hashsumVerify(file []byte, hashsum [sha512.Size]byte) bool {
		  liveHash := sha512.Sum512(file)
		  if liveHash == hashsum {
					 return true
		  } else {
					 return false
		  }
}

// Make a slice of a hashsum array
func hashSlice(hashsum [sha512.Size]byte) []byte {
		  sliceSum := make([]byte, sha512.Size)
		  for i, v := range hashsum {
					 sliceSum[i] = v
		  }
		  return sliceSum
}

// Make an array of a slice
func hashArr(hashSlice []byte) [sha512.Size]byte {
		  var arr [sha512.Size]byte
		  for i, v := range hashSlice {
					 arr[i] = v
		  }
		  return arr
}

// Check is the first n elemts of surslice match the elements of slice
func sliceHeadMatch(slice []byte, surslice[]byte) bool {
		  if len(slice) > len(surslice) {return false}
		  for i :=0; i < len(slice); i++ {
					 if slice[i] != surslice[i] {return false}
		  }
		  return true
}

// Finds the next occurence of the subslice
func nextSubslice(slice []byte, subslice []byte, start int) int {
		  for i := start; i < len(slice); i++ {
					 if sliceHeadMatch(subslice, slice[i:]) {
								return i
					 }
		  }
		  return -1
}

// Checks the if the occurences of the fingerprint add to the multiplicity.
