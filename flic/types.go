package flic

const (
	CmdCreateConnectionChannel = 3
	CmdRemoveConnectionChannel = 4
)

const (
	EvtCreateConnectionChannelResponse = 1
	EvtConnectionStatusChanged         = 2
	EvtConnectionChannelRemoved        = 3
	EvtButtonUpOrDown                  = 4
	EvtButtonClickOrHold               = 5
	EvtButtonSingleOrDoubleClick       = 6
	EvtButtonSingleOrDoubleClickOrHold = 7
)

type ClickType string

const (
	ButtonDown        ClickType = "ButtonDown"
	ButtonUp          ClickType = "ButtonUp"
	ButtonClick       ClickType = "ButtonClick"
	ButtonSingleClick ClickType = "ButtonSingleClick"
	ButtonDoubleClick ClickType = "ButtonDoubleClick"
	ButtonHold        ClickType = "ButtonHold"
)

type ConnectionStatus string

const (
	StatusDisconnected ConnectionStatus = "Disconnected"
	StatusConnected    ConnectionStatus = "Connected"
	StatusReady        ConnectionStatus = "Ready"
)

type ButtonEvent struct {
	ConnID    int32
	ClickType ClickType
	WasQueued bool
	TimeDiff  int32
}
