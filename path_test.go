package path

import (
	"strings"
	"testing"
)

func TestPathParsing(t *testing.T) {
	cases := map[string]bool{
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":             true,
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a":           true,
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b/c/d/e/f": true,
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":             true,
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a":           true,
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b/c/d/e/f": true,
		"/btns/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b/c/d/e/f": true,
		"/btns/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":             true,
		"QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b/c/d/e/f":       true,
		"QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":                   true,
		"/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":                  false,
		"/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a":                false,
		"/btfs/foo": false,
		"/btfs/":    false,
		"btfs/":     false,
		"btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n": false,
		"/ipld/foo": false,
		"/ipld/":    false,
		"ipld/":     false,
		"ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n": false,
	}

	for p, expected := range cases {
		_, err := ParsePath(p)
		valid := err == nil
		if valid != expected {
			t.Fatalf("expected %s to have valid == %t", p, expected)
		}
	}
}

func TestNoComponents(t *testing.T) {
	for _, s := range []string{
		"/btfs/",
		"/btns/",
		"/ipld/",
	} {
		_, err := ParsePath(s)
		if err == nil || !strings.Contains(err.Error(), "not enough path components") || !strings.Contains(err.Error(), s) {
			t.Error("wrong error")
		}
	}
}

func TestInvalidPaths(t *testing.T) {
	for _, s := range []string{
		"/btfs",
		"/testfs",
		"/",
	} {
		_, err := ParsePath(s)
		if err == nil || !strings.Contains(err.Error(), "invalid btfs path") || !strings.Contains(err.Error(), s) {
			t.Error("wrong error")
		}
	}
}

func TestIsJustAKey(t *testing.T) {
	cases := map[string]bool{
		"QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":           true,
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":     true,
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a":   false,
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b": false,
		"/btns/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":     false,
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b": false,
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":     true,
	}

	for p, expected := range cases {
		path, err := ParsePath(p)
		if err != nil {
			t.Fatalf("ParsePath failed to parse \"%s\", but should have succeeded", p)
		}
		result := path.IsJustAKey()
		if result != expected {
			t.Fatalf("expected IsJustAKey(%s) to return %v, not %v", p, expected, result)
		}
	}
}

func TestPopLastSegment(t *testing.T) {
	cases := map[string][]string{
		"QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":             []string{"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n", ""},
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n":       []string{"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n", ""},
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a":     []string{"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n", "a"},
		"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a/b":   []string{"/btfs/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/a", "b"},
		"/btns/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/x/y/z": []string{"/btns/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/x/y", "z"},
		"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/x/y/z": []string{"/ipld/QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n/x/y", "z"},
	}

	for p, expected := range cases {
		path, err := ParsePath(p)
		if err != nil {
			t.Fatalf("ParsePath failed to parse \"%s\", but should have succeeded", p)
		}
		head, tail, err := path.PopLastSegment()
		if err != nil {
			t.Fatalf("PopLastSegment failed, but should have succeeded: %s", err)
		}
		headStr := head.String()
		if headStr != expected[0] {
			t.Fatalf("expected head of PopLastSegment(%s) to return %v, not %v", p, expected[0], headStr)
		}
		if tail != expected[1] {
			t.Fatalf("expected tail of PopLastSegment(%s) to return %v, not %v", p, expected[1], tail)
		}
	}
}

func TestV0ErrorDueToLowercase(t *testing.T) {
	badb58 := "/btfs/qmbwqxbekc3p8tqskc98xmwnzrzdtrlmimpl8wbutgsmnr"
	_, err := ParsePath(badb58)
	if err == nil {
		t.Fatal("should have failed to decode")
	}
	if !strings.HasSuffix(err.Error(), "(possible lowercased CIDv0; consider converting to a case-agnostic CIDv1, such as base32)") {
		t.Fatal("should have meaningful info about case-insensitive fix")
	}
}
