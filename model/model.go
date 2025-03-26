package model

type Event struct {
	Type    string  `json:"event_type"` //时间类型
	Repo    string  `json:"repo"`       //操作的仓库
	Actor   string  `json:"actor"`      //即username
	Payload Payload `json:"payload"`
}

type Payload struct {
	Ref_type string `json:"ref_type"` //操作的对象
	Action   string `json:"action"`   //如何操作
}

func FetchTheRecentActivity(username string) error {
	return nil
}
