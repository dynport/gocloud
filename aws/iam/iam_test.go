package iam

import (
	"encoding/xml"
	"github.com/dynport/gocloud/testhelpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	f := testhelpers.MustReadFixture(t, "get_user.xml")
	rsp := &GetUserResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	user := rsp.User
	assert.Equal(t, user.Path, "/division_abc/subdivision_xyz/")
	assert.Equal(t, user.UserName, "Bob")
	assert.Contains(t, user.Arn, "arn:aws:iam::123456789012:user/division_abc/subdivision_xyz/Bob")
}

func TestAccountSummary(t *testing.T) {
	f := testhelpers.MustReadFixture(t, "get_account_summary.xml")
	rsp := &GetAccountSummaryResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	m := rsp.SummaryMap
	assert.Equal(t, len(m.Entries), 14)

	entry := m.Entries[0]
	assert.Equal(t, entry.Key, "Groups")
	assert.Equal(t, entry.Value, "31")
}

func TestListUsers(t *testing.T) {
	f := testhelpers.MustReadFixture(t, "list_users.xml")
	rsp := &ListUsersResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	assert.Equal(t, len(rsp.Users), 2)
	assert.Equal(t, rsp.Users[0].UserName, "Andrew")
}

func TestAccountAliases(t *testing.T) {
	f := testhelpers.MustReadFixture(t, "list_account_aliases.xml")
	rsp := &ListAccountAliasesResponse{}
	e := xml.Unmarshal(f, rsp)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	assert.Equal(t, len(rsp.AccountAliases), 1)
	assert.Equal(t, rsp.AccountAliases[0], "foocorporation")
}
