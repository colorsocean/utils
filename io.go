package utils

import (
	"bufio"
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

func LineByLinePipe(cb func(string)) io.Writer {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	pipeId := RandomNumericString(5)
	go func() {
		log.Printf("[pipe:%s] Pipe created", pipeId)
		reader := bufio.NewReader(r)
		msg := ""
		for {
			bytes, ispr, err := reader.ReadLine()

			if err != nil && err != io.EOF {
				panic(err)
			}

			msg += string(bytes)
			log.Print()
			if ispr {
				continue
			}
			line := msg
			msg = ""
			cb(line)

			if err == io.EOF {
				log.Printf("[pipe:%s] Pipe closed", pipeId)
				break
			}
		}
	}()
	return w
}
