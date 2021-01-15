package response

type CheckComment struct {
	ID    int  `json:"id"`
	Check bool `json:"check"`
}
type ReplyComment struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
}
