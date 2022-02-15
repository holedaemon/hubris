package types

type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = iota + 1
	ButtonStyleSeconday
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)
