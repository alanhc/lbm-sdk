package collection

import (
	"github.com/line/lbm-sdk/codec/types"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/msgservice"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContract{},
		&MsgIssueFT{},
		&MsgIssueNFT{},
		&MsgMintFT{},
		&MsgMintNFT{},
		&MsgAttach{},
		&MsgDetach{},
		&MsgTransferFT{},
		&MsgTransferFTFrom{},
		&MsgTransferNFT{},
		&MsgTransferNFTFrom{},
		&MsgApprove{},
		&MsgDisapprove{},
		&MsgBurnFT{},
		&MsgBurnFTFrom{},
		&MsgBurnNFT{},
		&MsgBurnNFTFrom{},
		&MsgModify{},
		&MsgGrantPermission{},
		&MsgRevokePermission{},
		&MsgAttachFrom{},
		&MsgDetachFrom{},
	)

	registry.RegisterInterface(
		"lbm.collection.v1.TokenClass",
		(*TokenClass)(nil),
		&FTClass{},
		&NFTClass{},
	)

	registry.RegisterInterface(
		"lbm.collection.v1.Token",
		(*Token)(nil),
		&FT{},
		&OwnerNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
