// Package msg provides message handling routines.
package msg

import spec "trpc.group/trpc-go/trpc-a2a-go/protocol"

// Text gets the first TextPart of a a2a protocol Message.
func Text(message spec.Message) string {
	for _, part := range message.Parts {
		if textPart, ok := part.(*spec.TextPart); ok {
			return textPart.Text
		}
	}
	return ""
}
