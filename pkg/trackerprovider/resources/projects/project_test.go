package projects_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt/ptfakes"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider/resources/projects"
)

func TestProject(t *testing.T) {
	RegisterTestingT(t)
	projectResource := projects.NewProjectResource()
	t.Run("Schema", func(t *testing.T) {
		t.Run("Should provide all fields supported by the API", func(t *testing.T) {
			schemaMap, _ := createControlDataset()
			fieldset := []string{}
			for f, _ := range schemaMap {
				fieldset = append(fieldset, f)
			}

			for k, v := range projectResource.Schema {
				Expect(fieldset).To(ContainElement(k), "schema element is not expected")
				Expect(v).NotTo(BeNil(), "schema value is not valid")
				Expect(v.Description).NotTo(BeEmpty(), "we shouldnt add elements without having a description")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		schemaMap, controlProjectRequest := createControlDataset()
		fakeData := projectResource.TestResourceData()
		for k, v := range schemaMap {
			fakeData.Set(k, v)
		}

		t.Run("when create fails", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeClient.NewProjectReturns(&pt.Project{}, nil, fmt.Errorf("some erroor msg"))
			err := projectResource.Create(fakeData, fakeClient)
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it creates a new project", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			controlResourceID := 1234
			fakeClient.NewProjectReturns(&pt.Project{
				ID: controlResourceID,
			}, nil, nil)
			err := projectResource.Create(fakeData, fakeClient)
			projectRequest := fakeClient.NewProjectArgsForCall(0)
			Expect(err).NotTo(HaveOccurred(),
				"it should not error",
			)

			Expect(fakeData.Id()).To(Equal(string(controlResourceID)),
				fmt.Sprint("it should set the id of the newly created resource"),
			)

			t.Run("it should call the client with the same values set in the resource data", func(t *testing.T) {
				Expect(projectRequest.Name).To(Equal(controlProjectRequest.Name))
				Expect(projectRequest.Status).To(Equal(controlProjectRequest.Status))
				Expect(projectRequest.IterationLength).To(Equal(controlProjectRequest.IterationLength))
				Expect(projectRequest.PointScale).To(Equal(controlProjectRequest.PointScale))
				Expect(projectRequest.BugsAndChoresAreEstimatable).To(Equal(controlProjectRequest.BugsAndChoresAreEstimatable))
				Expect(projectRequest.AutomaticPlanning).To(Equal(controlProjectRequest.AutomaticPlanning))
				Expect(projectRequest.EnableTasks).To(Equal(controlProjectRequest.EnableTasks))
				Expect(projectRequest.VelocityAveragedOver).To(Equal(controlProjectRequest.VelocityAveragedOver))
				Expect(projectRequest.NumberOfDoneIterationsToShow).To(Equal(controlProjectRequest.NumberOfDoneIterationsToShow))
				Expect(projectRequest.Description).To(Equal(controlProjectRequest.Description))
				Expect(projectRequest.ProfileContent).To(Equal(controlProjectRequest.ProfileContent))
				Expect(projectRequest.EnableIncomingEmails).To(Equal(controlProjectRequest.EnableIncomingEmails))
				Expect(projectRequest.InitialVelocity).To(Equal(controlProjectRequest.InitialVelocity))
				Expect(projectRequest.ProjectType).To(Equal(controlProjectRequest.ProjectType))
				Expect(projectRequest.Public).To(Equal(controlProjectRequest.Public))
				Expect(projectRequest.AtomEnabled).To(Equal(controlProjectRequest.AtomEnabled))
				Expect(projectRequest.AccountID).To(Equal(controlProjectRequest.AccountID))
				Expect(projectRequest.JoinAs).To(Equal(controlProjectRequest.JoinAs))
			})
		})
	})

	t.Run("Read", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when read fails", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it reads an existing project", func(t *testing.T) {
			err := projectResource.Read(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(nil).NotTo(BeNil(), "it should call the tracker api")
			Expect(nil).NotTo(BeNil(), "it should set the schema to the value from tracker")
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when delete fails", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it deletes an existing project", func(t *testing.T) {
			err := projectResource.Delete(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(nil).NotTo(BeNil(), "it should call delete on the project in the tracker api")
		})
	})

	t.Run("Exists", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when exists fails", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when project doesnt exist", func(t *testing.T) {
			exists, err := projectResource.Exists(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeFalse(), "it should return false")
		})

		t.Run("when project exists", func(t *testing.T) {
			exists, err := projectResource.Exists(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeTrue(), "it should return true")
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when update fails", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).To(HaveOccurred(), "it should error")
		})
	})
}

func createControlDataset() (map[string]interface{}, pt.ProjectsRequest) {

	controlProjects := pt.ProjectsRequest{
		NewAccountName: "testing",
		NoOwner:        false,
		ProjectRequest: pt.ProjectRequest{
			AccountID:                   1234,
			AtomEnabled:                 false,
			AutomaticPlanning:           false,
			BugsAndChoresAreEstimatable: false,
			Description:                 "testing",
			EnableIncomingEmails:        false,
			EnableTasks:                 false,
			InitialVelocity:             10,
			IterationLength:             1,
			JoinAs:                      "testing",
			Name:                        "testing",
			NumberOfDoneIterationsToShow: 1,
			PointScale:                   "testing",
			ProfileContent:               "testing",
			ProjectType:                  "testing",
			Public:                       false,
			StartDate:                    &pivotal.Date{},
			Status:                       "testing",
			TimeZone:                     &pivotal.TimeZone{},
			VelocityAveragedOver:         2,
			WeekStartDay:                 "testing",
		},
	}
	schemaMap := map[string]interface{}{
		"account_id":                        controlProjects.AccountID,
		"atom_enabled":                      controlProjects.AtomEnabled,
		"automatic_planning":                controlProjects.AutomaticPlanning,
		"bugs_and_chores_are_estimatable":   controlProjects.BugsAndChoresAreEstimatable,
		"description":                       controlProjects.Description,
		"enable_incoming_emails":            controlProjects.EnableIncomingEmails,
		"enable_tasks":                      controlProjects.EnableTasks,
		"initial_velocity":                  controlProjects.InitialVelocity,
		"iteration_length":                  controlProjects.IterationLength,
		"join_as":                           controlProjects.JoinAs,
		"name":                              controlProjects.Name,
		"new_account_name":                  controlProjects.NewAccountName,
		"no_owner":                          controlProjects.NoOwner,
		"number_of_done_iterations_to_show": controlProjects.NumberOfDoneIterationsToShow,
		"point_scale":                       controlProjects.PointScale,
		"profile_content":                   controlProjects.ProfileContent,
		"project_type":                      controlProjects.ProjectType,
		"public":                            controlProjects.Public,
		"start_date":                        controlProjects.StartDate,
		"status":                            controlProjects.Status,
		"time_zone":                         controlProjects.TimeZone,
		"velocity_averaged_over":            controlProjects.VelocityAveragedOver,
		"week_start_day":                    controlProjects.WeekStartDay,
	}
	return schemaMap, controlProjects
}
