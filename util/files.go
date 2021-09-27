package util

import (
	"encoding/binary"
	"log"
	"os"
)

func ReadBytes(file *os.File, qtdBytes int) []byte {
	bytes := make([]byte, qtdBytes)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func ReadIntegers(file *os.File, num int) []T {
	arr := make([]T, num)
	for i := 0; i < num; i++ {
		arr[i] = T(binary.LittleEndian.Uint32(ReadBytes(file, 4)))
	}
	return arr
}
