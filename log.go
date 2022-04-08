package main

import "log"

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
