package donationdomain_test

import (
	"github.com/google/uuid"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

func newUserID() idtype.UserID {
	return idtype.UserID(uuid.New())
}
