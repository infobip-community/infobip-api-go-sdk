package models

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidEmailMessage(t *testing.T) {
	tests := []struct {
		name     string
		instance EmailMsg
	}{
		{
			name: "minimum input",
			instance: EmailMsg{
				From:    "someone@infobip.com",
				To:      "someone@outside.com",
				Subject: "Some subject",
			},
		},
		{
			name:     "full input",
			instance: GenerateEmailMsg(),
		},
	}

	content := []byte("temporary file's content")
	attachment, err := ioutil.TempFile("", "example")
	require.NoError(t, err)
	_, err = attachment.Write(content)
	require.NoError(t, err)
	_, err = attachment.Seek(0, 0)
	require.NoError(t, err)

	image, err := os.Open("testdata/image.png")
	require.NoError(t, err)

	tests[1].instance.Attachment = attachment
	tests[1].instance.InlineImage = image

	defer image.Close()
	defer attachment.Close()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err = tc.instance.Validate()
			require.NoError(t, err)

			_, err = tc.instance.Marshal()
			require.NoError(t, err)

			require.NotEmpty(t, tc.instance.GetMultipartBoundary())
		})
	}
}

func TestEmailMsgConstraints(t *testing.T) {
	tests := []struct {
		name     string
		instance EmailMsg
	}{
		{
			name:     "empty",
			instance: EmailMsg{},
		},
		{
			name: "empty from",
			instance: EmailMsg{
				To:      "someone@outside.com",
				Subject: "Some subject",
			},
		},
		{
			name: "empty to",
			instance: EmailMsg{
				From:    "someone@infobip.com",
				Subject: "Some subject",
			},
		},
		{
			name: "empty subject",
			instance: EmailMsg{
				From: "someone@infobip.com",
				To:   "someone@outside.com",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.instance.Validate()
			require.Error(t, err)
		})
	}
}
