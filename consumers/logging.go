package consumers

import (
	"os"
	"fmt"
	"time"
	"sync"
	"encoding/json"

	"github.com/sensorsdata/sa-sdk-go/structs"
)

const (
	CHANNEL_SIZE = 1000
)

type LoggingConsumer struct {
	w       *LogWriter
	Fname   string
	Hour    bool
}

func InitLoggingConsumer(fname string, hour bool) (*LoggingConsumer, error) {
	w, err := InitLogWriter(fname, hour)
	if err != nil {
		return nil, err
	}

	c := &LoggingConsumer{Fname: fname, Hour: hour, w: w}
	return c, nil
}

func (c *LoggingConsumer) Send(data structs.EventData) error {
	return c.w.Write(data)
}

func (c *LoggingConsumer) Flush() error {
	c.w.Flush()
	return nil
}

func (c *LoggingConsumer) Close() error {
	c.w.Close()
	return nil
}

type LogWriter struct {
	rec        chan string

	fname      string
	file       *os.File

	day        int
	hour       int

	hourRotate bool

	wg         sync.WaitGroup
}

func (w *LogWriter) Write(data structs.EventData) error {
	bdata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.rec <- string(bdata)
	return nil
}

func (w *LogWriter) Flush() {
	w.file.Sync()
}

func (w *LogWriter) Close() {
	close(w.rec)
	w.wg.Wait()
}

func (w *LogWriter) intRotate() error {
	fname := ""

	if w.file != nil {
		w.file.Close()
	}

	now := time.Now()
	today := now.Format("2006-01-02")
	w.day = time.Now().Day()

	if w.hourRotate {
		hour := now.Hour()
		w.hour = hour

		fname = fmt.Sprintf("%s.%s.%d", w.fname, today, hour)
	} else {
		fname = fmt.Sprintf("%s.%s", w.fname, today)
	}

	fd, err := os.OpenFile(fname, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return err
	}
	w.file = fd

	return nil
}

func InitLogWriter(fname string, hourRotate bool) (*LogWriter, error) {
	w := &LogWriter{
		fname      : fname,
		day        : time.Now().Day(),
		hour       : time.Now().Hour(),
		hourRotate : hourRotate,
		rec        : make(chan string, CHANNEL_SIZE),
	}

	if err := w.intRotate(); err != nil {
		fmt.Fprintf(os.Stderr, "LogWriter(%q): %s\n", w.fname, err)
		return nil, err
	}

	w.wg.Add(1)

	go func() {
		defer func() {
			if w.file != nil {
				w.file.Sync()
				w.file.Close()
			}
			w.wg.Done()
		}()

		for {
			select {
			case rec, ok := <-w.rec:
				if !ok {
					return
				}

				now := time.Now()

				if (w.hourRotate && now.Hour() != w.hour) ||
					(now.Day() != w.day) {
					if err := w.intRotate(); err != nil {
						fmt.Fprintf(os.Stderr, "LogWriter(%q): %s\n", w.fname, err)
						return
					}
				}

				_, err := fmt.Fprintln(w.file, rec)
				if err != nil {
					fmt.Fprintf(os.Stderr, "LogWriter(%q): %s\n", w.fname, err)
					return
				}
			}
		}
	}()

	return w, nil
}
