package main

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:9003", Path: "/ws/test"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	// ตรวจสอบว่ามีการเชื่อมต่ออยู่แล้วหรือไม่
	// if existingConn, exists := connections[refKey]; exists {

	// var conn *websocket.Conn
	// var err error

	// for count < 10 {

	err = conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(int(0))))
	if err != nil {
		return
	}

	// _, message, err := conn.ReadMessage()
	_, _, err = conn.ReadMessage()
	if err != nil {
		return
	}

	startTime := time.Now()
	endTime := startTime.Add(5 * time.Second)
	count := 0
	countSuccess := 0
	countFail := 0

	// timestamp := time.Now().UnixNano()

	logrus.Info("start")
	for time.Now().Before(endTime) {
		count++

		timestamp := time.Now().UnixNano()
		err = conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(int(timestamp))))
		if err != nil {
			countFail++
			continue
		}

		// _, message, err := conn.ReadMessage()
		_, _, err := conn.ReadMessage()
		if err != nil {
			countFail++
			continue
			// log.Println("read:", err)
			// log.Fatalf("Did not read: %v", err)
		}

		// messageStr := string(message[:])
		// timestampStart, _ := strconv.ParseInt(messageStr, 10, 64)
		// timestampEnd := time.Now().UnixNano()
		// nanosecond := timestampEnd - timestampStart
		// millisecond := float64(timestampEnd-timestampStart) / float64(1000000)

		// logrus.Info("nanosecond: ", nanosecond)
		// logrus.Info("millisecond: ", millisecond)
		countSuccess++
	}

	logrus.Info("end of websocket service")
	logrus.Info("count: ", count)
	logrus.Info("countSuccess: ", countSuccess)
	logrus.Info("countFail: ", countFail)
}
