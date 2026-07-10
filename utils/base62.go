package utils

import "errors"

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var base62Lookup [128]int

func init() {
	// Initialize all entries to -1 (invalid character)
	for i := range base62Lookup {
		base62Lookup[i] = -1
	}

	// Populate lookup table
	for i := 0; i < len(base62Chars); i++ {
		base62Lookup[base62Chars[i]] = i
	}
}

// EncodeBase62 converts an int64 to a Base62 string.
func EncodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	// Maximum Base62 length for int64 is 11 characters.
	var buf [11]byte
	i := len(buf)

	for num > 0 {
		i--
		buf[i] = base62Chars[num%62]
		num /= 62
	}

	return string(buf[i:])
}

// DecodeBase62 converts a Base62 string back to int64.
func DecodeBase62(encoded string) (int64, error) {
	var num int64

	for i := 0; i < len(encoded); i++ {
		c := encoded[i]

		if c >= 128 || base62Lookup[c] == -1 {
			return 0, errors.New("invalid base62 character")
		}

		num = num*62 + int64(base62Lookup[c])
	}

	return num, nil
}