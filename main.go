package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const pidFile = "hup-catcher.pid"

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGHUP)

	err := writePidfile()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("awaiting SIGHUP")

	for {
		select {
		case sig := <-sigs:
			log.Println(sig)
		}
	}
}

func writePidfile() error {
	pidString := fmt.Sprintf("%d", os.Getpid())
	pidBytes := []byte(pidString)
	err := ioutil.WriteFile(pidFile, pidBytes, 0644)
	if err != nil {
		return err
	}
	log.Printf("pid written to %s\n", pidFile)
	return nil
}
