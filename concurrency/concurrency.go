package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// Quick summary

// - Go Routines, the basic unit of concurrency in Go, which let us
// check more than one website at the same time.

// - Anonymous Functions, which we used to start each of the
// concurrent processes that check websites.

// - Channels, to help organize and control the communication
// between the different processes, allowing us to avoid a race
// condition bug.

// - The Race Detector which helped us debug problems with concurrent code
