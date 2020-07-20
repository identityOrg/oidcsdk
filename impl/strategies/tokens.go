package strategies

import (
	"crypto/hmac"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	sdk "oidcsdk"
	"time"
)

type DefaultStrategy struct {
	PrivateKey               *rsa.PrivateKey
	PublicKey                *rsa.PublicKey
	HmacKey                  string
	AccessTokenEntropy       int
	AuthorizationCodeEntropy int
	RefreshTokenEntropy      int
	Issuer                   string
}

func NewDefaultStrategy(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *DefaultStrategy {
	return &DefaultStrategy{PrivateKey: privateKey, PublicKey: publicKey}
}

func (ds *DefaultStrategy) Configure(_ interface{}, config *sdk.Config, _ ...interface{}) {
	ds.AccessTokenEntropy = config.AccessTokenEntropy
	ds.AuthorizationCodeEntropy = config.AuthorizationCodeEntropy
	ds.RefreshTokenEntropy = config.RefreshTokenEntropy
	if ds.HmacKey == "" {
		ds.HmacKey = uuid.New().String()
	}
}

func (ds *DefaultStrategy) GenerateIDToken(profile sdk.RequestProfile, client sdk.IClient, expiry time.Time,
	transactionClaims map[string]interface{}) (idToken string, err error) {
	key := jose.SigningKey{
		Algorithm: client.GetIDTokenSigningAlg(),
		Key:       ds.PrivateKey,
	}

	signer, err := jose.NewSigner(key, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return
	}

	currentTime := time.Now()

	standardClaims := jwt.Claims{
		Issuer:    ds.Issuer,
		Subject:   profile.GetUsername(),
		Audience:  []string(profile.GetAudience()),
		NotBefore: jwt.NewNumericDate(currentTime),
		IssuedAt:  jwt.NewNumericDate(currentTime),
		Expiry:    jwt.NewNumericDate(expiry),
		ID:        uuid.New().String(),
	}
	idToken, err = jwt.Signed(signer).Claims(standardClaims).Claims(transactionClaims).CompactSerialize()
	return
}

func (ds *DefaultStrategy) GenerateRefreshToken() (token string, signature string) {
	return ds.generateAndSign(ds.RefreshTokenEntropy)
}

func (ds *DefaultStrategy) SignRefreshToken(token string) (signature string, err error) {
	decodeBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = base64.URLEncoding.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) GenerateAccessToken() (token string, signature string) {
	return ds.generateAndSign(ds.AccessTokenEntropy)
}

func (ds *DefaultStrategy) SignAccessToken(token string) (signature string, err error) {
	decodeBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = base64.URLEncoding.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) GenerateAuthCode() (code string, signature string) {
	return ds.generateAndSign(ds.AuthorizationCodeEntropy)
}

func (ds *DefaultStrategy) generateAndSign(length int) (code string, signature string) {
	codeBytes := generate(length)
	signedBytes := ds.sigh(codeBytes)
	code = base64.URLEncoding.EncodeToString(codeBytes)
	signature = base64.URLEncoding.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) SignAuthCode(code string) (signature string, err error) {
	decodeBytes, err := base64.URLEncoding.DecodeString(code)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = base64.URLEncoding.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) sigh(code []byte) (signature []byte) {
	mac := hmac.New(sha256.New, []byte(ds.HmacKey))
	mac.Write(code)
	signature = mac.Sum(nil)
	return
}

func generate(length int) (codeByte []byte) {
	if length < 1 {
		return []byte(uuid.New().String())
	}
	codeByte = make([]byte, length)
	_, _ = rand2.Read(codeByte)
	return
}
