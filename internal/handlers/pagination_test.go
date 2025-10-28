package handlers

import (
	"net/url"
	"reflect"
	"testing"
)

var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// no limit or offset param
func TestPaginationNoParams(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: ""}, 5, 10)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// no limit or offset param, default limit is 5 but max is 3, so we should only return 3
func TestPaginationNoParamsDefaultLimitGreaterThanMaxLimit(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: ""}, 5, 3)
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// no limit or offset param, default limit is larger than size of array, so we should return everything
func TestPaginationNoParamsDefaultLimitGreaterThanSize(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: ""}, 11, 20)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// limit param only, limit is smaller than size and max, so we should return limit number of items
func TestPaginationLimitParam(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: "?limit=4"}, 1, 20)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// limit param only, limit is larger than size but smaller than max, so we should return everything
func TestPaginationLimitParamLargerThanSize(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: "?limit=12"}, 1, 20)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// limit param only, limit is larger than max, so we should return max number of items
func TestPaginationLimitParamLargerThanMax(t *testing.T) {
	got, _ := paginationRequest(numbers[:], &url.URL{RawQuery: "?limit=15"}, 1, 8)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
