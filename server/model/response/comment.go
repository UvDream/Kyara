package response

type CheckComment struct {
	ID    int  `json:"id"`
	Check bool `json:"check"`
}
