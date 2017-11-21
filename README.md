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