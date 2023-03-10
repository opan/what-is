package whatis_rpc

import (
	"fmt"
)

type Listener int
type Reply struct {
	Data string
}

func (l *Listener) GetLine(line []byte, reply *Reply) error {
	rv := string(line)

	fmt.Printf("Receive: %v\n", rv)
	*reply = Reply{rv}
	return nil
}
