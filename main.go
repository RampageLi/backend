package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryEvent struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	StTime uint64 `json:"stTime"`
	EdTime uint64 `json:"edTime"`
}

type CountryInfo struct {
	ID            string         `json:"id"`
	Country       string         `json:"country"`
	HistoryEvents []HistoryEvent `json:"historyEvents"`
}

type Response struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []CountryInfo `json:"data,omitempty"`
}

var ChinaEvents = []HistoryEvent{
	{ID: "1", Title: "event1", StTime: 1591501341000, EdTime: 1592192541000},
	{ID: "2", Title: "event2", StTime: 1271302941000, EdTime: 1271734941000},
	{ID: "3", Title: "event3", StTime: 1113968541000, EdTime: 1114832541000},
}

var UsEvents = []HistoryEvent{
	{ID: "4", Title: "event4", StTime: 1654584530000, EdTime: 1655016530000},
	{ID: "5", Title: "event5", StTime: 1355294930000, EdTime: 1355554130000},
	{ID: "6", Title: "event6", StTime: 1200379730000, EdTime: 1201243730000},
}

var ChinaInfo = CountryInfo{
	ID:            "100",
	Country:       "China",
	HistoryEvents: ChinaEvents,
}

var UsInfo = CountryInfo{
	ID:            "200",
	Country:       "USA",
	HistoryEvents: UsEvents,
}

func main() {
	router := gin.Default()
	router.GET("/get_country_events", func(c *gin.Context) {
		ret := getCountryEvents(c)
		if !ret.Success {
			c.JSON(http.StatusBadRequest, ret)
		}
		c.JSON(http.StatusOK, ret)
	})

	router.Run(":8881")
}

func getCountryEvents(c *gin.Context) Response {
	stTime, err := strconv.ParseUint(c.Query("stTime"), 10, 64)
	if err != nil {
		return Response{
			Success: false,
			Message: "Please check stTime",
		}
	}

	edTime, err := strconv.ParseUint(c.Query("edTime"), 10, 64)
	if err != nil {
		return Response{
			Success: false,
			Message: "Please check edTime",
		}
	}

	events := []CountryInfo{}

	ChinaInfo := CountryInfo{
		ID:            "100",
		Country:       "China",
		HistoryEvents: []HistoryEvent{},
	}
	for _, event := range ChinaEvents {
		if checkTimeRage(stTime, edTime, event) {
			ChinaInfo.HistoryEvents = append(ChinaInfo.HistoryEvents, event)
		}
	}
	if len(ChinaInfo.HistoryEvents) > 0 {
		events = append(events, ChinaInfo)
	}

	UsInfo := CountryInfo{
		ID:            "200",
		Country:       "USA",
		HistoryEvents: []HistoryEvent{},
	}
	for _, event := range UsEvents {
		if checkTimeRage(stTime, edTime, event) {
			UsInfo.HistoryEvents = append(UsInfo.HistoryEvents, event)
		}
	}
	if len(UsInfo.HistoryEvents) > 0 {
		events = append(events, UsInfo)
	}

	return Response{
		Success: true,
		Message: "",
		Data:    events,
	}
}

func checkTimeRage(stTime, edTime uint64, event HistoryEvent) bool {
	eventStTime := event.StTime
	eventEdTime := event.EdTime
	if eventStTime >= stTime && eventEdTime <= edTime {
		return true
	}
	return false
}
