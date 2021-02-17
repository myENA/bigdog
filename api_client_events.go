package bigdog

// APIClientEventListener is away to have on-completion API call metadata pushed to a receiver
type APIClientEventListener interface {
	// Push is called each time an API call has been completed.
	Push(APIResponseMeta)
}

type funcEventHandler struct {
	fn func(APIResponseMeta)
}

func NewAPIClientEventListenerFunc(fn func(meta APIResponseMeta)) APIClientEventListener {
	return funcEventHandler{fn: fn}
}

func (ev funcEventHandler) Push(meta APIResponseMeta) {
	ev.fn(meta)
}

type chanEventHandler struct {
	ch chan<- APIResponseMeta
}

func NewAPIClientEventListenerChan(ch chan<- APIResponseMeta) APIClientEventListener {
	return chanEventHandler{ch: ch}
}

func (ev chanEventHandler) Push(meta APIResponseMeta) {
	ev.ch <- meta
}
