package hafas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateChecksum(t *testing.T) {
	payload := []byte("abcdefghijklmnopqrstuvwxyzABCDEF")
	expectedChecksum := "b8df5f2cccdef7ae45dd960813caa4fb"

	client := NewHafasClient(&Config{
		Url:  "",
		Salt: "32bebb99c2a8905f2f6890af311a5193",
		Aid:  "",
	})

	checksum := client.createChecksum(payload)

	assert.Equal(t, expectedChecksum, checksum)
}
