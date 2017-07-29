package utils

import (
	"fmt"
	"github.com/cohousing/cohousing-api-utils/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginatedList struct {
	CurrentPage int
	Count       int
	Objects     interface{} `json:"objects"`
	domain.DefaultHalResource
}

func CreatePaginatedList(baseUrl string, objects interface{}, currentPage, count, recordsPerPage int) PaginatedList {
	objectList := PaginatedList{
		CurrentPage: currentPage,
		Count:       count,
		Objects:     objects,
	}

	AddPaginationLinks(&objectList, baseUrl, recordsPerPage)

	return objectList
}

func GetCurrentPage(c *gin.Context) int {
	var page int
	if pageInt, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32); err == nil {
		page = int(pageInt)
	}
	if page < 1 {
		page = 1
	}
	return page
}

func AddPaginationLinks(objectList *PaginatedList, baseUrl string, recordsPerPage int) {
	var firstPage int = 1

	var lastPage int = objectList.Count / recordsPerPage
	if lastPage < 1 {
		lastPage = 1
	}

	var prevPage int = objectList.CurrentPage - 1
	var nextPage int = objectList.CurrentPage + 1

	objectList.AddLink(domain.REL_SELF, generatePaginationUrl(baseUrl, objectList.CurrentPage))

	if firstPage != lastPage {
		objectList.AddLink(domain.REL_FIRST, generatePaginationUrl(baseUrl, firstPage))
		objectList.AddLink(domain.REL_LAST, generatePaginationUrl(baseUrl, lastPage))
	}
	if prevPage >= firstPage && prevPage < lastPage {
		objectList.AddLink(domain.REL_PREV, generatePaginationUrl(baseUrl, prevPage))
	}
	if nextPage <= lastPage {
		objectList.AddLink(domain.REL_NEXT, generatePaginationUrl(baseUrl, nextPage))
	}
}

func generatePaginationUrl(baseUrl string, page int) string {
	if page > 1 {
		return fmt.Sprintf("%s?page=%d", baseUrl, page)
	} else {
		return baseUrl
	}
}

func GetStartRecord(currentPage, recordsPerPage int) int {
	return (currentPage - 1) * recordsPerPage
}
