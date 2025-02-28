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

	ConfigurationUpdatedEvent = "configuration-updated"

	ConfigurationDeletedEvent = "configuration-deleted"

	ConfigurationIndex = "id"

	CompositionKey = "Composition/value/"

	CompositionCountKey = "Composition/count/"

	CompositionCreatedEvent = "composition-created"

	CompositionUpdatedEvent = "composition-updated"

	CompositionDeletedEvent = "composition-deleted"

	CompositionIndex = "id"

	XRDKey = "XRD/value/"

	XRDCountKey = "XRD/count/"

	XRDCreatedEvent = "CompositeResource-created"

	XRDUpdateEvent = "CompositeResource-updated"

	XRDDeletedEvent = "CompositeResource-deleted"

	XRDIndex = "id"

	EnvironmentKey = "Environment/value/"

	EnvironmentCountKey = "Environment/count/"

	EnvironmentCreatedEvent = "environment-created"

	EnvironmentUpdatedEvent = "environment-updated"

	EnvironmentDeletedEvent = "environment-deleted"

	EnvironmentIndex = "id"

	ProviderKey = "Provider/value/"

	ProviderCountKey = "Provider/count/"

	ProviderCreatedEvent = "provider-created"

	ProviderUpdatedEvent = "provider-updated"

	ProviderDeletedEvent = "provider-deleted"

	ProviderIndex = "id"

	FunctionKey = "Function/value/"

	FunctionCountKey = "Function/count/"

	FunctionCreatedEvent = "function-created"

	FunctionUpdatedEvent = "function-updated"

	FunctionDeletedEvent = "function-deleted"

	FunctionIndex = "id"
)

var (
	ParamsKey = []byte("p_overlock")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
