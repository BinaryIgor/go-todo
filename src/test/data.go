package test

import "math/rand"

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345789"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		cIndx := rand.Intn(len(alphanumeric))
		b[i] = alphanumeric[cIndx]
	}
	return string(b)
}

func RandomBytes(n int) []byte {
	buff := make([]byte, n)
	_, err := rand.Read(buff)
	if err != nil {
		panic(err)
	}
	return buff
}
