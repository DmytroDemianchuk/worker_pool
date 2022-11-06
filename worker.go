package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"created record",
	"deleted record",
	"updated account",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	time.Sleep(time.Microsecond * 10)
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	wg := &sync.WaitGroup{}

	users := make(chan User, 1500)
	go generateUsers(100, users)

	for user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}
	wg.Wait()

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(user User, wg *sync.WaitGroup) error {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	wg.Done()
	return nil
}

func generateUsers(count int, users chan User) {
	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@company.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}

	}

	close(users)
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
