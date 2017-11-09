package observable

// Observer interface
type Observer interface {
	// Identifier each Observer should have a unique identifer
	Identifier() string

	// Update is called when sending an Observer a message.
	//
	// This method is called by the Observable, once for each Observer, for
	// every message that is being delivered.
	//
	// If this is a performance issue the Observer could chose to handle the
	// message in a goroutine.
	Update(messages interface{})
}
