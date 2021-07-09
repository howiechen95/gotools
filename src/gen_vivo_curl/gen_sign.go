package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func gen(pkgappId, pkgappkey, pkgappSecret string) {
	timestamp := time.Now().UnixNano() / 1000000
	// get new one
	signStr := strings.Trim(pkgappId+pkgappkey+strconv.FormatInt(timestamp, 10)+pkgappSecret, "")
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	paramMap := make(map[string]interface{})
	paramMap["appId"] = pkgappId
	paramMap["appKey"] = pkgappkey
	paramMap["sign"] = sign
	paramMap["timestamp"] = fmt.Sprint(timestamp)
	str := fmt.Sprintf("curl --location --request POST 'https://api-push.vivo.com.cn/message/auth' \\\n--header 'Content-Type: application/json' \\\n--data-raw '{\n    \"appId\":%s,\n    \"appKey\":\"%s\",\n    \"timestamp\":%d,\n    \"sign\":\"%s\"\n}'", pkgappId, pkgappkey, timestamp, sign)
	fmt.Println("\n\n", str, "\n\n")
	str1 := fmt.Sprintf("curl -X POST -H 'Content-Type:application/json' -d  ' {\"appId\":%s,\"appKey\":\"%s\",\"timestamp\":%d,\"sign\":\"%s\"}'  https://api-push.vivo.com.cn/message/auth", pkgappId, pkgappkey, timestamp, sign)
	fmt.Println("\n\n", str1, "\n\n")
}

func main() {
	if len(os.Args) > 3 {
		appid := os.Args[1]
		appkey := os.Args[2]
		secret := os.Args[3]

		fmt.Println("appid", appid)
		fmt.Println("appkey", appkey)
		fmt.Println("secret", secret)

		gen(appid, appkey, secret)
	}
}
