package app

import "time"

type Service interface {
	GetTimeByCity(chatID int64, city string) (string, error)
	GetStats(chatID int64) (string, error)
	GetLocation(chatID int64, city string) (string, error)
	//GetRegions() (string, error)
}

type Database interface {
	AddEntry(chatID int64, city string, time time.Time) error
	GetEntriesCount(chatID int64) (int64, error)
	GetFirstEntryDate(chatID int64) (time.Time, error)
	GetLastEntryDate(chatID int64) (time.Time, error)
}

type Bot interface {
	Start() error
	Shutdown()
}
