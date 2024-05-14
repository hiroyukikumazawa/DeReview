package types

func NewMsgCreatePost(creator string, title string, body string) *MsgCreatePost {
	return &MsgCreatePost{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func NewMsgUpdatePost(creator string, id uint64, title string, body string) *MsgUpdatePost {
	return &MsgUpdatePost{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func NewMsgDeletePost(creator string, id uint64) *MsgDeletePost {
	return &MsgDeletePost{
		Id:      id,
		Creator: creator,
	}
}
