package usecase

import (
	"testing"
)

type NameType string

func TestSearch(t *testing.T) {

	var name = "hoge"
	test_name := NameType(name)
	user, err := userUseCase.Search(test_name)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if user != nil {
		t.Fatalf("failed test")
	}
}
