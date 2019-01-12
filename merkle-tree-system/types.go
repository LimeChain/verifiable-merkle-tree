package main

type RootSaver interface {
	GetContractAddress() string
	GetSaverWalletAddress() string
	TriggerSave() (string, error)
}
