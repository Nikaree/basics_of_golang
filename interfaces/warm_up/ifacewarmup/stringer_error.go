package ifacewarmup

// StringerError оборачивает ошибку и реализует:
//
//	– error (метод Error)
//	– fmt.Stringer (метод String)
//	– Unwrap (для errors.Is / errors.As)
type StringerError struct {
	Err error // может быть nil
}

// Error возвращает Err.Error() или пустую строку, если Err == nil.
func (se *StringerError) Error() string {
	if se.Err == nil {
		return ""
	}
	return se.Err.Error()
}

// String возвращает то же, что и Error().
func (se *StringerError) String() string {
	return se.Err.Error()
}

// Unwrap возвращает исходную ошибку.
func (se *StringerError) Unwrap() error {
	return se.Err
}

// Wrap(err) создает *StringerError; если err == nil — возвращает nil.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &StringerError{Err: err}
}
