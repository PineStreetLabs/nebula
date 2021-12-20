package networks

// Params includes the scope and configuration of each network.
// In the cosmos-sdk, this is akin to the sdk.Config type.
// We choose not to use sdk.Config because of its use of init and how it is used globally in the cosmos-sdk library.
// Params tightly scopes relevant app chain configurations to work neatly across this library.
type Params struct {
	accountHRP          string
	validatorHRP        string
	consensusHRP        string
	VerifyAddressFormat func(b []byte) error
}

func (p Params) AccountHRP() string {
	return p.accountHRP
}

func (p Params) ValidatorHRP() string {
	return p.validatorHRP
}

func (p Params) ConsensusHRP() string {
	return p.consensusHRP
}
