package digitalocean

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {
	account := &Account{RegionId: 1122, SizeId: 10}
	assert.NotNil(t, account)
	droplet := account.DefaultDroplet()
	assert.NotNil(t, droplet)
	assert.Equal(t, account.RegionId, 1122)
	assert.Equal(t, account.SizeId, 10)
}
