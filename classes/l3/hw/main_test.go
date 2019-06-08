package test

import (
	"strings"

	"testing"
	
	"github.com/stretchr/testify/require"

	"hw/biglog" // You need to implement a biglog package.
)

// In this challenge, you need to implement a log system. It basically has 2 API, one for log and another for search.
func TestChallenge1(t *testing.T) {
	server := biglog.StartServer() // starts a server in its own goroutine and return the server.
	defer func() {
		err := server.Close()
		require.NoError(t, err)
	}()

	// The only requirement is that it accepts an io.Reader. You need to decide the storage mechanism.
	// 唯一的要求是 Log 函数接受一个 io.Reader。你自己需要决定储存机制。
	// Question:
	//		Should it accept io.ReadCloser instead?
	err := server.Log(strings.NewReader(`{"time": "2019-01-06", "number": 123}`)); require.NoError(t, err)	// Notice this log is the same.
	err := server.Log(strings.NewReader(`{"time": "2019-01-06", "number": 123}`)); require.NoError(t, err)	// Should you keep duplication?
	err = server.Log(strings.NewReader(`{"time": "2019-01-03", "number": 126}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-02", "number": 127}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-01", "number": 128}`)); require.NoError(t, err)

	// You should define a search API. It could have any signature.
	// Find the record between 01-01 and 01-03 and return in logging order.
	// Question:
	//		Should it return errors? If so, what kind?
	logs := server.Search(?)
	require.Equal(t, logs, []string{
		`{"time":"2019-01-03","number":126}`,  // Notice that JSON format is compact
		`{"time":"2019-01-02","number":127}`,
		`{"time":"2019-01-01","number":128}`
	})

	// Find the record between 01-01 and 01-03, select 'number' only and return in logging order
	logs := server.Search(?)
	require.Equal(t, logs, []string{
		`{"number":126}`,
		`{"number":127}`,
		`{"number":128}`
	})

	// Question:
	//		Because this is a log system, should you trust the 'time' which users input? Should the system have its own time?
	//		因为这是一个日志系统，你应该相信用户输入的时间吗？系统需要自己记录时间吗？ 
}
