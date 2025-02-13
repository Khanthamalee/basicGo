package linkedin

//Work with dates and times

import (
	"fmt"
	"time"
)

func Time() {
	t := time.Date(2009, time.February, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go lanched at %s\n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n)
	fmt.Printf("The object 's type is %T\n", n)
	fmt.Printf(n.Format(time.ANSIC) + "\n")

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow day is %s", tomorrow)
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n")

	format := "Mon 2006-02-01"
	fmt.Printf(tomorrow.Format(format) + "\n")

}
