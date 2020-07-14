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
			log.Printf("Chain Commit: %x %s", commit.GetChainIDHash(), commit.GetEntityState().String())
		case *eventmessages.FactomEvent_EntryCommit:
			commit := ev.Event.(*eventmessages.FactomEvent_EntryCommit).EntryCommit
			log.Printf("Entry Commit: %x %s", commit.GetEntryHash(), commit.GetEntityState().String())
		case *eventmessages.FactomEvent_EntryReveal:
			reveal := ev.Event.(*eventmessages.FactomEvent_EntryReveal).EntryReveal
			log.Printf("Entry Reveal: %x (chain %x) %s", reveal.GetEntry().GetHash(), reveal.GetEntry().GetChainID(), reveal.GetEntityState().String())
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
