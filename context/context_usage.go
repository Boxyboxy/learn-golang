package contextusage

import (
	"context"
	"fmt"
	"net/http"
)

/* Not idiomatic (denotaing expressions natural to native speaker)


From the go doc:

Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

From the Go Blog: Context again:

At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.
*/

/*
	type Store interface {
		Fetch() string
		Cancel()
	}

	func Server(store Store) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			data := make(chan string, 1)

			go func() {
				data <- store.Fetch()
			}()

			select {
			case d := <-data:
				fmt.Fprint(w, d)
			case <-ctx.Done():
				// context has a method Done() which reutrns a channel which gets sent a signal when the context is done or cancelled
				// we want to listen to the signal and call store.Cancel if we get it but ignore it if our store manages to Fetch before it.
				store.Cancel()
			}

		}
	}
*/
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}
