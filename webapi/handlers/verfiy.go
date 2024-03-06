package handlers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"gin-frame/build/utils"
	"gin-frame/webapi/model"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID        int64
	UserName      string
	// Issuer        string
	ExprieAt      int64
	jwt.RegisteredClaims
}

// func ValidTime() error {

// }

func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, _ = pem.Decode(buf)  //p, buf = pem.Decode(buf)
	key, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil || p == nil {
		return nil, errors.New("parse key error")
	}
	privateKey := key.(*rsa.PrivateKey)
	return privateKey, nil
 }
 

func CreateToken(user model.UserInfo) (string, error) {
	claim := MyCustomClaims{
		UserID:    user.Id,
		UserName:  user.UserName,
		ExprieAt:  time.Now().Add(time.Duration(utils.Config.Login.ExprieAt) * time.Hour).Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "null",        // 签发者
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(utils.Config.Login.ExprieAt) * time.Hour)),     // 过期时间
		},
	}

	key, err := parsePriKeyBytes([]byte(model.PriKey))
	if err != nil {
		return "", err
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	token, err := tk.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("signed failed: %s", err)
	}
	return token, nil
}

func VerfiyToken(token string) (*MyCustomClaims, error) {
	var claim *MyCustomClaims
	tk, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, parsePubKey)
	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	claim, ok := tk.Claims.(*MyCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claim type")
	}
	return claim, nil
}

func parsePubKey(tk *jwt.Token) (interface{}, error) {
	pub, err := parsePubKeyBytes([]byte(model.PubKey))
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func parsePubKeyBytes(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	if block == nil {
		return nil, fmt.Errorf("block nil")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := key.(*rsa.PublicKey)
	return publicKey, nil
}