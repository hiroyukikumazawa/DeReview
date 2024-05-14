package types

func NewMsgCreatePost1(creator string, title string, body string) *MsgCreatePost1 {
	return &MsgCreatePost1{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}
