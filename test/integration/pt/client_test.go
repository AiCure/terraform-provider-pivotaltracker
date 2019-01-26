package pt_integration_test

import (
	"os"
	"strconv"
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

	t.Run("Client should be able to create/list/update/delete account member", func(t *testing.T) {
		accountIDVal := os.Getenv("PVTL_TRACKER_ACCOUNT_ID")
		if accountIDVal == "" {
			t.Skip("This test requires PVTL_TRACKER_ACCOUNT_ID to be set")
		}

		accountID, _ := strconv.Atoi(accountIDVal)
		u := uuid.NewV4()
		memberEmail := u.String() + "@devnull.io"
		memberInitials := "JC"
		memberName := "mbr-" + u.String()
		client := pt.NewClient(token)

		t.Log("Calling ListAccountMembers endpoint")
		memberList, _, err := client.ListAccountMembers(accountID)
		Expect(err).NotTo(HaveOccurred(),
			"call to list account members should not fail",
		)
		found := false
		for _, m := range memberList {
			if m.Name == memberName {
				found = true
			}
		}
		Expect(found).To(BeFalse(),
			"the account member should not yet exist",
		)

		t.Log("Creating a new account member")
		memberRequest := pt.AccountMemberRequest{}
		memberRequest.Name = memberName
		memberRequest.Email = memberEmail
		memberRequest.Initials = memberInitials
		member, _, err := client.NewAccountMember(accountID, memberRequest)
		Expect(err).NotTo(HaveOccurred(),
			"call to create account member should not fail",
		)
		Expect(member.Person.ID).NotTo(BeZero(),
			"account member response should have an ID",
		)

		Expect(member.Person.Name).To(Equal(memberName),
			"account member should have the proper name",
		)
		memberList, _, err = client.ListAccountMembers(accountID)
		Expect(err).NotTo(HaveOccurred(),
			"call to list members should not fail",
		)
		found = false
		for _, p := range memberList {
			if p.Person.Name == memberName {
				found = true
			}
		}

		Expect(found).To(BeTrue(),
			"the member should have been created",
		)

		t.Log("Update the member")
		Expect(member.ProjectCreator).To(BeFalse(),
			"insure old value is set properly",
		)
		modifiedMember, _, err := client.UpdateAccountMember(accountID, member.Person.ID, pt.AccountMemberRequest{ProjectCreator: true})
		Expect(err).NotTo(HaveOccurred(),
			"call to update member should not fail",
		)
		Expect(modifiedMember.ProjectCreator).To(BeTrue(),
			"member should not have old attribute",
		)

		t.Log("Fetch the member")
		modifiedMember, _, err = client.GetAccountMember(accountID, member.Person.ID)
		Expect(err).NotTo(HaveOccurred(),
			"call to get member should not fail",
		)
		Expect(modifiedMember.Person.Name).To(Equal(memberName),
			"member should have name",
		)

		t.Log("Delete the member")
		_, err = client.DeleteAccountMember(accountID, member.Person.ID)
		Expect(err).NotTo(HaveOccurred(),
			"call to delete account member should not fail",
		)

		memberList, _, err = client.ListAccountMembers(accountID)
		Expect(err).NotTo(HaveOccurred(),
			"call to list account members should not fail",
		)
		found = false
		for _, p := range memberList {
			if p.Person.Name == memberName {
				found = true
			}
		}
		Expect(found).To(BeFalse(),
			"the account member should no longer exist",
		)
	})

	t.Run("Client should be able to create/list/update/delete a project", func(t *testing.T) {
		u := uuid.NewV4()
		projectName := "prj-" + u.String()
		modifiedProjectName := "mod-" + u.String()
		client := pt.NewClient(token)

		t.Log("Calling ListProjects endpoint")
		projectList, _, err := client.ListProjects()
		Expect(err).NotTo(HaveOccurred(),
			"call to list projects should not fail",
		)
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
		Expect(err).NotTo(HaveOccurred(),
			"call to create project should not fail",
		)
		Expect(project.ID).NotTo(BeZero(),
			"project response should have an ID",
		)
		Expect(project.Name).To(Equal(projectName),
			"project should have the proper name",
		)
		projectList, _, err = client.ListProjects()
		Expect(err).NotTo(HaveOccurred(),
			"call to list projects should not fail",
		)
		found = false
		for _, p := range projectList {
			if p.Name == projectName {
				found = true
			}
		}

		Expect(found).To(BeTrue(),
			"the project should have been created",
		)

		t.Log("Update the project")
		modifiedProject, _, err := client.UpdateProject(project.ID, pt.ProjectRequest{Name: modifiedProjectName})
		Expect(err).NotTo(HaveOccurred(),
			"call to update project should not fail",
		)
		Expect(modifiedProject.Name).NotTo(Equal(projectName),
			"project should not have old name",
		)
		Expect(modifiedProject.Name).To(Equal(modifiedProjectName),
			"project should have the new name",
		)

		t.Log("Fetch the project")
		modifiedProject, _, err = client.GetProject(project.ID)
		Expect(err).NotTo(HaveOccurred(),
			"call to get project should not fail",
		)
		Expect(modifiedProject.Name).NotTo(Equal(projectName),
			"project should not have old name",
		)
		Expect(modifiedProject.Name).To(Equal(modifiedProjectName),
			"project should have the new name",
		)

		t.Log("Delete the project")
		_, err = client.DeleteProject(project.ID)
		Expect(err).NotTo(HaveOccurred(),
			"call to delete project should not fail",
		)

		projectList, _, err = client.ListProjects()
		Expect(err).NotTo(HaveOccurred(),
			"call to list projects should not fail",
		)
		found = false
		for _, p := range projectList {
			if p.Name == modifiedProjectName || p.Name == projectName {
				found = true
			}
		}
		Expect(found).To(BeFalse(),
			"the project should no longer exist",
		)
	})
}
