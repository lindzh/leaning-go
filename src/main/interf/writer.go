package interf

type Writer interface {
	Write(buf []byte) (n int, err error)
}
