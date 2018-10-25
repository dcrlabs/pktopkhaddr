// Copyright (c) 2018, Jonathan Chappelow

package pktopkhaddr

import (
	"testing"
)

func TestPKAddr2PKHAddr(t *testing.T) {
	pkhAddr, err := PKAddr2PKHAddr("DkM3EyZ546GghVSkvzb6J47PvGDyntqiDtFgipQhNj78Xm2mUYRpf")
	if err != nil {
		t.Error(t)
	}

	expectedPkhAddr := "DsfFjaADsV8c5oHWx85ZqfxCZy74K8RFuhK"
	if pkhAddr != expectedPkhAddr {
		t.Errorf("Incorrect P2PKH address. Got %s, expected %s.",
			pkhAddr, expectedPkhAddr)
	}
}
