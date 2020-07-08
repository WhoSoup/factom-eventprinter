package main

import (
	"log"

	eater "github.com/WhoSoup/factom-eater"
	"github.com/WhoSoup/factom-eater/eventmessages"
)

func main() {
	eat, err := eater.Launch(":8040")
	if err != nil {
		panic(err)
	}

	for ev := range eat.Reader() {
		switch ev.Event.(type) {
		case *eventmessages.FactomEvent_ChainCommit:
			commit := ev.Event.(*eventmessages.FactomEvent_ChainCommit).ChainCommit
			log.Printf("Chain Commit: %x", commit.GetChainIDHash())
		case *eventmessages.FactomEvent_EntryCommit:
			commit := ev.Event.(*eventmessages.FactomEvent_EntryCommit).EntryCommit
			log.Printf("Entry Commit: %x", commit.GetEntryHash())
		case *eventmessages.FactomEvent_EntryReveal:
			reveal := ev.Event.(*eventmessages.FactomEvent_EntryReveal).EntryReveal
			log.Printf("Entry Reveal: %x (chain %x)", reveal.GetEntry().GetHash(), reveal.GetEntry().GetChainID())
		case *eventmessages.FactomEvent_StateChange:
			// not sure what this is
			change := ev.Event.(*eventmessages.FactomEvent_StateChange).StateChange
			log.Printf("State Change: %d %s", change.GetBlockHeight(), change.GetEntityState().String())
		case *eventmessages.FactomEvent_DirectoryBlockCommit:
			dbc := ev.Event.(*eventmessages.FactomEvent_DirectoryBlockCommit).DirectoryBlockCommit
			log.Printf("DBlock Commit: %x", dbc.GetDirectoryBlock().GetHash())
		case *eventmessages.FactomEvent_ProcessListEvent:
			ple := ev.Event.(*eventmessages.FactomEvent_ProcessListEvent).ProcessListEvent.GetProcessListEvent()
			switch ple.(type) {
			case *eventmessages.ProcessListEvent_NewMinuteEvent:
				minute := ple.(*eventmessages.ProcessListEvent_NewMinuteEvent)
				log.Printf("New Minute: %d/%d", minute.NewMinuteEvent.BlockHeight, minute.NewMinuteEvent.NewMinute)
			case *eventmessages.ProcessListEvent_NewBlockEvent:
				block := ple.(*eventmessages.ProcessListEvent_NewBlockEvent)
				log.Printf("New Height: %d", block.NewBlockEvent.NewBlockHeight)
			}

		case *eventmessages.FactomEvent_NodeMessage:
			nm := ev.Event.(*eventmessages.FactomEvent_NodeMessage).NodeMessage
			log.Printf("Node Message: %s", nm.GetMessageText())
		case *eventmessages.FactomEvent_DirectoryBlockAnchor:
			anchor := ev.Event.(*eventmessages.FactomEvent_DirectoryBlockAnchor).DirectoryBlockAnchor
			log.Printf("Anchor: %d (BTC: %v) (Eth: %v)", anchor.GetBlockHeight(), anchor.GetBtcConfirmed(), anchor.GetEthereumConfirmed())
		default:
			log.Println("unknown message type:", ev.Event)
		}
	}
}

/*

type FactomEvent_ChainCommit struct {
	ChainCommit *ChainCommit `protobuf:"bytes,4,opt,name=chainCommit,proto3,oneof" json:"chainCommit,omitempty"`
}
type FactomEvent_EntryCommit struct {
	EntryCommit *EntryCommit `protobuf:"bytes,5,opt,name=entryCommit,proto3,oneof" json:"entryCommit,omitempty"`
}
type FactomEvent_EntryReveal struct {
	EntryReveal *EntryReveal `protobuf:"bytes,6,opt,name=entryReveal,proto3,oneof" json:"entryReveal,omitempty"`
}
type FactomEvent_StateChange struct {
	StateChange *StateChange `protobuf:"bytes,7,opt,name=stateChange,proto3,oneof" json:"stateChange,omitempty"`
}
type FactomEvent_DirectoryBlockCommit struct {
	DirectoryBlockCommit *DirectoryBlockCommit `protobuf:"bytes,8,opt,name=directoryBlockCommit,proto3,oneof" json:"directoryBlockCommit,omitempty"`
}
type FactomEvent_ProcessListEvent struct {
	ProcessListEvent *ProcessListEvent `protobuf:"bytes,9,opt,name=processListEvent,proto3,oneof" json:"processListEvent,omitempty"`
}
type FactomEvent_NodeMessage struct {
	NodeMessage *NodeMessage `protobuf:"bytes,10,opt,name=nodeMessage,proto3,oneof" json:"nodeMessage,omitempty"`
}
type FactomEvent_DirectoryBlockAnchor struct {
	DirectoryBlockAnchor *DirectoryBlockAnchor `protobuf:"bytes,11,opt,name=directoryBlockAnchor,proto3,oneof" json:"directoryBlockAnchor,omitempty"`
}


*/
