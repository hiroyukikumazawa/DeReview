package types

func NewMsgDeletePost1(creator string, id uint64) *MsgDeletePost1 {
	return &MsgDeletePost1{
		Creator: creator,
		Id:      id,
	}
}
