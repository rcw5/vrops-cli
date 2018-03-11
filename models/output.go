package models

import "io"

type Output struct {
	Type   string
	Writer io.Writer
}
