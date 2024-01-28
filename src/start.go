package microserviceStarter

// Service contains the dependencies
type Service struct {
	routes *Routes
}

// NewService contains all the logic to initialize the service
func NewService(routes *Routes) *Service {
	return &Service{
		routes: routes,
	}
}

// Start starts the application
func (service *Service) Start() error {
	service.routes.Serve()

	return nil
}
