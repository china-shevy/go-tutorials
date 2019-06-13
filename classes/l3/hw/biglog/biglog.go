package biglog

import (
	"time"
	"encoding/json"
	"io"
	"fmt"
	"math/rand"
)
// StartServer returns a started logger.
func StartServer() logger {
	rand.Seed(time.Now().UnixNano())
	return logger{
		data: make(map[int64]record),
		nextID: func() int64 {
			return time.Now().Unix() + rand.Int63()
		},
	}
}

type record map[string]interface{}

type logger struct {
	data map[int64]record
	nextID func() int64
}

func (l *logger) Log(data io.Reader) error {	
	id := l.nextID()
	r := make(record)
	err := json.NewDecoder(data).Decode(&r)
	if err != nil {
		return err // todo: handle it
	}
	l.data[id] =  r
	fmt.Println(l.data)
	return nil
}

// Record is used for user to define their own filter function.
type Record map[string]interface{}

func (l *logger) All(f func(Record) bool) (r map[int64]Record, err error) {
	r = make(map[int64]Record)
	for id, v := range l.data {
		if f(Record(v)) {
			r[id] = Record(v)
		}
	}
	return
}


