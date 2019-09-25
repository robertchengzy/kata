package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
	"strings"
)

func RsaEncryptNoPadding(cipherText string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("no public key")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	encrypted := new(big.Int)
	e := big.NewInt(int64(pub.E))
	payload := new(big.Int).SetBytes([]byte(cipherText))
	encrypted.Exp(payload, e, pub.N)
	res := base64.StdEncoding.EncodeToString(encrypted.Bytes())
	return res, nil
}

func RsaDecryptNoPadding(base64CipherText string, privateKey []byte) (string, error) {
	der, _ := pem.Decode(privateKey)
	if der == nil {
		return "", errors.New("no private key")
	}

	private, err := x509.ParsePKCS1PrivateKey(der.Bytes)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return "", err
	}

	c := new(big.Int).SetBytes(cipherText)
	plainText := c.Exp(c, private.D, private.N).Bytes()
	res := strings.Replace(string(plainText), "\u0000", "", -1)
	return res, nil
}

func RsaSignWithSha1(data string, privateKey []byte) (string, error) {
	der, _ := pem.Decode(privateKey)
	if der == nil {
		return "", errors.New("no private key")
	}

	private, err := x509.ParsePKCS1PrivateKey(der.Bytes)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, hash[:])
	if err != nil {
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}
