package common

type InvalidRessourceError struct {
	Err error
}

func (i InvalidRessourceError) Error() string {
	return i.Err.Error()
}
