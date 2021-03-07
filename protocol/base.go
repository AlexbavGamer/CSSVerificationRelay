package protocol

import (
	"github.com/AlexbavGamer/CSSVerificationRelay/packet"
	"github.com/bwmarrin/discordgo"
)

type Deliverable interface {
	Type() MessageType
	Marshal() []byte
	Content() string
	Author() string
	Plain() string
	Embed() *discordgo.MessageEmbed
}

type BaseMessage struct {
	Type MessageType

	EntityName string

	// Internal relay purposes only
	SenderID string
}

func (b BaseMessage) Author() string {
	return b.SenderID
}

func ParseBaseMessage(r *packet.PacketReader) (BaseMessage, error) {
	m := BaseMessage{}

	r.SetPos(0)

	m.Type = ParseMessageType(r.ReadUint8())

	var ok bool

	m.EntityName, ok = r.TryReadString()

	if !ok {
		return m, ErrCannotReadString
	}

	return m, nil
}
