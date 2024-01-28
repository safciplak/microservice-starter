package providers

// InterfaceProviderService is the interface which each provider needs to comply to
type InterfaceProviderService interface {
	Initialize() error
	GetName() string
}
