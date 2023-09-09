package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"ibc_sequencer/x/sequencer/types"
)

func TestBlockMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateBlock(ctx, &types.MsgCreateBlock{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestBlockMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBlock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBlock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBlock{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBlock(ctx, &types.MsgCreateBlock{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateBlock(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBlockMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBlock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBlock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBlock{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateBlock(ctx, &types.MsgCreateBlock{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBlock(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
