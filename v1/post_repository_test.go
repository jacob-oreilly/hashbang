package v1

import (
	"reflect"
	"testing"
)

func TestPostRepository(t *testing.T) {
	t.Run("initially empty", func(t *testing.T) {
		got := NewPostRepository().FindAll()
		want := []Post{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("create and get post", func(t *testing.T) {
		postRepository := NewPostRepository()

		postRepository.Create(
			Post{
				ImageUri: "https://images.unsplash.com/photo-1603316851229-26637b4bd1b8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1400&q=80",
				Tags: []string{
					"#windy",
					"#tdd",
				},
			},
		)

		got := postRepository.FindAll()
		want := []Post{
			{
				ImageUri: "https://images.unsplash.com/photo-1603316851229-26637b4bd1b8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1400&q=80",
				Tags: []string{
					"#windy",
					"#tdd",
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}