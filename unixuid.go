package unixuid

import (
	"os/user"
	"strconv"
)

type UnixUID int

func (uid UnixUID) MarshalText() (text []byte, err error) {
	u, err := user.LookupId(strconv.FormatUint(uint64(uid), 10))
	if nil != err {
		return
	}
	return []byte(u.Username), nil
}

func (uidRef *UnixUID) UnmarshalText(text []byte) (err error) {
	u, err := user.Lookup(string(text))
	if nil != err {
		return
	}
	uid, err := strconv.ParseUint(u.Uid, 10, 64)
	if nil != err {
		return
	}
	*uidRef = UnixUID(uid)
	return
}
