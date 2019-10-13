package sqlBuild

import (
	"fmt"
	"testing"
)

func TestEscape(t *testing.T) {
	fmt.Println(Escape("xs's"))
	fmt.Println(Escape("xs\\'s"))
}
