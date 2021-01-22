package middleware

// ContextKey is just an alias for string to be used
// as key when assign a value in context.
// This is to avoid go-lint warning
// "should not use basic type untyped string as key in context.WithValue".
type ContextKey string
