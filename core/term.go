package core

import (
	tm "github.com/buger/goterm"
	"strconv"
	"time"
)

func Monitor() {
	tm.Clear()

	for {
		tm.MoveCursor(1, 1)
		tm.Println("Last beat:", time.Now().Format(time.RFC1123))

		res := Check()

		x := 1
		y := 3

		tm.MoveCursor(x, y)

		for key, val := range res {
			tm.Printf(key)
			tm.MoveCursor(35, y)

			color := tm.GREEN
			if val >= 500 && val < 599 {
				color = tm.RED
			} else if val >= 300 && val < 500 {
				color = tm.YELLOW
			}

			tm.Println(tm.Color(strconv.Itoa(val), color))
			y += 1
			tm.MoveCursor(x, y)
		}

		tm.Flush()
		tm.Clear()

		time.Sleep(time.Duration(GetBeatSeconds()) * time.Second)
	}
}
