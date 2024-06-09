package message

type Repository interface {
	FindMessageStatusByUUID(string) (string, error)
}
