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
eg: 0001 1111 1000 -- shift 1 bit --> 0011 1111 0000
```go
func ShiftLeft(input []byte, shiftNum int) (result []byte, leftMostCarryFlag bool)
```
shiftNum is the number of bits to be shifted

### func ShiftRight

ShiftRight performs right(>>) bit-shifting on entire byte array
```go
func ShiftRight(input []byte, shiftNum int) (result []byte, rightMostCarryFlag bool)
```
shiftNum is the number of bits to be shifted

### func SubBits

SubBits extracts bits in the form of []byte from an input byte array
```go
func SubBits(input []byte, startBitPos int, numOfBits int) (result []byte, resultPtr *[]byte, err error)
```
The startBitPos is starting from 1 (inclusive)<br />
numOfBits is the number of bits to be extracted, if length is set to 0, then all bits starting from startBitPos would be returned

### func Bool

Bool converts bit at startBitPos to boolean
```go
func Bool(input []byte, startBitPos int) (result bool, resultPtr *bool, err error)
```

### func Uint64

Uint64 converts []byte into unsigned 64 bits integer using 64 bits starting from the startBitPos
```go
func Uint64(input []byte, startBitPos int) (result uint64, resultPtr *uint64, err error)
```

### func Uint32

Uint32 converts []byte into unsigned 32 bits integer using 32 bits starting from the startBitPos
```go
func Uint32(input []byte, startBitPos int) (result uint32, resultPtr *uint32, err error)
```

### func Uint16

Uint16 converts []byte into unsigned 16 bits integer using 16 bits starting from the startBitPos
```go
func Uint16(input []byte, startBitPos int) (result uint16, resultPtr *uint16, err error)
```

### func Uint8

Uint8 converts []byte into unsigned 8 bits integer using 8 bits starting from the startBitPos
```go
func Uint8(input []byte, startBitPos int) (result uint8, resultPtr *uint8, err error)
```

### func Int

Int converts []byte into int using bits starting from the startBitPos
```go
func Int(input []byte, startBitPos int) (result int, resultPtr *int, err error)
```

### func Byte

Byte extracts single byte from input using bits starting from the startBitPos
```go
func Byte(input []byte, startBitPos int) (result byte, resultPtr *byte, err error)
```

### func HexString

HexString converts []byte into string
```go
func HexString(input []byte, startBitPos int, numOfBits int) (result string, resultPtr *string, err error)
```

### func String

String converts []byte into string
```go
func String(input []byte, startBitPos int, numOfBits int) (result string, resultPtr *string, err error)
```


