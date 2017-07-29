package domain

const (
	REL_SELF   RelType = "self"
	REL_CREATE RelType = "create"
	REL_UPDATE RelType = "update"
	REL_DELETE RelType = "delete"
	REL_NEXT   RelType = "next"
	REL_PREV   RelType = "prev"
	REL_FIRST  RelType = "first"
	REL_LAST   RelType = "last"
)

type RelType string

type HalResource interface {
	AddLink(rel RelType, href string)
}

type DefaultHalResource struct {
	Links map[RelType]Link `gorm:"-" json:"_links,omitempty"`
}

func (lr *DefaultHalResource) AddLink(rel RelType, href string) {
	if lr.Links == nil {
		lr.Links = make(map[RelType]Link)
	}

	lr.Links[rel] = Link{
		Href: href,
	}
}

type Link struct {
	Href string `json:"href"`
}
