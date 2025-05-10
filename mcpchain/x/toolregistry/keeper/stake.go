package keeper

import (
	"fmt"
	"mcpchain/x/toolregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetStakeRecord(ctx sdk.Context, record types.StakeRecord) error {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(&record)
	if err != nil {
		return err
	}
	return store.Set([]byte(getStakeKey(record.ServerId)), bz)
}

func getStakeKey(serverId uint64) string {
	return fmt.Sprintf("stake-record-%d", serverId)
}
