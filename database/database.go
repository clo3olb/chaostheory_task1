package database

import (
	"errors"
	"sort"
	"time"
)

const (
	timeFormat string = "2006-01-02T15:04:05Z"
)

type db struct {
	dataMap map[string]Data
}

type Data struct {
	Timestamp string `json:"timestamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

var theDatabase *db

func DB() *db {
	if theDatabase == nil {
		theDatabase = &db{dataMap: map[string]Data{}}
	}
	return theDatabase
}

func (d *db) Add(newData Data) {
	d.dataMap[newData.Key] = newData
}

func GetDataList(d *db) []Data {
	list := []Data{}
	for _, data := range d.dataMap {
		list = append(list, data)
	}
	return Sort(list)
}

func Sort(list []Data) []Data {
	sort.Slice(list, func(i, j int) bool {
		t1, _ := time.Parse(timeFormat, list[i].Timestamp)
		t2, _ := time.Parse(timeFormat, list[j].Timestamp)
		return t1.Unix() > t2.Unix()
	})
	return list
}

func FormatData(key string, value string) (*Data, error) {
	if key == "" {
		return nil, errors.New("\"key\" is empty")
	}
	if value == "" {
		return nil, errors.New("\"value\" is empty")
	}
	return &Data{Timestamp: time.Now().UTC().Format(timeFormat), Key: key, Value: value}, nil
}
