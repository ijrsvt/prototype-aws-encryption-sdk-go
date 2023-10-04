package keyproviders

type MasterKeyProviderConfig interface{}

type MasterKeyProvider interface{}

type MasterKeyConfig interface {
	ProviderId() string
}

type MasterKey interface{}
