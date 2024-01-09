package models

type LoopBackRequest struct {
	Enable string `json:"enable"`
	Pwd    string `json:"pwd"`
	Conft  string `json:"conft"`
	Loop   string `json:"loop"`
	Ipaddr string `json:"ipaddr"`
	End    string `json:"end"`
	Wr     string `json:"wr"`
}
