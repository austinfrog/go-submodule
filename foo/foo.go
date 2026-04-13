package foo

import "github.com/google/uuid"

func Hello() string {
	return "hello from foo/v2: " + uuid.NewString()
}
