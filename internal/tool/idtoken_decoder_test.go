package tool_test

import (
	"testing"

	"github.com/orvosi/api/internal/tool"
	"github.com/stretchr/testify/assert"
)

var (
	audience    = "test-audience"
	credentials = []byte(`{
		"type":"service_account",
		"project_id":"project-id",
		"private_key_id":"kmzwa8awaa",
		"private_key":"-----BEGIN PRIVATE KEY-----\nbody\n-----END PRIVATE KEY-----\n",
		"client_email":"projec@project-id.iam.gserviceaccount.com",
		"client_id":"1234567890",
		"auth_uri":"https://accounts.google.com/o/oauth2/auth",
		"token_uri":"https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/project%40project-id.iam.gserviceaccount.com"
	 }`)
)

type payload struct {
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	Claims   map[string]interface{} `json:"-"`
}

// jwt represents the segments of a jwt and exposes convenience methods for
// working with the different segments.
type jwt struct {
	header    string
	payload   string
	signature string
}

// jwtHeader represents a parted jwt's header segment.
type jwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	KeyID     string `json:"kid"`
}

func TestNewIDTokenDecoder(t *testing.T) {
	t.Run("fail to create an instance of IDTokenDecoder due to invalid credentials", func(t *testing.T) {
		dec, err := tool.NewIDTokenDecoder(audience, []byte(`invalid`))

		assert.NotNil(t, err)
		assert.Nil(t, dec)
	})

	t.Run("successfully create an instance of IDTokenDecoder", func(t *testing.T) {
		dec, err := tool.NewIDTokenDecoder(audience, credentials)

		assert.Nil(t, err)
		assert.NotNil(t, dec)
	})
}

func TestIDTokenDecoder_Decode(t *testing.T) {
	t.Run("fail to decode id token due to invalid audience", func(t *testing.T) {

	})
}
