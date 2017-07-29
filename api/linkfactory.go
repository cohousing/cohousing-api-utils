package utils

import (
	"github.com/cohousing/cohousing-api-utils/domain"
	"github.com/gin-gonic/gin"
	"reflect"
)

var (
	factory              map[string]LinkFactory = make(map[string]LinkFactory)
	halResourceInterface domain.HalResource
)

type LinkFactory func(c *gin.Context, halResource domain.HalResource, basePath string, detailed bool)

func AddLinkFactory(resource interface{}, linkFactory LinkFactory) {
	resourceType := reflect.TypeOf(resource)
	if resourceType.Implements(reflect.TypeOf(&halResourceInterface).Elem()) {
		factory[resourceType.String()] = linkFactory
	} else {
		panic("Must add a resource that implements HalResource. Also remember to parse it as a pointer.")
	}
}

// Add links to resource based on the type of it
func AddLinks(c *gin.Context, resource interface{}, basePath string, detailed bool) {
	addLinks := func(halResource domain.HalResource) {
		linkFactory := factory[reflect.TypeOf(halResource).String()]
		linkFactory(c, halResource, basePath, detailed)
	}

	resourceType := reflect.TypeOf(resource)
	if resourceType.Implements(reflect.TypeOf(&halResourceInterface).Elem()) {
		addLinks(resource.(domain.HalResource))
	} else if resourceType.Kind() == reflect.Ptr {
		AddLinks(c, reflect.ValueOf(resource).Elem().Interface(), basePath, detailed)
	} else if resourceType.Kind() == reflect.Slice {
		valueList := reflect.ValueOf(resource)
		listLength := valueList.Len()
		for i := 0; i < listLength; i++ {
			object := valueList.Index(i).Addr().Interface()
			addLinks(object.(domain.HalResource))
		}
	}
}
