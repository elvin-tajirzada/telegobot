package telegobot

import "log"

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
