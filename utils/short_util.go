package utils

import "github.com/spaolacci/murmur3"

var ntab = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenShortUrl(oriUrl string) (shortUrl string) {
	hash64 := murmur3.Sum32([]byte(oriUrl))
	return convert(hash64, 62)
}

func convert(num uint32, radix uint32) string {
	r := make([]byte, 0)
	for num != 0 {
		d := num / radix
		m := num % radix
		r = append(r, ntab[m])
		num = d
	}
	// 把结果反转过来
	for i := 0; i < len(r)/2; i++ {
		oi := len(r) - 1 - i
		r[i], r[oi] = r[oi], r[i]
	}
	return string(r)
}
