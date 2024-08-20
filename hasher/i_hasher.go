package hasher

type IHasher interface {
	Hash([]byte) ([]byte, error)
	Compare([]byte, []byte) error
	Cost([]byte) (int, error)
}
