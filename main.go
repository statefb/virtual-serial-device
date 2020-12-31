package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
	"strconv"
	"github.com/tarm/serial"
)

func main() {
	var(
		path = flag.String("path", "./conf.yml", "path to config yml file.")
	)
	c := GetConf(*path)
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
			
	val := 25.0
	base := 25.0
	for {
		val = base + rand.Float64()
		_, err = s.Write([]byte(strconv.FormatFloat(val, 'f', -1, 64) + "\r\n"))
		time.Sleep(time.Second * 1.0)
	}
}