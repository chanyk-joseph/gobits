# gobits
Package gobits provides a set of functions to extract bits from input byte array

## Usage

### func Len

```go
func Len(input []byte) int
```
Len returns the number of bits of the input bytes array

### func ShiftLeft

```go
func ShiftLeft(input []byte, shiftNum int) (result []byte, leftMostCarryFlag bool)
```
ShiftLeft performs left(<<) bit-shifting on entire byte array
shiftNum is the number of bits to be shifted

