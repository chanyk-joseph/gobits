package gobits

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math"
)

//Len returns the number of bits of the input bytes array
func Len(input []byte) int {
	return len(input) * 8
}

//ShiftLeft performs left(<<) bit-shifting on entire byte array
//shiftNum is the number of bits to be shifted
func ShiftLeft(input []byte, shiftNum int) (result []byte, leftMostCarryFlag bool) {
	if shiftNum >= 1 {
		result = make([]byte, len(input))
		copy(result, input)

		for i := 0; i < len(result); i++ {
			carryFlag := ((result[i] & 0x80) == 0x80)

			if i > 0 && carryFlag {
				result[i-1] |= 0x01
			} else if i == 0 {
				leftMostCarryFlag = carryFlag
			}

			result[i] <<= 1
		}

		if shiftNum == 1 {
			return result, leftMostCarryFlag
		}
		return ShiftLeft(result, shiftNum-1)
	}

	return input, false
}

//ShiftRight performs right(>>) bit-shifting on entire byte array
//shiftNum is the number of bits to be shifted
func ShiftRight(input []byte, shiftNum int) (result []byte, rightMostCarryFlag bool) {
	if shiftNum >= 1 {
		result = make([]byte, len(input))
		copy(result, input)
		for i := len(result) - 1; i >= 0; i-- {
			carryFlag := ((result[i] & 0x01) == 0x01)

			if i < len(result)-1 && carryFlag {
				result[i+1] |= 0x80
			} else if i == len(result)-1 {
				rightMostCarryFlag = carryFlag
			}

			result[i] >>= 1
		}

		if shiftNum == 1 {
			return result, rightMostCarryFlag
		}
		return ShiftRight(result, shiftNum-1)
	}

	return input, false
}

//SubBits extracts bits in the form of []byte from an input byte array
//The startBitPos is starting from 1 (inclusive)
//numOfBits is the number of bits to be extracted, if length is set to 0, then all bits starting from startBitPos would be returned
func SubBits(input []byte, startBitPos int, numOfBits int) (result []byte, resultPtr *[]byte, err error) {
	if numOfBits == 0 {
		numOfBits = Len(input) - startBitPos + 1
	}
	if startBitPos <= 0 {
		return nil, nil, errors.New("startBitPos must be > 0")
	}
	endBitPos := startBitPos + numOfBits - 1
	if endBitPos > Len(input) {
		return nil, nil, errors.New("Bit index out of bound")
	}

	startByteIndex := int((startBitPos - 1) / 8)
	startBitIndexInByte := (startBitPos - 1) % 8

	endByteIndex := int((endBitPos - 1) / 8)
	endBitIndexInByte := (endBitPos - 1) % 8

	result = make([]byte, endByteIndex-startByteIndex+1)
	for i := startByteIndex; i <= endByteIndex; i++ {
		result[i-startByteIndex] = input[i]
	}

	if endBitIndexInByte < 7 {
		result[len(result)-1] = result[len(result)-1] >> uint(7-endBitIndexInByte)
		result[len(result)-1] = result[len(result)-1] << uint(7-endBitIndexInByte)
	}
	result, _ = ShiftLeft(result, startBitIndexInByte)

	numOfBytesNeeded := int(math.Ceil(float64(endBitPos-startBitPos+1) / float64(8)))
	finalResult := make([]byte, numOfBytesNeeded)
	copy(finalResult, result)

	return finalResult, &finalResult, nil
}

//Bool converts bit at startBitPos to boolean
func Bool(input []byte, startBitPos int) (result bool, resultPtr *bool, err error) {
	if Len(input)-startBitPos+1 < 1 {
		return false, nil, errors.New("Input is less than 1 bit")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 1)
	result = (tmpArr[0] & 0x80) != 0
	return result, &result, err
}

//Uint64 converts []byte into unsigned 64 bits integer using bits starting from the startBitPos
func Uint64(input []byte, startBitPos int) (result uint64, resultPtr *uint64, err error) {
	if Len(input)-startBitPos+1 < 64 {
		return 0, nil, errors.New("Input is less than 64 bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 64)
	result = binary.BigEndian.Uint64(tmpArr)

	return result, &result, err
}

//Uint32 converts []byte into unsigned 32 bits integer using bits starting from the startBitPos
func Uint32(input []byte, startBitPos int) (result uint32, resultPtr *uint32, err error) {
	if Len(input)-startBitPos+1 < 32 {
		return 0, nil, errors.New("Input is less than 32 bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 32)
	result = binary.BigEndian.Uint32(tmpArr)

	return result, &result, err
}

//Uint16 converts []byte into unsigned 16 bits integer using bits starting from the startBitPos
func Uint16(input []byte, startBitPos int) (result uint16, resultPtr *uint16, err error) {
	if Len(input)-startBitPos+1 < 16 {
		return 0, nil, errors.New("Input is less than 16 bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 16)
	result = binary.BigEndian.Uint16(tmpArr)

	return result, &result, err
}

//Uint8 converts []byte into unsigned 8 bits integer using bits starting from the startBitPos
func Uint8(input []byte, startBitPos int) (result uint8, resultPtr *uint8, err error) {
	if Len(input)-startBitPos+1 < 8 {
		return 0, nil, errors.New("Input is less than 8 bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 8)
	result = tmpArr[0]

	return result, &result, err
}

//Int converts []byte into int using bits starting from the startBitPos
func Int(input []byte, startBitPos int) (result int, resultPtr *int, err error) {
	i64, _, err := Uint64(input, startBitPos)
	result = int(i64)

	return result, &result, err
}

//Byte extracts single byte from input using bits starting from the startBitPos
func Byte(input []byte, startBitPos int) (result byte, resultPtr *byte, err error) {
	if Len(input)-startBitPos+1 < 8 {
		return 0, nil, errors.New("Input is less than 8 bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, 8)
	result = tmpArr[0]

	return result, &result, err
}

//HexString converts []byte into string
func HexString(input []byte, startBitPos int, numOfBits int) (result string, resultPtr *string, err error) {
	if Len(input)-startBitPos+1 < numOfBits {
		return "", nil, errors.New("Input is less than " + string(numOfBits) + " bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, numOfBits)
	result = hex.EncodeToString(tmpArr)
	return result, &result, err
}

//String converts []byte into string
func String(input []byte, startBitPos int, numOfBits int) (result string, resultPtr *string, err error) {
	if Len(input)-startBitPos+1 < numOfBits {
		return "", nil, errors.New("Input is less than " + string(numOfBits) + " bits")
	}

	tmpArr, _, err := SubBits(input, startBitPos, numOfBits)
	result = string(tmpArr)
	return result, &result, err
}
