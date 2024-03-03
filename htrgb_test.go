package htrgb_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/aniket-hpp/htrgb"
)

func TestHexToRGB(t *testing.T) {
	r, g, b, e := htrgb.HexToRGB("#ffffff")

	if e != nil {
		log.Fatal(e.Error())
	}

	fmt.Println(r, g, b)
}
