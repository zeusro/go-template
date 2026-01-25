package local

import (
	"fmt"
	"sync"
	"time"
)

// ShowCurrentTimeZone prints the current time zone and time in a friendly format.
func ShowCurrentTimeZone() {
	// godotenv.Load("../../.env")
	now := time.Now()
	zoneName, offset := now.Zone()
	offsetHours := offset / 3600
	offsetMinutes := (offset % 3600) / 60

	fmt.Println("🕒 当前时间信息:")
	fmt.Printf("📍 时区: %s (UTC%+02d:%02d)\n", zoneName, offsetHours, offsetMinutes)
	fmt.Printf("📅 当前时间: %s\n", now.Format("2006-01-02 15:04:05 MST"))
}

func GetCurrentTime(limit int) {
	m := make(map[string]string, 1000)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < len(m); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			k := fmt.Sprintf("key%d", i)
			m[k] = fmt.Sprintf("value%d", i)
		}(limit)
		wg.Wait()
	}

	//读取
	mu.Lock()
	for k, v := range m {
		fmt.Println(k, v)
	}
	mu.Unlock()
	fmt.Println("done", len(m))
}
