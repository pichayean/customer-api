package services

import (
	"crypto/rsa"
	"fmt"
	"log"
	"macus/common"
	"macus/model"
	"macus/model/apperrors"
)

type TokenService interface {
	NewToken(u *model.LogIn) (*model.TokenPair, error)
}

type tokenService struct {
	PrivKey          *rsa.PrivateKey
	PubKey           *rsa.PublicKey
	IDExpirationSecs int64
}

type TSConfig struct {
	PrivKey          *rsa.PrivateKey
	PubKey           *rsa.PublicKey
	IDExpirationSecs int64
}

func NewTokenService(c *TSConfig) TokenService {
	return &tokenService{
		PrivKey:          c.PrivKey,
		PubKey:           c.PubKey,
		IDExpirationSecs: c.IDExpirationSecs,
	}
}

func (s *tokenService) NewToken(u *model.LogIn) (*model.TokenPair, error) {

	fmt.Println("sss")
	fmt.Println(s.PrivKey)
	idToken, err := common.GenerateIDToken(&u.UserName, s.PrivKey, s.IDExpirationSecs)

	fmt.Println("sss")
	if err != nil {
		log.Printf("Error generating idToken for uname: %v. Error: %v\n", u.UserName, err.Error())
		return nil, apperrors.NewInternal()
	}

	return &model.TokenPair{
		IDToken: model.IDToken{SS: idToken},
	}, nil
}

// https://kentakodashima.medium.com/generate-pem-keys-with-openssl-on-macos-ecac55791373
