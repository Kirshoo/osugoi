package options

import (
	"slices"
	"net/url"
)

func Filter(allowed []string, query *url.Values) {
	for paramKey := range *query {
		if !slices.Contains(allowed, paramKey) {
			query.Del(paramKey)	
		}	
	}
}
