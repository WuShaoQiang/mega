package vm

import (
	"github.com/WuShaoQiang/mega/model"
)

func UpdateLastSeen(username string) error {
	return model.UpdateLastSeen(username)
}
