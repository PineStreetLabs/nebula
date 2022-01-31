package messages

import (
	"github.com/PineStreetLabs/nebula/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// GovDeposit returns a MsgDeposit message.
func GovDeposit(proposalID uint64, depositor account.Address, coins sdk.Coins) govtypes.MsgDeposit {
	return govtypes.MsgDeposit{
		ProposalId: proposalID,
		Depositor:  depositor.String(),
		Amount:     coins,
	}
}

// GovVote returns a MsgVote message.
func GovVote(proposalID uint64, voter account.Address, vote govtypes.VoteOption) govtypes.MsgVote {
	return govtypes.MsgVote{
		ProposalId: proposalID,
		Voter:      voter.String(),
		Option:     vote,
	}
}

// GovSubmitProposal returns a MsgSubmitProposal message.
func GovSubmitProposal(content govtypes.Content, proposer sdk.Address, deposit sdk.Coins) govtypes.MsgSubmitProposal {
	proposal := govtypes.MsgSubmitProposal{
		InitialDeposit: deposit,
		Proposer:       proposer.String(),
	}
	proposal.SetContent(content)

	return proposal
}
