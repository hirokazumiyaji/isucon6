package main

import (
	"context"
	"time"
)

type Entry struct {
	ID          int
	AuthorID    int
	Keyword     string
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time

	Html  string
	Stars []Star
}

type User struct {
	ID        int
	Name      string
	Salt      string
	Password  string
	CreatedAt time.Time
}

type Star struct {
	ID        int       `json:"id"`
	Keyword   string    `json:"keyword"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type EntryWithCtx struct {
	Context context.Context
	Entry   Entry
}

type Keyword struct {
	Value  string
	Length int
	Hash   string
}

type KeywordList []Keyword

func (k KeywordList) Len() int {
	return len(k)
}

func (k KeywordList) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KeywordList) Less(i, j int) bool {
	return k[i].Length < k[j].Length
}
