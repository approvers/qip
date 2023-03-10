package activitypub

import (
	"strings"
)

const InstanceFQDN = "np.test.laminne33569.net"

// Acct ActivityPub上のユーザーの識別子
type Acct struct {
	Host     *string // nilを表現するためにポインタにする
	UserName string
}

func AcctParser(acct string) Acct {
	if string(acct[0]) == "@" {
		sp := strings.Split(acct[1:], "@")
		return Acct{Host: &sp[1], UserName: sp[0]}
	}

	sp := strings.Split(acct, "@")

	return Acct{Host: &sp[1], UserName: sp[0]}
}
