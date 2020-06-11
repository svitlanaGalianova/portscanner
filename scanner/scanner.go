package scanner

import (
	"net"
	"strconv"
	"time"
)

// Scan scan for the port on host with protocol
// error implies that is closed and unused
func Scan(protocol, hostname string, port int) error {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)

	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
