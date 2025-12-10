package jsonutil

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	tokenValues := make(map[string]json.Token)

	err := Scan(exampleJson, func(tokenPath *TokenPath, t json.Token) bool {
		tokenValues[tokenPath.String()] = t
		return true
	})
	require.NoError(t, err)

	assert.Equal(t, "0", tokenValues["$.version"])
	assert.Equal(t, "ddddd4-aaaa-7777-4444-345dd43cc333", tokenValues["$.id"])
	assert.Equal(t, "EC2 Instance State-change Notification", tokenValues["$.detail-type"])
	assert.Equal(t, "aws.ec2", tokenValues["$.source"])
	assert.Equal(t, "012345679012", tokenValues["$.account"])
	assert.Equal(t, "2017-10-02T16:24:49Z", tokenValues["$.time"])
	assert.Equal(t, "us-east-1", tokenValues["$.region"])
	assert.Equal(t, "arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00000", tokenValues["$.resources[0]"])
	assert.Equal(t, "arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00001", tokenValues["$.resources[1]"])
	assert.Equal(t, json.Number("5"), tokenValues["$.detail.c-count"])
	assert.Equal(t, json.Number("3"), tokenValues["$.detail.d-count"])
	assert.Equal(t, json.Number("301.8"), tokenValues["$.detail.x-limit"])
	assert.Equal(t, "10.0.0.33", tokenValues["$.detail.source-ip"])
	assert.Equal(t, "i-000000aaaaaa00000", tokenValues["$.detail.instance-id"])
	assert.Equal(t, "running", tokenValues["$.detail.state"])
}
