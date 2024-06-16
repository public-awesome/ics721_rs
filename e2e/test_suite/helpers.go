package test_suite

import (
	"encoding/hex"
	"fmt"

	b64 "encoding/base64"

	wasmibctesting "github.com/CosmWasm/wasmd/x/wasm/ibctesting"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ParseOptional(memo string) string {
	r := ""
	if memo != "" {
		r = fmt.Sprintf("\"%s\"", memo)
	} else {
		r = "null"
	}
	return r
}

func GetCw721SendIbcAwayMessage(path *wasmibctesting.Path, coordinator *wasmibctesting.Coordinator, tokenId string, bridge, receiver sdk.AccAddress, timeout int64, memo string) string {
	memo = ParseOptional(memo)
	ibcAway := fmt.Sprintf(`{ "receiver": "%s", "channel_id": "%s", "timeout": { "timestamp": "%d" }, "memo": %s }`, receiver.String(), path.EndpointA.ChannelID, timeout, memo)
	ibcAwayEncoded := b64.StdEncoding.EncodeToString([]byte(ibcAway))
	return fmt.Sprintf(`{ "send_nft": { "contract": "%s", "token_id": "%s", "msg": "%s" } }`, bridge, tokenId, ibcAwayEncoded)
}

func AccAddressFromHex(address string) (addr sdk.AccAddress, err error) {
	bz, err := addressBytesFromHexString(address)
	return sdk.AccAddress(bz), err
}

func addressBytesFromHexString(address string) ([]byte, error) {
	return hex.DecodeString(address)
}