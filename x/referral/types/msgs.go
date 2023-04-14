package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgAddReferral)(nil)
	_ sdk.Msg = (*MsgSetReferral)(nil)
)

func NewMsgAddReferral(creator sdk.AccAddress, code string) *MsgAddReferral {
	return &MsgAddReferral{
		Sender: creator.String(),
		Code:   code,
	}
}

func NewMsgSetReferral(sender sdk.AccAddress, parent string) *MsgSetReferral {
	return &MsgSetReferral{
		Sender: sender.String(),
		Parent: parent,
	}
}

func (msg *MsgAddReferral) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg *MsgAddReferral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address %s", err)
	}
	// if err := msg.Coins.Validate(); err != nil {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, err.Error())
	// }
	// if len(msg.Coins) != 2 {
	// 	return ErrWrongCoinNumber
	// }
	// TODO JIHON: mininum len of Code & max len of Code
	if msg.Code == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "code must not be empty")
	}

	return nil
}

func (msg *MsgSetReferral) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg *MsgSetReferral) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if msg.Parent == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "parent must not be empty")
	}
	return nil
}

const (
	TypeMsgAddReferral = "add_referral"
	TypeMsgSetReferral = "set_referral"
)

func (msg *MsgAddReferral) Route() string { return RouterKey }
func (msg *MsgAddReferral) Type() string  { return TypeMsgAddReferral }
func (msg *MsgAddReferral) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSetReferral) Route() string { return RouterKey }
func (msg *MsgSetReferral) Type() string  { return TypeMsgSetReferral }

func (msg *MsgSetReferral) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
