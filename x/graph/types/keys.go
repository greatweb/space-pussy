package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName 	  		= "graph"
	RouterKey 	  		= ModuleName
	StoreKey      	    = ModuleName
	QuerierRoute  		= ModuleName
)

var (
	GlobalStoreKeyPrefix  	 = []byte{0x00}
	CidStoreKeyPrefix	  	 = []byte{0x01}
	CidReverseStoreKeyPrefix = []byte{0x02}
	CyberlinkStoreKeyPrefix  = []byte{0x03}

	LastCidNumber    		 = append(GlobalStoreKeyPrefix, []byte("lastCidNumber")...)
	LinksCount 				 = append(GlobalStoreKeyPrefix, []byte("linksCount")...)
)

func CidStoreKey(cid Cid) []byte {
	return append(CidStoreKeyPrefix, []byte(cid)...)
}

func CidReverseStoreKey(num CidNumber) []byte {
	return append(CidReverseStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(num))...)
}

func CyberlinksStoreKey(id uint64) []byte {
	return append(CyberlinkStoreKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}