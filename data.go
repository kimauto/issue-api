package main

import (
	"errors"
	"sync"
	"time"
)

var (
	users  = make(map[uint]*User)
	issues = make(map[uint]*Issue)
	nextID uint = 1
	lock   sync.Mutex
)

func InitData() {
	users[1] = &User{ID: 1, Name: "김개발"}
	users[2] = &User{ID: 2, Name: "이디자인"}
	users[3] = &User{ID: 3, Name: "박기획"}
}

func CreateNewIssue(title, desc string, userID *uint) (*Issue, error) {
	lock.Lock()
	defer lock.Unlock()

	status := "PENDING"
	var user *User

	if userID != nil {
		u, ok := users[*userID]
		if !ok {
			return nil, errors.New("존재하지 않는 사용자입니다")
		}
		user = u
		status = "IN_PROGRESS"
	}

	issue := &Issue{
		ID:          nextID,
		Title:       title,
		Description: desc,
		Status:      status,
		User:        user,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	issues[nextID] = issue
	nextID++
	return issue, nil
}
