package sendingtcpdata_test

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestTemporaryError(t *testing.T) {
	var (
		err error
		n   int
		i   = 7 // maximum number of retries
	)

	conn, err := net.Dial("tcp", "10.0.0.1:80")
	if err != nil {
		t.Fatal(err)
	}

	for ; i > 0; i-- {
		n, err = conn.Write([]byte("hello world"))
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
				log.Println("temporary error:", nErr)
				time.Sleep(1 * time.Millisecond)
				continue
			}
			log.Println(err)
		}
		break
	}

	if i == 0 {
		log.Println("temporary write failure threshold exceeded")
	}

	log.Printf("wrote %d bytes to %s\n", n, conn.RemoteAddr())
}
