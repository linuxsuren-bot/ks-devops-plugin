/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package token

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/klog"
)

const (
	DefaultIssuerName = "kubesphere"
)

type Claims struct {
	Username  string              `json:"username"`
	Groups    []string            `json:"groups,omitempty"`
	Extra     map[string][]string `json:"extra,omitempty"`
	TokenType TokenType           `json:"token_type"`
	// Currently, we are not using any field in jwt.StandardClaims
	jwt.StandardClaims
}

type jwtTokenIssuer struct {
	name   string
	secret []byte
	// Maximum time difference
	maximumClockSkew time.Duration
}

func (s *jwtTokenIssuer) VerifyWithoutClaimsValidation(tokenString string) (user.Info, TokenType, error) {
	parser := &jwt.Parser{
		SkipClaimsValidation: true,
	}
	token, err := parser.Parse(tokenString, s.keyFunc)
	klog.V(9).Infof("verify the jwt token: %s", tokenString)
	if err != nil {
		if token != nil {
			mapClaims := token.Claims.(jwt.MapClaims)

			return &user.DefaultInfo{
				Name: mapClaims["sub"].(string),
			}, "clm.TokenType", nil
		}
		return nil, "", err
	}
	return &user.DefaultInfo{}, "clm.TokenType", nil
}

func (s *jwtTokenIssuer) Verify(tokenString string) (user.Info, TokenType, error) {
	clm := &Claims{}
	// verify token signature and expiration time
	_, err := jwt.ParseWithClaims(tokenString, clm, s.keyFunc)
	if err != nil {
		klog.V(4).Info(err)
		return nil, "", err
	}
	return &user.DefaultInfo{Name: clm.Username, Groups: clm.Groups, Extra: clm.Extra}, clm.TokenType, nil
}

func (s *jwtTokenIssuer) IssueTo(user user.Info, tokenType TokenType, expiresIn time.Duration) (string, error) {
	issueAt := time.Now().Unix() - int64(s.maximumClockSkew.Seconds())
	notBefore := issueAt
	clm := &Claims{
		Username:  user.GetName(),
		Groups:    user.GetGroups(),
		Extra:     user.GetExtra(),
		TokenType: tokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issueAt,
			Issuer:    s.name,
			NotBefore: notBefore,
		},
	}

	if expiresIn > 0 {
		clm.ExpiresAt = clm.IssuedAt + int64(expiresIn.Seconds())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)

	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		klog.V(4).Info(err)
		return "", err
	}

	return tokenString, nil
}

func (s *jwtTokenIssuer) keyFunc(token *jwt.Token) (i interface{}, err error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		return s.secret, nil
	} else {
		return nil, fmt.Errorf("expect token signed with HMAC but got %v", token.Header["alg"])
	}
}

func NewTokenIssuer(secret string, maximumClockSkew time.Duration) Issuer {
	return &jwtTokenIssuer{
		name:             DefaultIssuerName,
		secret:           []byte(secret),
		maximumClockSkew: maximumClockSkew,
	}
}
