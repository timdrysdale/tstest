package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("host", "relay.practable.io", "wss host")
var path = flag.String("path", "/out/pend01/video", "path to video stream")
var file = flag.String("file", "test.ts", "ts file")
var duration = flag.Int("duration", 60, "duration in seconds")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: *path}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	writeToFile := false
	// avoid shadowing in if block by declaring in advance
	var f *os.File

	if *file != "" {
		f, err = os.Create(*file)
		if err == nil {
			writeToFile = true
		}
		defer f.Close()
	}

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if writeToFile {
				_, err := f.Write(message)
				if err != nil {
					writeToFile = false //assume no more writing possible
				}
			} else {
				log.Printf("recv: %s", message)
			}
		}
	}()

	select {
	case <-interrupt:
	case <-time.After(time.Duration(*duration) * time.Second):
	}

}
