package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	appVersion string = "unset"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println(appVersion)
		os.Exit(0)
	}
	run(os.Stdin)
}

func run(input io.Reader) (string, error) {
	stsOutput, err := ParseAssumeRoleOutput(input)
	if err != nil {
		log.Fatal(err)
	}
	return GenerateSetEnvVarStatement(stsOutput), nil
}

func ParseAssumeRoleOutput(input io.Reader) (ParsedAssumeRoleCreds, error) {
	output := ParsedAssumeRoleCreds{}

	decoder := json.NewDecoder(input)

	assumeRoleOutput := AssumeRoleOutput{}

	err := decoder.Decode(&assumeRoleOutput)
	if err != nil {
		return output, nil
	}

	output.AccessKeyID = assumeRoleOutput.Credentials.AccessKeyID
	output.SecretKey = assumeRoleOutput.Credentials.SecretAccessKey
	output.SessionToken = assumeRoleOutput.Credentials.SessionToken

	return output, nil
}

func GenerateSetEnvVarStatement(creds ParsedAssumeRoleCreds) string {
	format := `export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s
export AWS_SESSION_TOKEN=%s`

	return fmt.Sprintf(format, creds.AccessKeyID, creds.SecretKey, creds.SessionToken)
}

type ParsedAssumeRoleCreds struct {
	AccessKeyID  string
	SecretKey    string
	SessionToken string
}

type AssumeRoleOutput struct {
	Credentials struct {
		AccessKeyID     string    `json:"AccessKeyId"`
		SecretAccessKey string    `json:"SecretAccessKey"`
		SessionToken    string    `json:"SessionToken"`
		Expiration      time.Time `json:"Expiration"`
	} `json:"Credentials"`
	AssumedRoleUser struct {
		AssumedRoleID string `json:"AssumedRoleId"`
		Arn           string `json:"Arn"`
	} `json:"AssumedRoleUser,omitempty"`
}
