package timetablev1conv

import (
	sharedconv "github.com/twin-te/twin-te/back/handler/api/rpc/shared/conv"
	timetablev1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/timetable/v1"
	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func ToPBTag(tag *timetabledomain.Tag) *timetablev1.Tag {
	pbTag := &timetablev1.Tag{
		Id:     sharedconv.ToPBUUID(tag.ID),
		UserId: sharedconv.ToPBUUID(tag.UserID),
		Name:   tag.Name.String(),
		Order:  int32(tag.Order),
	}
	return pbTag
}
