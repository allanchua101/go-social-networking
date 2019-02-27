package models

// Activity struct is used for representing
// instance of activity records stored in the 
// applicaiton.
type Activity struct {
	// The actor performing the activity represented by user name
	Actor string
	// The verb of the activity, please include at 
	// least 3 types. e.g. like, share, post
	Verb string
	// Optional: The object of the activity, e.g. photo:1
	Object string
	// Optional: Target of the activity, usually another user name
	Target string
}
