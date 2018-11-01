package learning

import (
	"fmt"
)

func Learning() {
	fmt.Println("learn...")

	func() {
		for i := 0; i < 100; i++ {
			go fmt.Println(i)
		}
	}()

}
