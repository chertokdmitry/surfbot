package subscription

var IsSpotId, IsHour bool
var SpotId, Hour int

//type Sub struct {
//	Id    int
//	Title string
//	Hour  string
//}

type Sub struct {
	Id     int    `json:"id"`
	ChatId int64  `json:"chat_id"`
	SpotId int64  `json:"spot_id"`
	Title  string `json:"title"`
	Hour   string `json:"hour"`
}

//type NewSub struct {
//	ChatId	int64 `json:"chat_id"`
//	SpotId	int64 `json:"spot_id"`
//	Hour	string `json:"hour"`
//}
