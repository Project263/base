package handlers

func returnError(err error) *map[string]string {
	return &map[string]string{
		"error": err.Error(),
	}
}
