package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	// IDトークンから、ヘッダー・ペイロードを入手するデモ

	// 変数 idTokenには自分のIDトークンを代入
	// >- export ID_TOKEN="自分のIDトークン"をしてから go run .
	idToken := os.Getenv("ID_TOKEN")

	dataArray := strings.Split(idToken, ".")
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]
	// fmt.Println("sig", sig)

	// headerをbase64 decodeする
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64 decodeする
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 公開鍵構造体を作る
	N := "rpqQOnhPWBqtg1DIQ3cQ1ozZWsyViJvQgvg1IvKYeMlpRHTMI2ySsQeCtUftWtohH_zTg2xY-Yg09bZ5Iq-gCS8haEUVb5nFOlmteomObtS5W3S8J1laBPfABiEICAcw-Q7YhP_WXlrnzDSQqvUMl2fR2Sl9WlWp30s13OS4mJv9EE0oa8k22BNxCv5_fhH4YXcTLKdp2U4F5nnoxFS-HaAsCQlCAK7g2yEn1eNn0h14yQ3IaGVx8ilOxDZxJpAeeqVUVJ_lA-1S47mHdRjrqmfTZ5JD60IL906mVBc_uSEyIh7XPbeuECjFDtyu_4N5cHqMekfbxoSQ5iOVugMkTQ"
	E := "AQAB"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// 検証するデータ
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名をbase64 decodeする
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("vaild token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
