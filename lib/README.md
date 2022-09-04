## Library Functions

#### Bit Type

The `bit` type is intended to represent a single bit, the 1 or 0 of a digital file. The `bit` type is an alias for a byte in which its enumerated values `zero` and `one` represent `byte(0)` and `byte(255)` repsectively.

#### Bit - Byte Conversion

`ByteToBit()` and `BitToByte()` take slices of one type an converte them to the other. *No mutiplication occures at this stage.*

#### Multiply Bit - Creating redundant files

`MultiplyBits()` allows the multiplication of a bit slice. The multiplicity is limited to the range of type `byte` as any greater multiplicity would overflow in the single byte header.

`BigMultiplyBits(data []bit, multiplicity uint) []bit` allows the multiplication of a bit slice by an arbitraty unsinged integer values. Not compatible with headers.

#### Deviding Bits - Reducing redundant files to their original bit sequences

`DevideBits` allows the devision of a R.E.D file of known multiplicity. *This is not corruption aware. The first bit of each multiplicity segement is used.*

`CorruptionAwareBitDevide(data []bit, multiplicity byte) []bit` allows the devision of a R.E.D file of known multiplicity. The corruption awareness is accomplished by using the most frequent bit value in a give multiplicity segment.

#### Corruption Awareness

`IsCorrupt(data []bit, multiplicity byte) int` checks is a slice of type `bit` with known multiplity has any sign of corruption. A return values of `-1` indicates no corruption, otherwise the return value shows the index of possible corruption.

