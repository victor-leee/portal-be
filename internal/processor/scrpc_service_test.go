package processor

import (
	"fmt"
	"testing"
)

func Test_BuildServiceKey(t *testing.T) {
	completePath := "police-github-com"
	fmt.Printf("%s", buildServiceKey(completePath))
}
