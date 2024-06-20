package concurrency

type WebsiteChecker func(string) bool

// anon values, useful when it's hard to know what to name a value
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// anonymous function declaration
		// By giving each anonymous function a parameter for the url - u - and then calling the anonymous function with the url as the argument,
		// we make sure that the value of u is fixed as the value of url for the iteration of the loop that we're launching the goroutine in.
		go func(u string) {
			// sending a result struct for each call to wc to the resultChannel.
			// send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	// none of the goroutines started had enough time to add their result to the results map
	//time.Sleep(2 * time.Second)

	for i := 0; i < len(urls); i++ {
		// receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

// Channels: Go data structure that can both receive and send values.
// These operations, along with their details, allow communication between different processes.
/*
By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time. Although each of the calls of wc, and each send to the result channel, is happening concurrently inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.
*/
