package strategies

import (
	"crypto"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	sdk "oauth2-oidc-sdk"
	"time"
)

type DefaultStrategy struct {
	PrivateKey               *rsa.PrivateKey
	PublicKey                *rsa.PublicKey
	AccessTokenEntropy       int
	AuthorizationCodeEntropy int
	RefreshTokenEntropy      int
	Issuer                   string
}

func (ds *DefaultStrategy) Configure(_ interface{}, config *sdk.Config, _ ...interface{}) {
	ds.AccessTokenEntropy = config.AccessTokenEntropy
	ds.AuthorizationCodeEntropy = config.AuthorizationCodeEntropy
	ds.RefreshTokenEntropy = config.RefreshTokenEntropy
}

func NewDefaultStrategy(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *DefaultStrategy {
	return &DefaultStrategy{PrivateKey: privateKey, PublicKey: publicKey}
}

func (ds *DefaultStrategy) GenerateIDToken(profile sdk.IProfile, client sdk.IClient, expiry time.Time,
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
	idToken, err = jwt.Signed(signer).Claims(standardClaims).Claims(transactionClaims).FullSerialize()
	return
}

func (ds *DefaultStrategy) GenerateRefreshToken() (token string, signature string) {
	return ds.generateAndSign(ds.RefreshTokenEntropy)
}

func (ds *DefaultStrategy) ValidateRefreshToken(token string, signature string) error {
	return ds.validate(token, signature)
}

func (ds *DefaultStrategy) SignRefreshToken(token string) string {
	return ds.sigh(token)
}

func (ds *DefaultStrategy) GenerateAccessToken() (token string, signature string) {
	return ds.generateAndSign(ds.AccessTokenEntropy)
}

func (ds *DefaultStrategy) ValidateAccessToken(token string, signature string) error {
	return ds.validate(token, signature)
}

func (ds *DefaultStrategy) GenerateAuthCode() (code string, signature string) {
	return ds.generateAndSign(ds.AuthorizationCodeEntropy)
}

func (ds *DefaultStrategy) ValidateAuthCode(code string, signature string) error {
	return ds.validate(code, signature)
}

func (ds *DefaultStrategy) generateAndSign(length int) (code string, signature string) {
	codeBytes := generate(length)
	hash := sha256.Sum256(codeBytes)
	signed, err := rsa.SignPKCS1v15(rand2.Reader, ds.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}
	code = base64.URLEncoding.EncodeToString(codeBytes)
	signature = base64.URLEncoding.EncodeToString(signed)
	return
}

func (ds *DefaultStrategy) SignAuthCode(code string) string {
	return ds.sigh(code)
}

func (ds *DefaultStrategy) validate(code string, signature string) (err error) {
	codeBytes, err := base64.URLEncoding.DecodeString(code)
	if err != nil {
		return
	}
	signatureBytes, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return
	}
	err = rsa.VerifyPKCS1v15(ds.PublicKey, crypto.SHA256, codeBytes, signatureBytes)
	return
}

func generate(length int) (codeByte []byte) {
	codeByte = make([]byte, length)
	_, _ = rand2.Read(codeByte)
	return
}

func (ds *DefaultStrategy) sigh(code string) (signature string) {
	hash := sha256.Sum256([]byte(code))
	signed, err := rsa.SignPKCS1v15(rand2.Reader, ds.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}
	signature = base64.URLEncoding.EncodeToString(signed)
	return
}
