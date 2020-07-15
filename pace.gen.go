// Code generated by oto; DO NOT EDIT.

package pace

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Client is used to access Pace services.
type Client struct {
	// RemoteHost is the URL of the remote server that this Client should
	// access. Default: https://pace.dev
	RemoteHost string
	// HTTPClient is the http.Client to use when making HTTP requests.
	HTTPClient *http.Client
	// Debug writes a line of debug log output.
	Debug func(s string)
	// apiKey is the API key to use when interacting with the server.
	apiKey string
}

// New makes a new Client.
func New(apiKey string) *Client {
	c := &Client{
		RemoteHost: "https://pace.dev",
		apiKey:     apiKey,
		Debug:      func(s string) {},
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	return c
}

// CardsService is used to work with cards.
type CardsService struct {
	client *Client
}

// NewCardsService makes a new client for accessing CardsService services.
func NewCardsService(client *Client) *CardsService {
	return &CardsService{
		client: client,
	}
}

// CreateCard creates a new Card.
func (s *CardsService) CreateCard(ctx context.Context, r CreateCardRequest) (*CreateCardResponse, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.CreateCard: marshal CreateCardRequest")
	}
	url := s.client.RemoteHost + "/api/CardsService.CreateCard"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.CreateCard: NewRequest")
	}
	req.Header.Set("X-API-KEY", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.CreateCard")
	}
	defer resp.Body.Close()
	var response CreateCardResponse
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "CardsService.CreateCard: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.CreateCard: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("CardsService.CreateCard: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	return &response, nil
}

// GetCard gets a card.
func (s *CardsService) GetCard(ctx context.Context, r GetCardRequest) (*GetCardResponse, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.GetCard: marshal GetCardRequest")
	}
	url := s.client.RemoteHost + "/api/CardsService.GetCard"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.GetCard: NewRequest")
	}
	req.Header.Set("X-API-KEY", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.GetCard")
	}
	defer resp.Body.Close()
	var response GetCardResponse
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "CardsService.GetCard: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "CardsService.GetCard: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("CardsService.GetCard: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	return &response, nil
}

type CommentsService struct {
	client *Client
}

// NewCommentsService makes a new client for accessing CommentsService services.
func NewCommentsService(client *Client) *CommentsService {
	return &CommentsService{
		client: client,
	}
}

// AddComment adds a comment.
func (s *CommentsService) AddComment(ctx context.Context, r AddCommentRequest) (*AddCommentResponse, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "CommentsService.AddComment: marshal AddCommentRequest")
	}
	url := s.client.RemoteHost + "/api/CommentsService.AddComment"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "CommentsService.AddComment: NewRequest")
	}
	req.Header.Set("X-API-KEY", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "CommentsService.AddComment")
	}
	defer resp.Body.Close()
	var response AddCommentResponse
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "CommentsService.AddComment: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "CommentsService.AddComment: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("CommentsService.AddComment: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	return &response, nil
}

type AddCommentRequest struct {
	OrgID string `json:"OrgID"`

	TargetKind string `json:"TargetKind"`

	TargetID string `json:"TargetID"`

	Body string `json:"Body"`
}

// Person is a human who uses Pace.
type Person struct {
	ID string `json:"ID"`

	Username string `json:"Username"`

	Name string `json:"Name"`

	PhotoURL string `json:"PhotoURL"`
}

// File represents an attached file.
type File struct {

	// ID is the identifier for this file.
	ID string `json:"ID"`

	// CTime is the time the file was uploaded.
	CTime string `json:"CTime"`

	// Name is the name of the file.
	Name string `json:"Name"`

	// Path is the path of the file.
	Path string `json:"Path"`

	// ContentType is the type of the file.
	ContentType string `json:"ContentType"`

	// FileType is the type of file. Can be &#34;file&#34;, &#34;video&#34;, &#34;image&#34;, &#34;audio&#34; or
	// &#34;screenshare&#34;.
	FileType string `json:"FileType"`

	// Size is the size of the file in bytes.
	Size int `json:"Size"`

	// DownloadURL URL which can be used to get the file.
	DownloadURL string `json:"DownloadURL"`

	// ThumbnailURL is an optional thumbnail URL for this file.
	ThumbnailURL string `json:"ThumbnailURL"`

	// Author is the person who uploaded the file.
	Author Person `json:"Author"`
}

type Comment struct {
	ID string `json:"ID"`

	CTime string `json:"CTime"`

	MTime string `json:"MTime"`

	Body string `json:"Body"`

	BodyHTML string `json:"BodyHTML"`

	Author Person `json:"Author"`

	Files []File `json:"Files"`
}

type AddCommentResponse struct {
	Comment Comment `json:"Comment"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"Error,omitempty"`
}

type RelatedCardsSummary struct {
	Total int `json:"Total"`

	Done int `json:"Done"`

	Progress int `json:"Progress"`
}

type Card struct {

	// ID is the unique ID of the card within the org.
	ID string `json:"ID"`

	CTime string `json:"CTime"`

	MTime string `json:"MTime"`

	Order float64 `json:"Order"`

	TeamID string `json:"TeamID"`

	Slug string `json:"Slug"`

	Title string `json:"Title"`

	Status string `json:"Status"`

	Author Person `json:"Author"`

	Body string `json:"Body"`

	BodyHTML string `json:"BodyHTML"`

	Tags []string `json:"Tags"`

	// TakenByCurrentUser indicates whether the current user has taken this card or
	// not.
	TakenByCurrentUser bool `json:"TakenByCurrentUser"`

	// TakenByPeople is a list of people who have taken responsibility for this Card.
	TakenByPeople []Person `json:"TakenByPeople"`

	// Files are the list of files that are attached to this Card.
	Files []File `json:"Files"`

	RelatedCardsSummary RelatedCardsSummary `json:"RelatedCardsSummary"`
}

type CreateCardRequest struct {

	// OrgID is the org ID in which to create the card.
	OrgID string `json:"OrgID"`

	// TeamID is the team ID in which to create the card.
	TeamID string `json:"TeamID"`

	// Title is the title of the card.
	Title string `json:"Title"`

	// ParentTargetKind is the kind of target to relate this card to (e.g. card or
	// message)
	ParentTargetKind string `json:"ParentTargetKind"`

	// ParentTargetID is the ID of the item to relate this new card to.
	ParentTargetID string `json:"ParentTargetID"`
}

type CreateCardResponse struct {

	// Card is the card that was just created.
	Card Card `json:"Card"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"Error,omitempty"`
}

type GetCardRequest struct {
	OrgID string `json:"OrgID"`

	CardID string `json:"CardID"`
}

type GetCardResponse struct {
	Card Card `json:"Card"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"Error,omitempty"`
}
