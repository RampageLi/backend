package main

type HistoryEvent struct {
	id      string `json:"id"`
	content string `json:"content"`
	stTime  uint64 `json:"st_time"`
	edTime  uint64 `json:"ed_time"`
}

func main() {
	getEventBrief()
}

func getEventBrief() {

}
