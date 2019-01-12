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
type Person pivotal.Person
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
