package time

import (
	"fmt"
	"time"
)

func CurrentTime() {

	statement := "The current time is "
	currentTime := time.Now
	fmt.Println(statement + currentTime)
}
