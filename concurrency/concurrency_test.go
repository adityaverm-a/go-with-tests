package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url == "http://batman.com"
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://batman.com",
		"http://yahoo.com",
	}

	want := map[string]bool{
		"http://google.com": false,
		"http://batman.com": true,
		"http://yahoo.com":  false,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}
