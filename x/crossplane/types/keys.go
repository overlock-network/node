package types

const (
	// ModuleName defines the module name
	ModuleName = "overlock"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_overlock"

	ConfigurationKey = "Configuration/value/"

	ConfigurationCountKey = "Configuration/count/"

	ConfigurationCreatedEvent = "configuration-created"

	ConfigurationDeletedEvent = "configuration-deleted"

	ConfigurationIndex = "id"

	CompositionKey = "Composition/value/"

	CompositionCountKey = "Composition/count/"

	CompositionCreatedEvent = "composition-created"

	CompositionDeletedEvent = "composition-deleted"

	CompositionIndex = "id"

	XRDKey = "XRD/value/"

	XRDCountKey = "XRD/count/"

	XRDCreatedEvent = "CompositeResource-created"

	XRDDeletedEvent = "CompositeResource-deleted"

	XRDIndex = "id"

	EnvironmentKey = "Environment/value/"

	EnvironmentCountKey = "Environment/count/"

	EnvironmentCreatedEvent = "environment-created"

	EnvironmentDeletedEvent = "environment-deleted"

	EnvironmentIndex = "id"

	ProviderKey = "Provider/value/"

	ProviderCountKey = "Provider/count/"

	ProviderCreatedEvent = "provider-created"

	ProviderDeletedEvent = "provider-deleted"

	ProviderIndex = "id"

	FunctionKey = "Function/value/"

	FunctionCountKey = "Function/count/"

	FunctionCreatedEvent = "function-created"

	FunctionDeletedEvent = "function-deleted"

	FunctionIndex = "id"
)

var (
	ParamsKey = []byte("p_overlock")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
