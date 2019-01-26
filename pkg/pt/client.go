package pt

import (
	"fmt"
	"net/http"

	"github.com/salsita/go-pivotaltracker/v5/pivotal"
)

type Client struct {
	RequestDoer
	ClientCaller
}

const (
	ProjectOwner   string = "owner"
	ProjectMemeber string = "member"
	ProjectViewer  string = "viewer"
)

type ProjectMembership pivotal.ProjectMembership
type Project pivotal.Project
type ProjectRequest struct {
	Name                         string            `json:"name,omitempty"`
	Status                       string            `json:"status,omitempty"`
	IterationLength              int               `json:"iteration_length,omitempty"`
	WeekStartDay                 pivotal.Day       `json:"week_start_day,omitempty"`
	PointScale                   string            `json:"point_scale,omitempty"`
	BugsAndChoresAreEstimatable  bool              `json:"bugs_and_chores_are_estimatable,omitempty"`
	AutomaticPlanning            bool              `json:"automatic_planning,omitempty"`
	EnableTasks                  bool              `json:"enable_tasks,omitempty"`
	StartDate                    *pivotal.Date     `json:"start_date,omitempty"`
	TimeZone                     *pivotal.TimeZone `json:"time_zone,omitempty"`
	VelocityAveragedOver         int               `json:"velocity_averaged_over,omitempty"`
	NumberOfDoneIterationsToShow int               `json:"number_of_done_iterations_to_show,omitempty"`
	Description                  string            `json:"description,omitempty"`
	ProfileContent               string            `json:"profile_content,omitempty"`
	EnableIncomingEmails         bool              `json:"enable_incoming_emails,omitempty"`
	InitialVelocity              int               `json:"initial_velocity,omitempty"`
	ProjectType                  string            `json:"project_type,omitempty"`
	Public                       bool              `json:"public,omitempty"`
	AtomEnabled                  bool              `json:"atom_enabled,omitempty"`
	AccountID                    int               `json:"account_id,omitempty"`
	JoinAs                       string            `json:"join_as,omitempty"`
}

type Person struct {
	Kind     string `json:"kind,omitempty"`
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Initials string `json:"initials,omitempty"`
	Username string `json:"username,omitempty"`
}

type AccountMember struct {
	AccountMemberRequest
	Person Person `json:"person,omitempty"`
}

type AccountMemberRequest struct {
	Name           string `json:"name,omitempty"`
	AccountID      int    `json:"account_id,omitempty"`
	PersonID       int    `json:"person_id,omitempty"`
	Email          string `json:"email,omitempty"`
	Initials       string `json:"initials,omitempty"`
	Admin          bool   `json:"admin,omitempty"`
	ProjectCreator bool   `json:"project_creator,omitempty"`
}

type ProjectsRequest struct {
	NoOwner        bool   `json:"no_owner,omitempty"`
	NewAccountName string `json:"new_account_name,omitempty"`
	ProjectRequest
}

//go:generate counterfeiter . RequestDoer
type RequestDoer interface {
	Do(req *http.Request, v interface{}) (*http.Response, error)
	NewRequest(method, urlPath string, body interface{}) (*http.Request, error)
}

//go:generate counterfeiter . ClientCaller
type ClientCaller interface {
	ProjectCaller
	AccountMemberCaller
}

//go:generate counterfeiter . AccountMemberCaller
type AccountMemberCaller interface {
	ListAccountMembers(accountID int) ([]AccountMember, *http.Response, error)
	GetAccountMember(accountID int, accountMemberID int) (*AccountMember, *http.Response, error)
	NewAccountMember(accountID int, member AccountMemberRequest) (*AccountMember, *http.Response, error)
	UpdateAccountMember(accountID int, accountMemberID int, project AccountMemberRequest) (*AccountMember, *http.Response, error)
	DeleteAccountMember(accountID int, accountMemberID int) (*http.Response, error)
}

//go:generate counterfeiter . ProjectCaller
type ProjectCaller interface {
	ListProjects() ([]*Project, *http.Response, error)
	GetProject(projectID int) (*Project, *http.Response, error)
	NewProject(project ProjectsRequest) (*Project, *http.Response, error)
	UpdateProject(projectID int, project ProjectRequest) (*Project, *http.Response, error)
	DeleteProject(projectID int) (*http.Response, error)
}

func NewClient(apiToken string) ClientCaller {
	return &Client{
		RequestDoer: pivotal.NewClient(apiToken),
	}
}

// ListProjects returns all active projects for the current user.
func (service *Client) ListProjects() ([]*Project, *http.Response, error) {
	req, err := service.NewRequest("GET", "projects", nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := service.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, err
}

// GetProject returns a specific project's information.
func (service *Client) GetProject(projectID int) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/%v", projectID)
	req, err := service.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var project Project
	resp, err := service.Do(req, &project)
	if err != nil {
		return nil, resp, err
	}

	return &project, resp, err
}

// NewProject returns the created project's information.
func (service *Client) NewProject(project ProjectsRequest) (*Project, *http.Response, error) {
	req, err := service.NewRequest("POST", "projects", project)
	if err != nil {
		return nil, nil, err
	}
	responseProject := &Project{}
	resp, err := service.Do(req, responseProject)
	if err != nil {
		return nil, resp, err
	}

	return responseProject, resp, err
}

// UpdateProject returns the updated project's information.
func (service *Client) UpdateProject(projectID int, project ProjectRequest) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/%v", projectID)
	req, err := service.NewRequest("PUT", u, project)
	if err != nil {
		return nil, nil, err
	}

	responseProject := &Project{}
	resp, err := service.Do(req, responseProject)
	if err != nil {
		return nil, resp, err
	}

	return responseProject, resp, err
}

// DeleteProject deletes a given project by id.
func (service *Client) DeleteProject(projectID int) (*http.Response, error) {
	u := fmt.Sprintf("projects/%v", projectID)
	req, err := service.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := service.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// ListAccountMembers - list all account members
func (service *Client) ListAccountMembers(accountID int) ([]AccountMember, *http.Response, error) {
	req, err := service.NewRequest("GET", fmt.Sprintf("accounts/%v/memberships", accountID), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating request: %v", err)
	}

	responseMembers := make([]AccountMember, 0)
	resp, err := service.Do(req, &responseMembers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed calling service: %v", err)
	}

	return responseMembers, resp, nil
}

// GetAccountMember - retrieve an account member's details from the api
func (service *Client) GetAccountMember(accountID int, accountMemberID int) (*AccountMember, *http.Response, error) {
	req, err := service.NewRequest("GET", fmt.Sprintf("accounts/%v/memberships/%v", accountID, accountMemberID), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating request: %v", err)
	}

	responseMember := &AccountMember{}
	resp, err := service.Do(req, responseMember)
	if err != nil {
		return nil, resp, fmt.Errorf("failed calling service: %v", err)
	}

	return responseMember, resp, nil
}

// NewAccountMember - creates a new account member record
func (service *Client) NewAccountMember(accountID int, member AccountMemberRequest) (*AccountMember, *http.Response, error) {
	req, err := service.NewRequest("POST", fmt.Sprintf("accounts/%v/memberships", accountID), member)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating request: %v", err)
	}

	responseMember := &AccountMember{}
	resp, err := service.Do(req, responseMember)
	if err != nil {
		return nil, resp, fmt.Errorf("failed calling service: %v", err)
	}

	return responseMember, resp, nil
}

// UpdateAccountMember - updates a given account member by id.
func (service *Client) UpdateAccountMember(accountID int, accountMemberID int, member AccountMemberRequest) (*AccountMember, *http.Response, error) {
	req, err := service.NewRequest("PUT", fmt.Sprintf("accounts/%v/memberships/%v", accountID, accountMemberID), member)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating request: %v", err)
	}

	responseMember := &AccountMember{}
	resp, err := service.Do(req, responseMember)
	if err != nil {
		return nil, resp, fmt.Errorf("failed calling service: %v", err)
	}

	return responseMember, resp, nil
}

// DeleteAccountMember deletes a given member by id.
func (service *Client) DeleteAccountMember(accountID int, accountMemberID int) (*http.Response, error) {
	req, err := service.NewRequest("DELETE", fmt.Sprintf("accounts/%v/memberships/%v", accountID, accountMemberID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request: %v", err)
	}

	resp, err := service.Do(req, nil)
	if err != nil {
		return resp, fmt.Errorf("failed calling service: %v", err)
	}

	return resp, nil
}
