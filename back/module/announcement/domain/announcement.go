package announcementdomain

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twin-te/twin-te/back/base"
	shareddomain "github.com/twin-te/twin-te/back/module/shared/domain"
	"github.com/twin-te/twin-te/back/module/shared/domain/idtype"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=AnnouncementTag -trimprefix=AnnouncementTag -output=announcement_tag_string.gen.go
type AnnouncementTag int

const (
	AnnouncementTagInformation AnnouncementTag = iota + 1
	AnnouncementTagNotification
)

var AllAnnouncementTags = []AnnouncementTag{
	AnnouncementTagInformation,
	AnnouncementTagNotification,
}

func ParseAnnouncementTag(s string) (AnnouncementTag, error) {
	ret, ok := base.FindByString(AllAnnouncementTags, s)
	if ok {
		return ret, nil
	}
	return 0, fmt.Errorf("failed to parse AnnouncementTag %v", s)
}

var (
	ParseTitle   = shareddomain.NewRequiredStringParser("title")
	ParseContent = shareddomain.NewRequiredStringParser("content")
)

// Announcement is identified by one of the following fields.
//   - ID
type Announcement struct {
	ID          idtype.AnnouncementID
	Tags        []AnnouncementTag
	Title       shareddomain.RequiredString
	Content     shareddomain.RequiredString
	PublishedAt time.Time
}

func (a *Announcement) Clone() *Announcement {
	ret := lo.ToPtr(*a)
	ret.Tags = base.CopySlice(a.Tags)
	return ret
}

func ConstructAnnouncement(fn func(a *Announcement) (err error)) (*Announcement, error) {
	a := new(Announcement)
	if err := fn(a); err != nil {
		return nil, err
	}

	if a.ID.IsZero() || a.Title.IsZero() || a.Content.IsZero() || a.PublishedAt.IsZero() {
		return nil, fmt.Errorf("failed to construct %+v", a)
	}

	return a, nil
}

func ParseAnnouncement(id string, tags []string, title, content string, publishedAt time.Time) (announcement *Announcement, err error) {
	return ConstructAnnouncement(func(a *Announcement) (err error) {
		a.ID, err = idtype.ParseAnnouncementID(id)
		if err != nil {
			return err
		}

		a.Tags, err = base.MapWithErr(tags, ParseAnnouncementTag)
		if err != nil {
			return err
		}

		a.Title, err = ParseTitle(title)
		if err != nil {
			return err
		}

		a.Content, err = ParseContent(content)
		if err != nil {
			return err
		}

		a.PublishedAt = publishedAt

		return nil
	})
}
