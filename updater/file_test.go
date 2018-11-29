package updater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EnumeFilesRecursive(t *testing.T) {

	stopCh := make(chan struct{})

	f, e := EnumFilesRecursive(`..\`, stopCh)
	for path := range f {
		fmt.Println(path)
	}

	err := <-e
	fmt.Println(err)

	assert.True(t, err == nil)

}
