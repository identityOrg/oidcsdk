package strategies

import (
	"crypto/hmac"
	rand2 "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	sdk "github.com/identityOrg/oidcsdk"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

type DefaultStrategy struct {
	SecretStore              sdk.ISecretStore
	HmacKey                  string
	AccessTokenEntropy       int
	AuthorizationCodeEntropy int
	RefreshTokenEntropy      int
	Issuer                   string
}

func NewDefaultStrategy() *DefaultStrategy {
	return &DefaultStrategy{}
}

var b64 = base64.URLEncoding.WithPadding(base64.NoPadding)

func (ds *DefaultStrategy) Configure(config *sdk.Config, args ...interface{}) {
	ds.AccessTokenEntropy = config.AccessTokenEntropy
	ds.AuthorizationCodeEntropy = config.AuthorizationCodeEntropy
	ds.RefreshTokenEntropy = config.RefreshTokenEntropy
	if ds.HmacKey == "" {
		ds.HmacKey = uuid.New().String()
	}
	for _, arg := range args {
		if ss, ok := arg.(sdk.ISecretStore); ok {
			ds.SecretStore = ss
		}
	}
}

func (ds *DefaultStrategy) GenerateIDToken(profile sdk.RequestProfile, client sdk.IClient, expiry time.Time,
	transactionClaims map[string]interface{}) (idToken string, err error) {
	signingKey := jose.SigningKey{
		Algorithm: client.GetIDTokenSigningAlg(),
	}
	keySet := ds.SecretStore.GetAllSecrets()
	for _, key := range keySet.Keys {
		if key.Use == "sign" && key.Algorithm == string(client.GetIDTokenSigningAlg()) {
			signingKey.Key = key
		}
	}
	if signingKey.Key == nil {
		err = fmt.Errorf("no key available for algorithm %s", client.GetIDTokenSigningAlg())
		return
	}

	signer, err := jose.NewSigner(signingKey, (&jose.SignerOptions{}).WithType("JWT"))
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
	decodeBytes, err := b64.DecodeString(token)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = b64.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) GenerateAccessToken() (token string, signature string) {
	return ds.generateAndSign(ds.AccessTokenEntropy)
}

func (ds *DefaultStrategy) SignAccessToken(token string) (signature string, err error) {
	decodeBytes, err := b64.DecodeString(token)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = b64.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) GenerateAuthCode() (code string, signature string) {
	return ds.generateAndSign(ds.AuthorizationCodeEntropy)
}

func (ds *DefaultStrategy) generateAndSign(length int) (code string, signature string) {
	codeBytes := generate(length)
	signedBytes := ds.sigh(codeBytes)
	code = b64.EncodeToString(codeBytes)
	signature = b64.EncodeToString(signedBytes)
	return
}

func (ds *DefaultStrategy) SignAuthCode(code string) (signature string, err error) {
	decodeBytes, err := b64.DecodeString(code)
	if err != nil {
		return
	}
	signedBytes := ds.sigh(decodeBytes)
	signature = b64.EncodeToString(signedBytes)
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
