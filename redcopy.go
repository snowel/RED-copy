package main

import (
		  "os"
		  "log"
		  "fmt"
		  . "redcopy/lib"
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
					 // --- Open file
					 file := OpenFile(cmd[2])
					 fmt.Println("The base file is : ", len(file))
					 for i := 0; i < len(file); i++ {
								fmt.Printf("%08b", file[i])
					 }
					 fmt.Println("")

					 // -- Base file converted to []bit
					 fileBits := ByteToBit(file)
					 fmt.Println("The base file bits : ", len(fileBits))

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
								case "ludicrous": mult = 254
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
