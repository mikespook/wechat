package wechat

import (
	"strings"
	"testing"
)

const encryptXML = "<xml><ToUserName><![CDATA[gh_10f6c3c3ac5a]]></ToUserName><Encrypt><![CDATA[hQM/NS0ujPGbF+/8yVe61E3mUVWVO1izRlZdyv26zrVUSE3zUEBdcXITxjbjiHH38kexVdpQLCnRfbrqny1yGvgqqKTGKxJWWQ9D5WiiUKxavHRNzYVzAjYkp7esNGy7HJcl/P3BGarQF3+AWyNQ5w7xax5GbOwiXD54yri7xmNMHBOHapDzBslbnTFiEy+8sjSl4asNbn2+ZVBpqGsyKDv0ZG+DlSlXlW+gNPVLP+YxeUhJcyfp91qoa0FJagRNlkNul4mGz+sZXJs0WF7lPx6lslDGW3J66crvIIx/klpl0oa/tC6n/9c8OFQ9pp8hrLq7B9EaAGFlIyz5UhVLiWPN97JkL6JCfxVooVMEKcKRrrlRDGe8RWVM3EW/nxk9Ic37lYY5j97YZfq375AoTBdGDtoPFZsvv3Upyut1i6G0JRogUsMPlyZl9B8Pl/wcA7k7i4LYMr2yK4SxNFrBUw==]]></Encrypt></xml>"

func TestNewEncryptXML(t *testing.T) {
	r := strings.NewReader(encryptXML)
	data, err := NewEncryptXML(r)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s:%s", data.ToUserName, data.Encrypt)
}
