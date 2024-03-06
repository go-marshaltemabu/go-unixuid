package unixuid_test

import (
	"testing"

	unixuid "github.com/go-marshaltemabu/go-unixuid"
)

func TestMarshalText_0_root(t *testing.T) {
	uidObject := unixuid.UnixUID(0)
	textBuf, err := uidObject.MarshalText()
	if nil != err {
		t.Errorf("MarshalText failed: %v", err)
	}
	if resultValue := string(textBuf); resultValue != "root" {
		t.Errorf("unexpect result: [%s]", resultValue)
	}
}

func TestMarshalText_65534(t *testing.T) {
	uidObject := unixuid.UnixUID(65534)
	textBuf, err := uidObject.MarshalText()
	if nil != err {
		t.Errorf("MarshalText failed: %v", err)
	}
	if resultValue := string(textBuf); (resultValue != "nobody") && (resultValue != "nfsnobody") {
		t.Errorf("unexpect result: [%s]", resultValue)
	}
}

func TestMarshalText_765765123(t *testing.T) {
	uidObject := unixuid.UnixUID(765765123)
	textBuf, err := uidObject.MarshalText()
	if nil == err {
		t.Errorf("unexpect result: [%s] %v", string(textBuf), err)
	}
}

func TestUnmarshalText_0_root(t *testing.T) {
	uidObject := unixuid.UnixUID(32767)
	textBuf := []byte("root")
	err := uidObject.UnmarshalText(textBuf)
	if nil != err {
		t.Errorf("UnmarshalText failed: %v", err)
	}
	if resultValue := int(uidObject); resultValue != 0 {
		t.Errorf("unexpect result: [%d]", resultValue)
	}
}

func TestUnmarshalText_nobody(t *testing.T) {
	var uidObject unixuid.UnixUID
	textBuf := []byte("nobody")
	err := uidObject.UnmarshalText(textBuf)
	if nil != err {
		t.Errorf("UnmarshalText failed: %v", err)
	}
	if resultValue := int(uidObject); (resultValue != 65534) && (resultValue != 99) {
		t.Errorf("unexpect result: [%d]", resultValue)
	}
}

func TestUnmarshalText_shouldnotexist(t *testing.T) {
	var uidObject unixuid.UnixUID
	textBuf := []byte("should-not-exist-user")
	err := uidObject.UnmarshalText(textBuf)
	if nil == err {
		t.Errorf("unexpected result: [%d] %v", int(uidObject), err)
	}
}
