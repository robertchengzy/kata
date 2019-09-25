package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"kata/utils"
)

func main() {
	/*fmt.Println("good good study, day day up")
	fmt.Println(HashCode("1234567889"))

	fmt.Println(float64(1) / float64(3) * 10000)

	data := time.Now().Format("2006-01-02T15:04:05Z")

	fmt.Println(data)

	fmt.Println(time.Now().Add(time.Hour * 37))
	*/
	prikey, _ := ioutil.ReadFile("prikey_4gkaer.pem")
	pubkey, _ := ioutil.ReadFile("pubkey_4gkaer.pem")

	res, err := utils.RsaEncryptNoPadding("834882283", pubkey)
	fmt.Println(res, err)

	ress, err := utils.RsaDecryptNoPadding(res, prikey)

	fmt.Println(ress, err)

	/*batch := int32(math.Ceil(float64(5000) / 3000))
	fmt.Println(batch)

	var keywords []string
	keywords = append(keywords, "123", "1234")
	fmt.Println(keywords)*/
}

// 获取hashcode
func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
