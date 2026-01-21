package resources

// resourceOptionSet represents a configuration set for resource options, such as determining if the resource is core.
type resourceOptionSet struct {
	isCore bool
}

// resourceOption represents a configuration option for a resource.
type resourceOption func(*resourceOptionSet)

// WithCore is an option func that sets the resource as a Firebolt Core resource.
func WithCore() resourceOption {
	return func(opts *resourceOptionSet) {
		opts.isCore = true
	}
}
