package token

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTMaker struct {
	secretKey string
}

type Payload struct {
	Username string `json: "username"`
}

// Valid implements jwt.Claims.
func (*Payload) Valid() error {
	panic("unimplemented")
}

func main() {
	r := gin.Default()

}

func (maker *JWTMaker) VerifyToken(tokenstring string) (_, error) {

	// claims := jwt.MapClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	/*
	   Function jwt.ParseWithClaims accept an interface of jwt.Claims
	   as the second argument. Besides struct-based custom claims,
	   the package also provides map-based claims, i.e. jwt.MapClaims.
	   So, you can simply decode the token into a MapClaims
	*/

	_, err := jwt.ParseWithClaims(tokenstring, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	for key, val := range claims {
		fmt.Println("key: %v, value: %v\n", key, val)
	}
	// payload, ok := jwtToken.Claims.(*Payload)
	// if !ok {
	// 	return nil, ErrInvalidToken
	// }
	// return payload, nil
}

// _, err := jwt.ParseWithClaims(tokenstring, claims, keyFunc)
// if err != nil {
// 	verr, ok := err.(*jwt.ValidationError)
// 	if ok && errors.Is(verr.Inner, ErrExpiredToken) {
// 		return nil, ErrExpiredToken
// 	}
// 	return nil, ErrInvalidToken
// }

// for key, val := range claims {
// 	fmt.Println("key: %v, value: %v\n", key, val)
// }
// // payload, ok := jwtToken.Claims.(*Payload)
// // if !ok {
// // 	return nil, ErrInvalidToken
// // }
// // return payload, nil

// if request.Header["Token"] != nil {
// 	token, err := jwt.Parse(request.Header\["Token"\][0],func(token *jwt.Token)(interface{}, error){
// 		_,ok := token.Method.(*jwt.SigningMethodHS256)
// 		if !ok{
// 			writer.WriteHeader(http.StatusUnauthorized)
// 			_, err := writer.Write([]byte("Unauthorized"))
// 			if err!=nil{
// 				return nil, err
// 			}
// 		}
// 		return "",nil
// 	})
// }

// token, err := jwt.Parse(request.Header\["Token"\][0], func(token *jwt.Token) (interface{}, error) {
// 	_, ok := token.Method.(*jwt.SigningMethodECDSA)
// 	if !ok {
// 	   writer.WriteHeader(http.StatusUnauthorized)
// 	   _, err := writer.Write([]byte("You're Unauthorized!"))
// 	   if err != nil {
// 		  return nil, err

// 	   }
// 	}
// 	return "", nil

//  })
