package ctx

type respondToPrintTextIface interface {
	PrintText()
}

type Context struct {
	A respondToPrintTextIface
	B respondToPrintTextIface
}
