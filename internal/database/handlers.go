package database

import "time"

func (a adapter) AddEntry(chatID int64, city string, time time.Time) error {
	query := `
insert into entry 
    (chat_id, city, created_at) 
values ($1, $2, $3)`

	_, err := a.db.Exec(query, chatID, city, time)
	if err != nil {
		a.logger.Error().Err(err).Msg("add entry")
		return err
	}

	return nil
}

func (a adapter) GetEntriesCount(chatID int64) (int64, error) {
	var count int64
	query := `
select count(1) 
from entry
where chat_id=$1`

	err := a.db.Get(&count, query, chatID)
	if err != nil {
		a.logger.Error().Err(err).Msg("get entries count")
		return count, err
	}

	return count, nil
}

func (a adapter) GetFirstEntryDate(chatID int64) (time.Time, error) {
	var date time.Time
	query := `
select created_at 
from entry 
where chat_id=$1 
order by created_at asc 
limit 1`

	err := a.db.Get(&date, query, chatID)
	if err != nil {
		a.logger.Error().Err(err).Msg("get first entry date")
		return date, err
	}

	return date, nil
}

func (a adapter) GetLastEntryDate(chatID int64) (time.Time, error) {
	var date time.Time
	query := `
select created_at 
from entry 
where chat_id=$1 
order by created_at desc 
limit 1`

	err := a.db.Get(&date, query, chatID)
	if err != nil {
		a.logger.Error().Err(err).Msg("get first entry date")
		return date, err
	}

	return date, nil
}
