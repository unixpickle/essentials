package essentials

// CtxError is an error with some added context.
type CtxError struct {
	Context  string
	Original error
}

// AddCtx creates a *CtxError by adding some context
// to an existing error.
// If the original error is nil, then this returns nil.
func AddCtx(ctx string, err error) error {
	if err == nil {
		return nil
	}
	return &CtxError{Context: ctx, Original: err}
}

// AddCtxTo is like doing
//
//     *err = AddCtx(ctx, *err)
//
// It is useful for adding context to a named return
// argument, like in:
//
//     func MyMethod() (err error) {
//         defer essentials.AddCtxTo("MyMethod", &err)
//         // Code here...
//     }
func AddCtxTo(ctx string, err *error) {
	*err = AddCtx(ctx, *err)
}

// Error returns an error message with added context.
// The message is of the form "Context: Original.Error()".
func (c *CtxError) Error() string {
	return c.Context + ": " + c.Original.Error()
}
