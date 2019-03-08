package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"kata/utils"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	TxAppid   = "1258777577"
	SecretId  = "AKIDMZM0ToOLPKCNNHP6nOi098warxudxbtI"
	SecretKey = "OhqrhU6Xov0uUvJMEvTvhhXIj0JVglhY"

	VoiceMaxSize = 200 * 1024
)

var txurlPrfix = "https://aai.qcloud.com/asr/v1/" + TxAppid

func main() {
	fmt.Println("good good study, day day up")

	now := time.Now().Unix()
	expire := now + 24*60*60*30
	nonce := randInt64(1, 1000000000)
	seq := 0
	end := 1 // 是否为最后一片，最后一片语音片为 1，其余为 0
	timeout := 1000 * 5
	voiceId := createVoiceId()

	txurl := txurlPrfix + fmt.Sprintf("?end=%d&engine_model_type=16k_0&expired=%d&nonce=%d&projectid=0&res_type=0"+
		"&result_text_format=0&secretid=%s&seq=%d&source=0&sub_service_type=1&timeout=%d"+
		"&timestamp=%d&voice_format=1&voice_id=%s", end, expire, nonce, SecretId, seq, timeout, now, voiceId)

	fileData, err := os.Open("16k-1.pcm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileData.Close()
	fileBytes, err := ioutil.ReadAll(fileData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var voiceLen = int64(len(fileBytes))
	fmt.Println(voiceLen)
	if voiceLen > 200*1024 {
		voiceLen = 200 * 1024
		fileBytes = fileBytes[0 : 200*1024]
	}

	sign := createSign(txurl, SecretKey)
	header := map[string]string{
		"Host":           "aai.qcloud.com",
		"Authorization":  sign,
		"Content-Type":   "application/octet-stream",
		"Content-Length": strconv.FormatInt(voiceLen, 10),
	}

	result, err := utils.PostData(txurl, fileBytes, header)
	if err != nil {
		fmt.Printf("err [%v] [%+v]\n", err, header)
		return
	}

	fmt.Println(string(result))
}

func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

var letter = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// 订单号
func createVoiceId() string {
	return time.Now().Format("20060102150405") + letter[randInt64(1, 25)] + letter[randInt64(1, 25)]
}

func createSign(url, key string) string {
	strToBeEncoded := "POST" + url[8:]
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(strToBeEncoded))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func TxVoiceToText(voiceData []byte) {

}
