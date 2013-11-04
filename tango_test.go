package tango

import (
	"tango"
 	"fmt"
)

func Example200OK() {
	fmt.Println(tango.FetchMeaning("200"))
  // Output: OK
}

func Example500InternalServerError() {
	fmt.Println(tango.FetchMeaning("500"))
  // Output: Internal Server Error
}

func Example000Unknown() {
	fmt.Println(tango.FetchMeaning("000"))
  // Output: Unknown
}
