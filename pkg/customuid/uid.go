// customuid implement custom unique id, for moment use ulid
package customuid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

var t = time.Now()
var entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

func GetUniqueID() string {
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return ""
	}
	return id.String()
}
