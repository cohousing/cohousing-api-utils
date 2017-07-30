package utils

import (
	"fmt"
	"github.com/cohousing/cohousing-api-utils/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestResource struct {
	ID    int64
	Valid bool
	domain.DefaultHalResource
}

func dummyLinkFactory(c *gin.Context, halResource domain.HalResource, basePath string, detailed bool) {
	tr := halResource.(*TestResource)
	tr.Valid = true
	tr.AddLink(domain.REL_SELF, fmt.Sprintf("%s/resource/%d", basePath, tr.ID))
}

func TestAddLinkFactory(t *testing.T) {
	assert.Empty(t, factory)

	AddLinkFactory(&TestResource{}, dummyLinkFactory)

	assert.Equal(t, 1, len(factory))

	c := gin.Context{}
	tr := TestResource{
		ID: 1,
	}
	factory["*utils.TestResource"](&c, &tr, "basePath", true)
	assert.True(t, tr.Valid)

	factory = make(map[string]LinkFactory)
}

func TestAddLinkFactory_IncorrectResource(t *testing.T) {
	assert.Empty(t, factory)

	assert.Panics(t, func() {
		AddLinkFactory(TestResource{}, dummyLinkFactory)
	})

	factory = make(map[string]LinkFactory)
}

func TestAddLinksSingleObject(t *testing.T) {
	AddLinkFactory(&TestResource{}, dummyLinkFactory)

	c := gin.Context{}
	tr := TestResource{
		ID: 1,
	}
	AddLinks(&c, &tr, "/basepath", true)

	assert.Equal(t, "/basepath/resource/1", tr.Links[domain.REL_SELF].Href)

	factory = make(map[string]LinkFactory)
}

func TestAddLinksArray(t *testing.T) {
	AddLinkFactory(&TestResource{}, dummyLinkFactory)

	c := gin.Context{}
	trl := []TestResource{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
		{
			ID: 3,
		},
	}

	AddLinks(&c, &trl, "/basepath", true)

	assert.Equal(t, "/basepath/resource/1", trl[0].Links[domain.REL_SELF].Href)
	assert.Equal(t, "/basepath/resource/2", trl[1].Links[domain.REL_SELF].Href)
	assert.Equal(t, "/basepath/resource/3", trl[2].Links[domain.REL_SELF].Href)

	factory = make(map[string]LinkFactory)
}
