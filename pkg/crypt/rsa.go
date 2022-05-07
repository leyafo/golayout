package crypt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const maxEncryptLen = 39

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
// var privateKey1 = []byte(`

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
// var publicKey1 = []byte(`

var (
	privateKey *rsa.PrivateKey
	publicKey  rsa.PublicKey
)

func SetKey(privateKeyBytes []byte) error {
	var err error
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	publicKey = privateKey.PublicKey

	return nil
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &publicKey, origData)
}

func RsaEncryptBase64(origData []byte) ([]byte, error) {
	ciphertext, err := RsaEncrypt(origData)
	if err != nil {
		return nil,  err
	}
	var data []byte
	data = make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(data, ciphertext)
	return data, nil
}

func DecryptBase64(ciphertext []byte) ([]byte, error) {
	var data []byte
	data = make([]byte, base64.StdEncoding.DecodedLen(len(ciphertext)))
	n, err := base64.StdEncoding.Decode(data, ciphertext)
	if err != nil {
		fmt.Println(n, err)
		return nil, err
	}
	data = data[:n]

	// fmt.Printf("%s\n", data)
	plainText, err := RsaDecrypt(data)
	if err != nil {
		return nil, err
	}
	return plainText, err
}

func DecryptBase64Str(ciphertext string)([]byte, error){
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil{
		return nil, err 
	}
	plainText, err := RsaDecrypt(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", plainText)
	return plainText, err
}


// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

func SplitEncrypt(plainText []byte)([]byte, error){
	var (
		err error
		cipherText []byte
	)
	result := new(bytes.Buffer)
	for i:=0; i<len(plainText); i+= maxEncryptLen{
		if i+maxEncryptLen > len(plainText){
			cipherText, err = RsaEncryptBase64(plainText[i:])
		}else{
			cipherText, err = RsaEncryptBase64(plainText[i:i+maxEncryptLen])
		}

		if err != nil {
			return nil, err
		}

		result.Write(cipherText)
		result.Write([]byte(";"))
	}
	return result.Bytes(), nil
}