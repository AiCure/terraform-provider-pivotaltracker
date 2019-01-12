package pt_integration_test

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
	uuid "github.com/satori/go.uuid"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
)

func TestClientIntegrations(t *testing.T) {
	RegisterTestingT(t)
	token := os.Getenv("PVTL_TRACKER_TOKEN")
	if token == "" {
		t.Skip("This test requires PVTL_TRACKER_TOKEN to be set")
	}
	t.Run("Client should be able to create/list/update/delete a project", func(t *testing.T) {
		u, err := uuid.NewV4()
		if err != nil {
			t.Fatal(err)
		}

		projectName := "prj-" + u.String()
		modifiedProjectName := "mod-" + u.String()
		client := pt.NewClient(token)

		t.Log("Calling ListProjects endpoint")
		projectList, _, err := client.ListProjects()
		Expect(err).NotTo(HaveOccurred(), "call to list projects should not fail")
		found := false
		for _, p := range projectList {
			if p.Name == projectName {
				found = true
			}
		}
		Expect(found).To(BeFalse(), "the project should not yet exist")

		t.Log("Creating a new project")
		projectsRequest := pt.ProjectsRequest{}
		projectsRequest.Name = projectName
		projectsRequest.WeekStartDay = pivotal.DayTuesday
		projectsRequest.ProjectType = pivotal.ProjectTypePrivate
		projectsRequest.JoinAs = pt.ProjectViewer
		project, _, err := client.NewProject(projectsRequest)
		Expect(err).NotTo(HaveOccurred(), "call to create project should not fail")
		Expect(project.ID).NotTo(BeZero(), "project response should have an ID")
		Expect(project.Name).To(Equal(projectName), "project should have the proper name")
		projectList, _, err = client.ListProjects()
		Expect(err).NotTo(HaveOccurred(), "call to list projects should not fail")
		found = false
		for _, p := range projectList {
			if p.Name == projectName {
				found = true
			}
		}

		Expect(found).To(BeTrue(), "the project should have been created")

		t.Log("Update the project")
		modifiedProject, _, err := client.UpdateProject(project.ID, pt.ProjectRequest{Name: modifiedProjectName})
		Expect(err).NotTo(HaveOccurred(), "call to update project should not fail")
		Expect(modifiedProject.Name).NotTo(Equal(projectName), "project should not have old name")
		Expect(modifiedProject.Name).To(Equal(modifiedProjectName), "project should have the new name")

		t.Log("Fetch the project")
		modifiedProject, _, err = client.GetProject(project.ID)
		Expect(err).NotTo(HaveOccurred(), "call to get project should not fail")
		Expect(modifiedProject.Name).NotTo(Equal(projectName), "project should not have old name")
		Expect(modifiedProject.Name).To(Equal(modifiedProjectName), "project should have the new name")

		t.Log("Delete the project")
		_, err = client.DeleteProject(project.ID)
		Expect(err).NotTo(HaveOccurred(), "call to delete project should not fail")

		projectList, _, err = client.ListProjects()
		Expect(err).NotTo(HaveOccurred(), "call to list projects should not fail")
		found = false
		for _, p := range projectList {
			if p.Name == modifiedProjectName || p.Name == projectName {
				found = true
			}
		}
		Expect(found).To(BeFalse(), "the project should no longer exist")
	})
}
