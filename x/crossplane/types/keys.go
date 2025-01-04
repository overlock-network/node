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

	ConfigurationIndex = "id"

	CompositionKey = "Composition/value/"

	CompositionCountKey = "Composition/count/"

	CompositionCreatedEvent = "composition-created"

	CompositionIndex = "id"

	XRDKey = "XRD/value/"

	XRDCountKey = "XRD/count/"

	XRDCreatedEvent = "CompositeResource-created"

	XRDIndex = "id"

	EnvironmentKey = "Environment/value/"

	EnvironmentCountKey = "Environment/count/"

	EnvironmentCreatedEvent = "environment-created"

	EnvironmentIndex = "id"

	ProviderKey = "Provider/value/"

	ProviderCountKey = "Provider/count/"

	ProviderCreatedEvent = "provider-created"

	ProviderIndex = "id"
)

var (
	ParamsKey = []byte("p_overlock")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
