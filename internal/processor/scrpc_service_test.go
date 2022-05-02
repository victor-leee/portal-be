package processor

import (
	"fmt"
	"testing"
)

func Test_BuildServiceKey(t *testing.T) {
	completePath := "police-repo-com"
	fmt.Printf("%s", buildServiceKey(completePath))
}
