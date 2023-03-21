// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcnotify

import (
	"encoding/json"
)

type (
	// Optional link that can be added as a button to the email.
	Link struct {

		// Where the link should point to.
		//
		// Min length: 1
		// Max length: 1024
		Href string `json:"href"`

		// Text to display on link.
		//
		// Min length: 1
		// Max length: 40
		Text string `json:"text"`
	}

	// List of notification addresses.
	ListOfNotificationAdresses struct {
		Addresses []NotificationTypeAndAddress `json:"addresses"`

		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// Type of notification and its corresponding address.
	NotificationTypeAndAddress struct {
		NotificationAddress string `json:"notificationAddress"`

		// Possible values:
		//   * "email"
		//   * "pulse"
		//   * "matrix-room"
		//   * "slack-channel"
		NotificationType string `json:"notificationType"`
	}

	// Request to post a message on pulse.
	PostPulseMessageRequest struct {

		// Pulse message to send as plain text.
		//
		// Additional properties allowed
		Message json.RawMessage `json:"message"`

		// Routing-key to use when posting the message.
		//
		// Max length: 255
		RoutingKey string `json:"routingKey"`
	}

	// Request to send an email
	SendEmailRequest struct {

		// E-mail address to which the message should be sent
		Address string `json:"address"`

		// Content of the e-mail as **markdown**, will be rendered to HTML before
		// the email is sent. Notice that markdown allows for a few HTML tags, but
		// won't allow inclusion of script tags and other unpleasantries.
		//
		// Min length: 1
		// Max length: 102400
		Content string `json:"content"`

		// Optional link that can be added as a button to the email.
		Link Link `json:"link,omitempty"`

		// Reply-to e-mail (this property is optional)
		ReplyTo string `json:"replyTo,omitempty"`

		// Subject line of the e-mail, this is plain-text
		//
		// Min length: 1
		// Max length: 255
		Subject string `json:"subject"`

		// E-mail html template used to format your content.
		//
		// Possible values:
		//   * "simple"
		//   * "fullscreen"
		//
		// Default:    "simple"
		Template string `json:"template,omitempty" default:"simple"`
	}

	// Request to send a Matrix notice. Many of these fields are better understood by
	// checking the matrix spec itself. The precise definitions of these fields is
	// beyond the scope of this document.
	SendMatrixNoticeRequest struct {

		// Unformatted text that will be displayed in the room if you do not
		// specify `formattedBody` or if a user's client can not render the format.
		Body string `json:"body"`

		// The format for `formattedBody`. For instance, `org.matrix.custom.html`
		Format string `json:"format,omitempty"`

		// Text that will be rendered by matrix clients that support the given
		// format in that format. For instance, `<h1>Header Text</h1>`.
		FormattedBody string `json:"formattedBody,omitempty"`

		// Which of the `m.room.message` msgtypes to use. At the moment only the
		// types that take `body`/`format`/`formattedBody` are supported.
		//
		// Possible values:
		//   * "m.notice"
		//   * "m.text"
		//   * "m.emote"
		//
		// Default:    "m.notice"
		Msgtype string `json:"msgtype,omitempty" default:"m.notice"`

		// The fully qualified room name, such as `!whDRjjSmICCgrhFHsQ:mozilla.org`
		// If you are using riot, you can find this under the advanced settings for a room.
		RoomID string `json:"roomId"`
	}

	// Request to send a message to a Slack channel. The most interesting field in
	// this request is the `blocks` field which allows you to specify advanced
	// display layout for messages. This is best understood via the Slack API
	// documentation.
	SendSlackMessage struct {

		// An array of Slack attachments. See https://api.slack.com/messaging/composing/layouts#attachments.
		Attachments []interface{} `json:"attachments,omitempty"`

		// An array of Slack layout blocks. See https://api.slack.com/reference/block-kit/blocks.
		Blocks []interface{} `json:"blocks,omitempty"`

		// The unique Slack channel ID, such as `C123456GZ`.
		// In the app, this is the last section of the 'copy link' URL for a channel.
		ChannelID string `json:"channelId"`

		// The main message text. If no blocks are included, this is used as the
		// message text, otherwise this is used as alternative text and the blocks
		// are used.
		Text string `json:"text"`
	}
)
