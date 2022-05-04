package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidSendSMSRequest(t *testing.T) {
	req := GenerateSendSMSRequest()

	err := req.Validate()
	assert.NoError(t, err)
}

func TestInvalidSendSMSRequest(t *testing.T) {

}

func TestSMSMsgMarshalJSON(t *testing.T) {

}
