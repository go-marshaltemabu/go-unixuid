package unixuid

import (
	"os/user"
	"strconv"
)

type UnixUID uint32

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

// MakeUIDIntMap returns a map of int from a slice of UnixUID.
func MakeUIDIntMap(uidSlice []UnixUID) (uidMap map[int]struct{}) {
	uidMap = make(map[int]struct{})
	for _, uid := range uidSlice {
		uidMap[int(uid)] = struct{}{}
	}
	return
}

// MakeUIDIntMap returns a map of int from a slice of UnixUID.
func MakeUIDUint32Map(uidSlice []UnixUID) (uidMap map[uint32]struct{}) {
	uidMap = make(map[uint32]struct{})
	for _, uid := range uidSlice {
		uidMap[uint32(uid)] = struct{}{}
	}
	return
}

// CloneAsIntSlice returns a slice of int from a slice of UnixUID.
func CloneAsIntSlice(uidSlice []UnixUID) (result []int) {
	result = make([]int, len(uidSlice))
	for idx, uid := range uidSlice {
		result[idx] = int(uid)
	}
	return
}

func CloneAsUint32Slice(uidSlice []UnixUID) (result []uint32) {
	result = make([]uint32, len(uidSlice))
	for idx, uid := range uidSlice {
		result[idx] = uint32(uid)
	}
	return
}
