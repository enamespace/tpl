package store

type Factory interface {
	User() UserStore
}

var client Factory

func SetClient(factory Factory) {
	client = factory
}

func Client() Factory {
	return client
}
