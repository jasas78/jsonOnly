package main

import (
	"fmt"
	"os"
)

// note : change the exe-bin file name to *.cgi and chmod 755,
// then put it into the shared host then all ok.

func main() {
	fmt.Printf("Content-type: text/html\n\n")
	//print "$ENV{REMOTE_ADDR}\n";
	//print "$ENV{REMOTE_PORT}\n";
	REMOTE_ADDR := os.Getenv("REMOTE_ADDR")
	REMOTE_PORT := os.Getenv("REMOTE_PORT")
	fmt.Printf("addr:%s\n\nport:%s\n\n",
		REMOTE_ADDR,
		REMOTE_PORT)

}
