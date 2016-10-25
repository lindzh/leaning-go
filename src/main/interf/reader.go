package interf

type Reader interface {
	Read(buf []byte) (n int, err error)
}
