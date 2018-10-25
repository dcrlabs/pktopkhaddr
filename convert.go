// Copyright (c) 2018, Jonathan Chappelow

package pktopkhaddr

import (
	"github.com/decred/dcrd/dcrec"
	"github.com/decred/dcrd/dcrutil"
)

// PKAddr2PKHAddr converts a P2PK (pay-to-pubkey) address to a P2PKH
// (pay-to-pubkey-hash) address. The generated address specifies an ECDSA
// signature over the secp256k1 elliptic curve.
func PKAddr2PKHAddr(p2pk string) (p2pkh string, err error) {
	// Attempt to decode the pay-to-pubkey address.
	var addr dcrutil.Address
	addr, err = dcrutil.DecodeAddress(p2pk)
	if err != nil {
		return
	}

	// Extract the pubkey hash.
	addrHash := addr.Hash160()

	// Create a new pay-to-pubkey-hash address.
	addrPKH, err := dcrutil.NewAddressPubKeyHash(addrHash[:], addr.Net(), dcrec.STEcdsaSecp256k1)
	if err != nil {
		return
	}
	return addrPKH.EncodeAddress(), err
}
