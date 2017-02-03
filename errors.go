package essentials

// CtxError is an error with some added context.
type CtxError struct {
	Context  string
	Original error
}

// AddCtx creates a CtxError by adding some context
// to an existing error.
// If the original error is nil, then this returns nil.
func AddCtx(ctx string, err error) *CtxError {
	if err == nil {
		return nil
	}
	return &CtxError{Context: ctx, Original: err}
}

// Error returns an error message with added context.
// The message is of the form "Context: Original.Error()".
func (c *CtxError) Error() string {
	return c.Context + ": " + c.Original.Error()
}
