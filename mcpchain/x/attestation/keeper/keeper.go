package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"mcpchain/x/attestation/types"
)

const (
	AttestationPrefix   = "AttestationRecord-value-"
	AttestationCountKey = "AttestationRecord-count"
)

func GetAttestationKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetAttestation(ctx sdk.Context, record types.AttestationRecord) error {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(&record)
	if err != nil {
		return err
	}
	return store.Set(GetAttestationKey(k.GetAttestationCount(ctx)), bz)
}

func (k Keeper) GetAttestationCount(ctx sdk.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get([]byte(AttestationCountKey))
	if err != nil || bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAttestationCount(ctx sdk.Context, count uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	return store.Set([]byte(AttestationCountKey), bz)
}

func (k Keeper) AppendAttestation(ctx sdk.Context, record types.AttestationRecord) (uint64, error) {
	id := k.GetAttestationCount(ctx)

	if err := k.SetAttestation(ctx, record); err != nil {
		return 0, err
	}
	if err := k.SetAttestationCount(ctx, id+1); err != nil {
		return 0, err
	}

	return id, nil
}

func (k Keeper) GetAllAttestations(ctx sdk.Context) []types.AttestationRecord {
	store := k.storeService.OpenKVStore(ctx)
	iterator, err := store.Iterator(nil, nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()

	var list []types.AttestationRecord
	for ; iterator.Valid(); iterator.Next() {
		var v types.AttestationRecord
		if err := k.cdc.Unmarshal(iterator.Value(), &v); err != nil {
			continue
		}
		list = append(list, v)
	}
	return list
}

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
