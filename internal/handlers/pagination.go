package handlers

import (
	"errors"
	"net/url"
	"strconv"
)

type PagedResponse struct {
	PageSize     int   `json:"pageSize"`
	Page         int   `json:"page"`
	ResultSet    []any `json:"resultSet"`
	TotalPages   int   `json:"totalPages"`
	TotalResults int   `json:"totalResults"`
}

func ParsePaginationParams(u *url.URL, defaultLimit int, maxLimit int) (limit int, offset int, err error) {
	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return 0, 0, errors.New("Invalid query parameters")
	}

	if params["limit"] != nil {
		limit, err = strconv.Atoi(params["limit"][0])
		// limit param can only be specified once and must be > 0
		if err != nil || limit < 1 || len(params["limit"]) > 1 {
			return 0, 0, errors.New("Invalid limit parameter")
		}
	} else {
		// set default limit so we don't return everything when user doesn't specify limit
		limit = defaultLimit
	}

	if params["offset"] != nil {
		offset, err = strconv.Atoi(params["offset"][0])
		// offset param can only be specified once and must be > 0
		if err != nil || offset < 1 || len(params["offset"]) > 1 {
			return 0, 0, errors.New("Invalid offset parameter")
		}
	}

	// cap limit to maxLimit
	limit = min(limit, maxLimit)

	return limit, offset, nil
}

func PaginateRespons[T any](array []T, limit, offset int) []T {
	size := len(array)

	i := min(size, offset)
	j := min(size, offset+limit)

	return array[i:j]
}
