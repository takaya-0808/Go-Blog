package usecase

import (
	"testing"
)

func TestSearch(t *testing.T) {

	test_name := "hoge"
	user, err := userUseCase.Search(test_name)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if user != nil {
		t.Fatalf("failed test")
	}
}
