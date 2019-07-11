package panik

// OnError panics if err != nil
func OnError(err error) {
	if err != nil {
		panic(err)
	}
}
