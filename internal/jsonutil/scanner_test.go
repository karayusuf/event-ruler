package jsonutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFoo(t *testing.T) {
	exampleJson := strings.NewReader(`
		{
			"version": "0",
			"id": "ddddd4-aaaa-7777-4444-345dd43cc333",
			"detail-type": "EC2 Instance State-change Notification",
			"source": "aws.ec2",
			"account": "012345679012",
			"time": "2017-10-02T16:24:49Z",
			"region": "us-east-1",
			"resources": [
				"arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00000",
				"arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00001"
			],
			"detail": {
				"c-count": 5,
				"d-count": 3,
				"x-limit": 301.8,
				"source-ip": "10.0.0.33",
				"instance-id": "i-000000aaaaaa00000",
				"state": "running"
			}
		}
	`)

	s := &Scanner{}

	err := s.Scan(exampleJson, func(tokenPath *tokenPath, t json.Token) bool {
		fmt.Printf("path:'%s' value: %v\n", tokenPath.String(), t)

		return true
	})
	require.Error(t, err)
}
