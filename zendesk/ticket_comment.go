package zendesk

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type TicketComment struct {
	ID          *int64       `json:"id,omitempty"`
	Type        *string      `json:"type,omitempty"`
	Body        *string      `json:"body,omitempty"`
	HTMLBody    *string      `json:"html_body,omitempty"`
	Public      *bool        `json:"public,omitempty"`
	AuthorID    *int64       `json:"author_id,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	Uploads     []string     `json:"uploads,omitempty"`
}

// ListTicketCommentsOptions specifies the optional parameters for the list ticket comments methods.
type ListTicketCommentsOptions struct {
	ListOptions

	SortOrder string `url:"sort_order"` // One of asc, desc. Defaults to asc
}

func (c *client) ListTicketComments(id int64, opts *ListTicketCommentsOptions) ([]TicketComment, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/tickets/%d/comments.json?%s", id, params.Encode()), out)
	return out.Comments, err
}
