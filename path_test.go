package path_test

import (
	"fmt"
	"testing"

	"github.com/tempcke/path"
)

func TestPath(t *testing.T) {
	var tests = map[string]struct {
		expect, actual string
	}{
		"simple": {
			expect: "/foo",
			actual: path.New("/foo").String(),
		},
		"base": {
			expect: "http://127.0.0.1:42407/foo/:foo",
			actual: path.New("/foo/:foo").
				WithBaseURL("http://127.0.0.1:42407").String(),
		},
		"base and prefix": {
			expect: "http://127.0.0.1:42407/v1/foo/:foo",
			actual: path.New("/foo/:foo").
				WithBaseURL("http://127.0.0.1:42407").
				WithPrefix("v1").String(),
		},
		"extra slashes removed properly": {
			expect: "http://127.0.0.1:42407/v1/foo/:foo",
			actual: path.New("/foo/:foo/").
				WithBaseURL("http://127.0.0.1:42407/").
				WithPrefix("/v1/").String(),
		},
		"param without :": {
			expect: "/foo/bar/baz",
			actual: path.New("/foo/:foo/baz").
				WithParam("foo", "bar").String(),
		},
		"param with :": {
			expect: "/foo/bar/baz",
			actual: path.New("/foo/:foo/baz").
				WithParam(":foo", "bar").String(),
		},
		"query args": {
			expect: "/foo?a=A&b=B",
			actual: path.New("/foo").
				WithQuery("a", "A").
				WithQuery("b", "B").String(),
		},
		"omit empty query args": {
			// useful when building a path not knowing if a config val has been set or not
			expect: "/foo?a=A",
			actual: path.New("/foo").
				WithQuery("a", "A").
				WithQuery("b", "").String(),
		},
		"keep empty query when no values": {
			// useful when building a path not knowing if a config val has been set or not
			expect: "/foo?a=",
			actual: path.New("/foo").
				WithQuery("a").String(),
		},
		"query with multiple values": {
			expect: "/foo?id=A&id=B",
			actual: path.New("/foo").
				WithQuery("id", "A", "B").String(),
		},
		"query with multiple values via multiple WithQuery": {
			expect: "/foo?id=A&id=B",
			actual: path.New("/foo").
				WithQuery("id", "A").
				WithQuery("id", "B").String(),
		},
		"multiple params": {
			expect: "/foo/abc/bar/def",
			actual: path.New("/foo/:foo/bar/:bar").WithParams(map[string]string{
				"foo": "abc",
				"bar": "def",
			}).String(),
		},
		"all features": {
			expect: "http://127.0.0.1:42407/v1/foo/p1/bar/p2/p3?a=A&b=B&c=C&id=1&id=2",
			actual: path.New("/foo/:foo/bar/:bar/:baz").
				WithBaseURL("http://127.0.0.1:42407").
				WithPrefix("v1").
				WithParam(":foo", "p1").
				WithParams(map[string]string{
					"bar": "p2",
					"baz": "p3",
				}).
				WithQuery("id", "1", "2").
				WithQuery("a", "A").
				WithQueryArgs(map[string]string{
					"b": "B",
					"c": "C",
				}).
				String(),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.expect != tc.actual {
				t.Errorf("\n want: %s\n got:  %s", tc.expect, tc.actual)
			}
		})
	}
}
func ExamplePath() {
	const pathFoo = "/foo/:foo"
	uri := path.New(pathFoo).
		WithParam(":foo", "bar")
	fmt.Println(uri.String())
	// Output: /foo/bar
}
func ExamplePath_WithQuery() {
	const pathFooBarBaz = "/foo/:foo/bar/:bar/:baz"
	uri := path.New(pathFooBarBaz).
		WithBaseURL("https://example.com").
		WithPrefix("v1").
		WithParam(":foo", "p1").
		WithParams(map[string]string{
			"bar": "p2",
			"baz": "p3",
		}).
		WithQuery("id", "1", "2").
		WithQuery("a", "A").
		WithQueryArgs(map[string]string{
			"b": "B",
			"c": "C",
		})
	fmt.Println(uri.String())
	// Output: https://example.com/v1/foo/p1/bar/p2/p3?a=A&b=B&c=C&id=1&id=2
}
