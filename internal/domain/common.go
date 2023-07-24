package domain

type Bridger interface {
	None()
}

func Bridge[T Bridger](collectionName string, bridges map[string]Bridger) T {
	return bridges[collectionName].(T)
}
