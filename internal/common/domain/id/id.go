package id

type ID interface {
	Next() (int64, error)
}
