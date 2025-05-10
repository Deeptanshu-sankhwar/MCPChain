package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"mcpchain/x/toolregistry/types"
)

const (
	ServerPrefix   = "RegisteredServer/value/"
	ServerCountKey = "RegisteredServer-count"
)

func GetServerKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetServer(ctx sdk.Context, server types.RegisteredServer) error {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(&server)
	if err != nil {
		return err
	}
	return store.Set(GetServerKey(server.Id), bz)
}

func (k Keeper) GetServer(ctx sdk.Context, id uint64) (types.RegisteredServer, bool) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(GetServerKey(id))
	if err != nil || bz == nil {
		return types.RegisteredServer{}, false
	}
	var out types.RegisteredServer
	if err := k.cdc.Unmarshal(bz, &out); err != nil {
		return types.RegisteredServer{}, false
	}
	return out, true
}

func (k Keeper) DeleteServer(ctx sdk.Context, id uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(GetServerKey(id))
}

func (k Keeper) GetServerCount(ctx sdk.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get([]byte(ServerCountKey))
	if err != nil || bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetServerCount(ctx sdk.Context, count uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	return store.Set([]byte(ServerCountKey), bz)
}

func (k Keeper) AppendRegisteredServer(ctx sdk.Context, server types.RegisteredServer) (uint64, error) {
	id := k.GetServerCount(ctx)
	server.Id = id

	if err := k.SetServer(ctx, server); err != nil {
		return 0, err
	}
	if err := k.SetServerCount(ctx, id+1); err != nil {
		return 0, err
	}

	return id, nil
}

func (k Keeper) GetAllRegisteredServers(ctx sdk.Context) []types.RegisteredServer {
	store := k.storeService.OpenKVStore(ctx)
	iterator, err := store.Iterator(nil, nil)
	if err != nil {
		return nil
	}
	defer iterator.Close()

	var servers []types.RegisteredServer
	for ; iterator.Valid(); iterator.Next() {
		if string(iterator.Key()) == ServerCountKey {
			continue
		}
		var server types.RegisteredServer
		if err := k.cdc.Unmarshal(iterator.Value(), &server); err != nil {
			continue
		}
		servers = append(servers, server)
	}
	return servers
}

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority  string
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		bankKeeper:   bankKeeper,
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
