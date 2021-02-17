package parser

import (
	"goTest/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "123")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+ "elementl but was %v", result.Items)
		profile := result.Items[0].(model.Profile)
		expected := model.Profile{
			Name:       "",
			Gender:     "",
			Age:        0,
			Height:     0,
			Weight:     0,
			Income:     "",
			Marriage:   "",
			Edoucation: "",
			Occupation: "",
			Hokou:      "",
			Xinzuo:     "",
			House:      "",
			Cat:        "",
		}

		if profile != expected {
			t.Errorf("expected %v; but was %v", expected, profile)
		}
	}
}
