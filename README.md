# gobits
Package gobits provides a set of functions to extract bits from input byte array

## Usage

### func Len

Len returns the number of bits of the input bytes array
```go
func Len(input []byte) int
```

### func ShiftLeft

ShiftLeft performs left(<<) bit-shifting on entire byte array
```go
func ShiftLeft(input []byte, shiftNum int) (result []byte, leftMostCarryFlag bool)
```
shiftNum is the number of bits to be shifted

### func ShiftRight

ShiftRight performs right(>>) bit-shifting on entire byte array
```go
func ShiftRight(input []byte, shiftNum int) (result []byte, leftMostCarryFlag bool)
```
shiftNum is the number of bits to be shifted

### func SubBits

SubBits extracts bits in the form of []byte from an input byte array
```go
func SubBits(input []byte, startBitPos int, numOfBits int) (result []byte, resultPtr *[]byte, err error)
```
The startBitPos is starting from 1 (inclusive)<br />
numOfBits is the number of bits to be extracted, if length is set to 0, then all bits starting from startBitPos would be returned