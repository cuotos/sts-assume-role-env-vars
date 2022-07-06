package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseAssumeRoleOutput(t *testing.T) {
	input := `{
    "Credentials": {
        "AccessKeyId": "ASIA...UUP2",
        "SecretAccessKey": "fYS...in",
        "SessionToken": "IQoJ...P46W4F/IEX",
        "Expiration": "2022-07-05T14:28:30+00:00"
    },
    "AssumedRoleUser": {
        "AssumedRoleId": "AROA...5YCM:danp",
        "Arn": "arn:aws:sts::111111111111:assumed-role/runner-role/cuotos"
    }
}`

	expected := ParsedAssumeRoleCreds{
		AccessKeyID:  "ASIA...UUP2",
		SecretKey:    "fYS...in",
		SessionToken: "IQoJ...P46W4F/IEX",
	}

	actual, err := ParseAssumeRoleOutput(strings.NewReader(input))

	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestOutputExportEnvVarsStatements(t *testing.T) {
	input := ParsedAssumeRoleCreds{
		AccessKeyID:  "accessKey",
		SecretKey:    "secretKey",
		SessionToken: "sessionToken",
	}

	expected := `export AWS_ACCESS_KEY_ID=accessKey
export AWS_SECRET_ACCESS_KEY=secretKey
export AWS_SESSION_TOKEN=sessionToken`

	actual := GenerateSetEnvVarStatement(input)

	assert.Equal(t, expected, actual)
}
