package strategies

import (
	"context"
	"crypto/hmac"
	rand2 "crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	sdk "github.com/identityOrg/oidcsdk"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"hash"
	"time"
)

type DefaultStrategy struct {
	SecretStore sdk.ISecretStore
	Config      *sdk.Config
	HmacKey     string
}

func NewDefaultStrategy(secretStore sdk.ISecretStore, config *sdk.Config) *DefaultStrategy {
	return &DefaultStrategy{
		SecretStore: secretStore,
		Config:      config,
		HmacKey:     uuid.New().String(),
	}
}

var b64 = base64.URLEncoding.WithPadding(base64.NoPadding)

func (ds *DefaultStrategy) GenerateIDToken(ctx context.Context, profile sdk.RequestProfile, client sdk.IClient, expiry time.Time,
	transactionClaims map[string]interface{}, tokens sdk.Tokens) (string, error) {
	signingKey := jose.SigningKey{
		Algorithm: client.GetIDTokenSigningAlg(),
	}
	keySet, err := ds.SecretStore.GetAllSecrets(ctx)
	if err != nil {
		return "", err
	}
	for _, key := range keySet.Keys {
		if key.Use == "sig" && key.Algorithm == string(client.GetIDTokenSigningAlg()) {
			signingKey.Key = key
		}
	}
	if signingKey.Key == nil {
		err = fmt.Errorf("no key available for algorithm %s", client.GetIDTokenSigningAlg())
		return "", err
	}

	signer, err := jose.NewSigner(signingKey, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	currentTime := time.Now()

	standardClaims := jwt.Claims{
		Issuer:    ds.Config.Issuer,
		Subject:   profile.GetUsername(),
		NotBefore: jwt.NewNumericDate(currentTime),
		IssuedAt:  jwt.NewNumericDate(currentTime),
		Expiry:    jwt.NewNumericDate(expiry),
		ID:        uuid.New().String(),
	}
	if len(profile.GetAudience()) > 0 {
		standardClaims.Audience = []string(profile.GetAudience())
	}
	if transactionClaims == nil {
		transactionClaims = make(map[string]interface{}, 0)
	}
	if profile.GetNonce() != "" {
		transactionClaims["nonce"] = profile.GetNonce()
	}

	transactionClaims["azp"] = client.GetID()
	transactionClaims["typ"] = "ID"
	if tokens.AccessToken != "" {
		rtHash := createHashOctet(signingKey.Algorithm, tokens.AccessToken)
		transactionClaims["at_hash"] = rtHash
	}

	return jwt.Signed(signer).Claims(standardClaims).Claims(transactionClaims).CompactSerialize()
}

func createHashOctet(algorithm jose.SignatureAlgorithm, token string) string {
	var h hash.Hash
	switch algorithm {
	case jose.RS256:
		fallthrough
	case jose.ES256:
		fallthrough
	case jose.PS256:
		h = sha256.New()
	case jose.RS512:
		fallthrough
	case jose.ES512:
		fallthrough
	case jose.PS512:
		h = sha512.New()
	case jose.PS384:
		fallthrough
	case jose.ES384:
		fallthrough
	case jose.RS384:
		h = sha512.New384()
	default:
		return ""
	}
	_, _ = h.Write([]byte(token))
	sum := h.Sum(nil)
	half := sum[:len(sum)/2]
	return b64.EncodeToString(half)
}

func (ds *DefaultStrategy) GenerateRefreshToken() (token string, signature string) {
	return ds.generateAndSign(ds.Config.RefreshTokenEntropy)
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
	return ds.generateAndSign(ds.Config.AccessTokenEntropy)
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
	return ds.generateAndSign(ds.Config.AuthorizationCodeEntropy)
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
