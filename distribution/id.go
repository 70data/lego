package distribution

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

func GetUUID() string {
	tempID := fmt.Sprintf("%s", uuid.NewV4())
	uuID := strings.Replace(tempID, "-", "", -1)
	return uuID
}
