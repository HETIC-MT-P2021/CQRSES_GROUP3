package services

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/config"
	"strconv"
	"errors"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// Claims is the struct for the token claims
type Claims struct {
	Email string
	jwt.StandardClaims
}

// Rsa is the struct to get the rsa keys used to generate and verify tokens from the environment variables
type Rsa struct {
	PublicKeyPath  string
	PrivateKeyPath string
	PublicKey      interface{}
	PrivateKey     interface{}
}

var rsa Rsa

func initRsaKeys() error {
	if rsa.PrivateKey != nil && rsa.PublicKey != nil {
		return nil
	}

	rsa.PublicKeyPath = config.GoDotEnvVariable("RSA_PUBLIC_PATH")
	rsa.PrivateKeyPath = config.GoDotEnvVariable("RSA_PRIVATE_PATH")

	// Get the public key
	publicKeyData, err := ioutil.ReadFile(rsa.PublicKeyPath)
	if err != nil {
		log.Error(err)
		return err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		log.Error("error public key: ", err)
		return err
	}
	rsa.PublicKey = publicKey

	// Get the private key
	privateKeyData, err := ioutil.ReadFile(rsa.PrivateKeyPath)
	if err != nil {
		log.Error(err)
		return err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(privateKeyData, config.GoDotEnvVariable("RSA_PASSWORD"))
	if err != nil {
		log.Error("error public key: ", err)
		return err
	}
	rsa.PrivateKey = privateKey

	return nil
}

// GenerateToken creates a JWT with email and expiration time in the payload
func GenerateToken(email string) (string, error) {
	err := initRsaKeys()
	if err != nil {
		return "", errors.New("Couldn't init rsa keys")
	}

	validTime,_ := strconv.ParseInt(config.GoDotEnvVariable("TOKEN_VALID_DURATION"), 10, 64)
	// Generate Expiration date
	expirationTime := time.Now().Add(time.Duration(validTime) * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// JWT takes unix timestamps
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(rsa.PrivateKey)

	if err != nil {
		log.Error("error while generating token: ", err)
		return "", err
	}

	return tokenString, nil
}

// DecodeToken decode and validates a token
func DecodeToken(tokenString string) (*jwt.Token, *Claims, error) {
	err := initRsaKeys()
	if err != nil {
		return nil, &Claims{}, errors.New("Couldn't init rsa keys")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return rsa.PublicKey, nil
	})
	if err != nil {
		log.Error("Couldn't parse the token : ", err)
		return nil, &Claims{}, err
	}

	return token, claims, nil
}