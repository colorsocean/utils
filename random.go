package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

/****************************************************************
** Random utils
********/

/*
** UUID generation
 */

func RandomUUIDBytes() []byte {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(fmt.Sprintf("RandomUUIDBytes# len:%d, err:%v", n, err))
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return uuid
}

func RandomUUID() string {
	uuid := RandomUUIDBytes()
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func RandomUUIDNoDashes() string {
	uuid := RandomUUIDBytes()
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

/*
** Random data generation from alphabet
 */

const (
	AlphabetAlphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	AlphabetAlpha        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AlphabetNumeric      = "0123456789"
)

func RandomAlphanumericString(length int) string {
	return string(RandomData(length, []byte(AlphabetAlphanumeric)))
}

func RandomAlphaString(length int) string {
	return string(RandomData(length, []byte(AlphabetAlpha)))
}

func RandomNumericString(length int) string {
	return string(RandomData(length, []byte(AlphabetNumeric)))
}

func RandomString(length int, alphabet string) string {
	return string(RandomData(length, []byte(alphabet)))
}

func RandomData(length int, alphabet []byte) []byte {
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(alphabet))
	maxrb := byte(256 - (256 % len(alphabet)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, r); err != nil {
			panic("error reading from random source: " + err.Error())
			return nil
		}
		for _, c := range r {
			if c >= maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = alphabet[c%clen]
			i++
			if i == length {
				return b
			}
		}
	}
	panic("unreachable")
	return nil
}
