package strategies

import (
	"crypto"
	rand2 "crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	sdk "oauth2-oidc-sdk"
)

type DefaultStrategy struct {
	PrivateKey          *rsa.PrivateKey
	PublicKey           *rsa.PublicKey
	AccessTokenEntropy  uint8
	AccessCodeEntropy   uint8
	RefreshTokenEntropy uint8
}

func NewDefaultStrategy(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *DefaultStrategy {
	return &DefaultStrategy{PrivateKey: privateKey, PublicKey: publicKey}
}

func (ds *DefaultStrategy) GenerateIDToken(profile sdk.IProfile, client sdk.IClient,
	transactionClaims map[string]interface{}) (idToken string, err error) {
	client.GetIDTokenSigningAlg()
	key := jose.SigningKey{
		Algorithm: client.GetIDTokenSigningAlg(),
		Key:       ds.PrivateKey,
	}
	signer, err := jose.NewSigner(key, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return
	}
	idToken, err = jwt.Signed(signer).Claims(transactionClaims).Claims(profile.GetClaims()).FullSerialize()
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
	return ds.generateAndSign(ds.AccessCodeEntropy)
}

func (ds *DefaultStrategy) ValidateAuthCode(code string, signature string) error {
	return ds.validate(code, signature)
}

func (ds *DefaultStrategy) generateAndSign(length uint8) (code string, signature string) {
	codeBytes := generate(length)
	signed, err := rsa.SignPKCS1v15(rand2.Reader, ds.PrivateKey, crypto.SHA256, codeBytes)
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

func generate(length uint8) (codeByte []byte) {
	codeByte = make([]byte, length)
	_, _ = rand2.Read(codeByte)
	return
}

func (ds *DefaultStrategy) sigh(code string) (signature string) {
	signed, err := rsa.SignPKCS1v15(rand2.Reader, ds.PrivateKey, crypto.SHA256, []byte(code))
	if err != nil {
		panic(err)
	}
	signature = base64.URLEncoding.EncodeToString(signed)
	return
}
