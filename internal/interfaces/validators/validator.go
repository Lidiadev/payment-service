package validators

type ValidatorInterface interface {
	Validate(interface{}) error
}
