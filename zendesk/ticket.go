package zendesk

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type Ticket struct {
	ID              *int64         `json:"id,omitempty"`
	URL             *string        `json:"url,omitempty"`
	ExternalID      *string        `json:"external_id,omitempty"`
	Type            *string        `json:"type,omitempty"`
	Subject         *string        `json:"subject,omitempty"`
	RawSubject      *string        `json:"raw_subject,omitempty"`
	Description     *string        `json:"description,omitempty"`
	Comment         *TicketComment `json:"comment,omitempty"`
	Priority        *string        `json:"priority,omitempty"`
	Status          *string        `json:"status,omitempty"`
	Recipient       *string        `json:"recipient,omitempty"`
	RequesterID     *int64         `json:"requester_id,omitempty"`
	SubmitterID     *int64         `json:"submitter_id,omitempty"`
	AssigneeID      *int64         `json:"assignee_id,omitempty"`
	OrganizationID  *int64         `json:"organization_id,omitempty"`
	GroupID         *int64         `json:"group_id,omitempty"`
	CollaboratorIDs []int64        `json:"collaborator_ids,omitempty"`
	ForumTopicID    *int64         `json:"forum_topic_id,omitempty"`
	ProblemID       *int64         `json:"problem_id,omitempty"`
	HasIncidents    *bool          `json:"has_incidents,omitempty"`
	DueAt           *time.Time     `json:"due_at,omitempty"`
	Tags            []string       `json:"tags,omitempty"`
	Via             *Via           `json:"via,omitempty"`
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`
	CustomFields    []CustomField  `json:"custom_fields,omitempty"`
}

type CustomField struct {
	ID    *int64      `json:"id"`
	Value interface{} `json:"value"`
}

func (c *client) ShowTicket(id int64) (*Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d.json", id), out)
	return out.Ticket, err
}

func (c *client) CreateTicket(ticket *Ticket) (*Ticket, error) {
	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.post("/api/v2/tickets.json", in, out)
	return out.Ticket, err
}

func (c *client) UpdateTicket(id int64, ticket *Ticket) (*Ticket, error) {
	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/tickets/%d.json", id), in, out)
	return out.Ticket, err
}

func (c *client) UpdateManyTickets(tickets []Ticket) ([]Ticket, error) {
	in := &APIPayload{Tickets: tickets}
	out := new(APIPayload)
	err := c.put("/api/v2/tickets/update_many.json", in, out)
	return out.Tickets, err
}

// ListTicketsOptions specifies the optional parameters for the list tickets methods.
type ListTicketsOptions struct {
	ListOptions

	Include   string `url:"include"`    // Possible values are comment_count
	SortBy    string `url:"sort_by"`    // Possible values are assignee, assignee.name, created_at, group, id, locale, requester, requester.name, status, subject, updated_at
	SortOrder string `url:"sort_order"` // One of asc, desc. Defaults to asc
}

func (c *client) ListUserTicketsRequested(id int64, opts *ListTicketsOptions) ([]Ticket, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/users/%d/tickets/requested.json?%s", id, params.Encode()), out)
	return out.Tickets, err
}

func (c *client) ListUserTicketsCCd(id int64, opts *ListTicketsOptions) ([]Ticket, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/users/%d/tickets/ccd.json?%s", id, params.Encode()), out)
	return out.Tickets, err
}

func (c *client) ListOrganizationTickets(id int64, opts *ListTicketsOptions) ([]Ticket, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/organizations/%d/tickets.json?%s", id, params.Encode()), out)
	return out.Tickets, err
}

// ListTicketIncidents list all incidents related to the problem
func (c *client) ListTicketIncidents(problemID int64) ([]Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/incidents.json", problemID), out)

	return out.Tickets, err
}
