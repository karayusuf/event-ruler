package rule

type Value interface {
	Raw() []byte
	String() (string, error)
}
