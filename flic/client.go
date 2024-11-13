package flic

import (
	"encoding/binary"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	OnButton func(ButtonEvent)
}

func NewClient(wsURL string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("dial error: %w", err)
	}

	return &Client{
		conn: conn,
	}, nil
}

// Connect to a specific Flic button
func (c *Client) Connect(bdAddr string) error {
	return c.writeCommand(CmdCreateConnectionChannel, 1, bdAddr)
}

func (c *Client) Listen() error {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}

		opcode := message[0]
		pos := 1

		switch opcode {
		case EvtButtonSingleOrDoubleClickOrHold:

			connID := int32(binary.LittleEndian.Uint32(message[pos:]))
			pos += 4

			clickType := message[pos]
			pos++

			wasQueued := message[pos] != 0
			pos++

			timeDiff := int32(binary.LittleEndian.Uint32(message[pos:]))

			event := ButtonEvent{
				ConnID:    connID,
				ClickType: parseClickType(clickType),
				WasQueued: wasQueued,
				TimeDiff:  timeDiff,
			}

			if c.OnButton != nil {
				c.OnButton(event)
			}
		}
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func parseClickType(ct uint8) ClickType {
	switch ct {
	case 0:
		return ButtonDown
	case 1:
		return ButtonUp
	case 2:
		return ButtonClick
	case 3:
		return ButtonSingleClick
	case 4:
		return ButtonDoubleClick
	case 5:
		return ButtonHold
	default:
		return ButtonClick
	}
}
