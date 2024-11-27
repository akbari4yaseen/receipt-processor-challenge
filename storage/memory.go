package storage

import "sync"

var receiptStore = make(map[string]int)
var mutex = &sync.Mutex{}

// SaveReceipt stores the receipt ID and points
func SaveReceipt(id string, points int) {
	mutex.Lock()
	defer mutex.Unlock()
	receiptStore[id] = points
}

// GetPoints retrieves the points for a given receipt ID
func GetPoints(id string) (int, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	points, exists := receiptStore[id]
	return points, exists
}
