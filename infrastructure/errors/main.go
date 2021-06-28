package errors

type Causer interface {
	Cause() error
}

func Cause(err error) error {
	for err != nil {
		cause, ok := err.(Causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

type Wrapper interface {
	Unwrap() error
}

func Unwrap(err error) error {
	if unwrapper, ok := err.(Wrapper); ok {
		return unwrapper.Unwrap()
	}

	return nil
}

type Contextor interface {
	Context() map[string]interface{}
}

func Context(err error) map[string]interface{} {
	if contextor, ok := err.(Contextor); ok {
		return contextor.Context()
	}

	return make(map[string]interface{})
}
