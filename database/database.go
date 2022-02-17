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

func (d *db) Exists(key string) bool {
	_, exists := d.dataMap[key]
	return exists
}

func (d *db) Find(key string) (*Data, error) {
	data, exists := d.dataMap[key]
	if exists {
		return &data, nil
	}
	return nil, errors.New("data not found")
}

func (d *db) Create(key string, value string) error {
	if d.Exists(key) {
		return errors.New("data already exits")
	}
	d.dataMap[key] = *FormatData(key, value)
	return nil
}

func (d *db) Update(key string, value string) error {
	if !d.Exists(key) {
		return errors.New("data not found")
	}
	d.dataMap[key] = *FormatData(key, value)
	return nil
}

func (d *db) Delete(key string) {
	if d.Exists(key) {
		delete(d.dataMap, key)
	}
}

func (d *db) GetDataList() []Data {
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

func getCurrentTimestamp() string {
	return time.Now().UTC().Format(timeFormat)
}

func FormatData(key string, value string) *Data {
	newData := &Data{Key: key, Value: value}
	newData.Timestamp = getCurrentTimestamp()
	return newData
}
