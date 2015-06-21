package wechat

import "testing"

func TestAccessToken(t *testing.T) {
	token, exp, err := AccessToken("-- replace your own id --", "-- replace your own secret --")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s (%d)", token, exp)
}
