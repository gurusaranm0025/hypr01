package logout

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logout(t *testing.T) {
	// bug logout command not closing after closing the logout window - only occurs in testing
	err := Logout(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
}
