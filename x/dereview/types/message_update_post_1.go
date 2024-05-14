package types

func NewMsgUpdatePost1(creator string, title string, body string, id uint64) *MsgUpdatePost1 {
	return &MsgUpdatePost1{
		Creator: creator,
		Title:   title,
		Body:    body,
		Id:      id,
	}
}
