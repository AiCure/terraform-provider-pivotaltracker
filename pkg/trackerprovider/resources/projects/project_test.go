package projects_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	. "github.com/onsi/gomega"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt/ptfakes"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider/resources/projects"
)

func TestProject(t *testing.T) {
	RegisterTestingT(t)
	t.Run("Schema", func(t *testing.T) {
		t.Run("Should provide all fields supported by the API", func(t *testing.T) {
			schemaMap, _, projectResource, _ := createControlDataset()
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
		_, controlProjectRequest, projectResource, fakeData := createControlDataset()
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

			Expect(fakeData.Id()).To(Equal(strconv.Itoa(controlResourceID)),
				"it should set the id of the newly created resource",
			)

			t.Run("it should call the client with the same values set in the resource data", func(t *testing.T) {
				Expect(projectRequest.AccountID).To(Equal(controlProjectRequest.AccountID))
				Expect(projectRequest.AtomEnabled).To(Equal(controlProjectRequest.AtomEnabled))
				Expect(projectRequest.AutomaticPlanning).To(Equal(controlProjectRequest.AutomaticPlanning))
				Expect(projectRequest.BugsAndChoresAreEstimatable).To(Equal(controlProjectRequest.BugsAndChoresAreEstimatable))
				Expect(projectRequest.Description).To(Equal(controlProjectRequest.Description))
				Expect(projectRequest.EnableIncomingEmails).To(Equal(controlProjectRequest.EnableIncomingEmails))
				Expect(projectRequest.EnableTasks).To(Equal(controlProjectRequest.EnableTasks))
				Expect(projectRequest.InitialVelocity).To(Equal(controlProjectRequest.InitialVelocity))
				Expect(projectRequest.IterationLength).To(Equal(controlProjectRequest.IterationLength))
				Expect(projectRequest.JoinAs).To(Equal(controlProjectRequest.JoinAs))
				Expect(projectRequest.Name).To(Equal(controlProjectRequest.Name))
				Expect(projectRequest.NumberOfDoneIterationsToShow).To(Equal(controlProjectRequest.NumberOfDoneIterationsToShow))
				Expect(projectRequest.PointScale).To(Equal(controlProjectRequest.PointScale))
				Expect(projectRequest.ProfileContent).To(Equal(controlProjectRequest.ProfileContent))
				Expect(projectRequest.ProjectType).To(Equal(controlProjectRequest.ProjectType))
				Expect(projectRequest.Public).To(Equal(controlProjectRequest.Public))
				Expect(projectRequest.Status).To(Equal(controlProjectRequest.Status))
				Expect(projectRequest.VelocityAveragedOver).To(Equal(controlProjectRequest.VelocityAveragedOver))
			})
		})
	})

	t.Run("Delete", func(t *testing.T) {
		_, _, projectResource, fakeData := createControlDataset()
		t.Run("when delete fails", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeClient.DeleteProjectReturns(nil, fmt.Errorf("some erroor msg"))
			err := projectResource.Delete(fakeData, fakeClient)
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it deletes an existing project", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeData.SetId("1234")
			err := projectResource.Delete(fakeData, fakeClient)
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(fakeClient.DeleteProjectCallCount()).To(Equal(1), "it should call delete exactly once")
			Expect(fakeClient.DeleteProjectArgsForCall(0)).To(Equal(1234), "it should call delete on the project ID in the tracker api")
		})
	})

	t.Run("Exists", func(t *testing.T) {
		_, _, projectResource, fakeData := createControlDataset()
		fakeData.SetId("1234")
		t.Run("when exists call fails", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeClient.GetProjectReturns(&pt.Project{}, nil, fmt.Errorf("some erroor msg"))
			_, err := projectResource.Exists(fakeData, fakeClient)
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when project doesnt exist", func(t *testing.T) {
			fakeCaller := &ptfakes.FakeClientCaller{}
			fakeCaller.GetProjectReturns(&pt.Project{}, nil, nil)
			exists, err := projectResource.Exists(fakeData, fakeCaller)
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeFalse(), "it should return false")
		})

		t.Run("when project exists", func(t *testing.T) {
			fakeCaller := &ptfakes.FakeClientCaller{}
			fakeCaller.GetProjectReturns(&pt.Project{ID: 1234}, nil, nil)
			exists, err := projectResource.Exists(fakeData, fakeCaller)
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeTrue(), "it should return true")
		})
	})

	t.Run("Read", func(t *testing.T) {
		_, _, projectResource, fakeData := createControlDataset()
		fakeData.SetId("1234")
		t.Run("when read fails", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeClient.GetProjectReturns(&pt.Project{}, nil, fmt.Errorf("some erroor msg"))
			err := projectResource.Read(fakeData, fakeClient)
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it reads an existing project", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			controlProjectResponse := &pt.Project{ID: 1234, AccountID: 12345, AtomEnabled: true, Description: "blah"}
			fakeClient.GetProjectReturns(controlProjectResponse, nil, nil)
			err := projectResource.Read(fakeData, fakeClient)
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(fakeClient.GetProjectCallCount()).To(Equal(1), "it should call the tracker api")
			Expect(fakeData.Id()).To(Equal(strconv.Itoa(1234)),
				"it should set the id of the resource",
			)

			t.Run("it set the resource data with the values from the tracker API", func(t *testing.T) {
				Expect(fakeData.Get("account_id")).To(Equal(controlProjectResponse.AccountID), "account_id")
				Expect(fakeData.Get("atom_enabled")).To(Equal(controlProjectResponse.AtomEnabled), "atom_enabled")
				Expect(fakeData.Get("automatic_planning")).To(Equal(controlProjectResponse.AutomaticPlanning), "automatic_planning")
				Expect(fakeData.Get("bugs_and_chores_are_estimatable")).To(Equal(controlProjectResponse.BugsAndChoresAreEstimatable), "bugs_and_chores_are_estimateable")
				Expect(fakeData.Get("description")).To(Equal(controlProjectResponse.Description), "description")
				Expect(fakeData.Get("enable_incoming_emails")).To(Equal(controlProjectResponse.EnableIncomingEmails), "enable_incoming_emails")
				Expect(fakeData.Get("enable_tasks")).To(Equal(controlProjectResponse.EnableTasks), "enable_tasks")
				Expect(fakeData.Get("initial_velocity")).To(Equal(controlProjectResponse.InitialVelocity), "initial_velocity")
				Expect(fakeData.Get("iteration_length")).To(Equal(controlProjectResponse.IterationLength), "iteration_length")
				Expect(fakeData.Get("name")).To(Equal(controlProjectResponse.Name), "name")
				Expect(fakeData.Get("number_of_done_iterations_to_show")).To(Equal(controlProjectResponse.NumberOfDoneIterationsToShow), "number_of_done_iterations_to_show")
				Expect(fakeData.Get("point_scale")).To(Equal(controlProjectResponse.PointScale), "point_scale")
				Expect(fakeData.Get("profile_content")).To(Equal(controlProjectResponse.ProfileContent), "profile_content")
				Expect(fakeData.Get("project_type")).To(Equal(controlProjectResponse.ProjectType), "project_type")
				Expect(fakeData.Get("public")).To(Equal(controlProjectResponse.Public), "public")
				Expect(fakeData.Get("velocity_averaged_over")).To(Equal(controlProjectResponse.VelocityAveragedOver), "velocity_averaged_over")
			})
		})
	})

	t.Run("Update", func(t *testing.T) {
		_, _, projectResource, fakeData := createControlDataset()
		fakeData.SetId("1234")
		t.Run("when read fails", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeClient.UpdateProjectReturns(&pt.Project{}, nil, fmt.Errorf("some erroor msg"))
			err := projectResource.Update(fakeData, fakeClient)
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it reads an existing project", func(t *testing.T) {
			fakeClient := &ptfakes.FakeClientCaller{}
			fakeData.Set("name", "someprojects")
			fakeData.Set("account_id", 12345)
			fakeData.Set("atom_enabled", true)
			fakeData.Set("description", "blah")
			fakeClient.UpdateProjectReturns(&pt.Project{ID: 1234}, nil, nil)
			err := projectResource.Update(fakeData, fakeClient)
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(fakeClient.UpdateProjectCallCount()).To(Equal(1), "it should call the tracker api")
			_, updatedProject := fakeClient.UpdateProjectArgsForCall(0)
			Expect(fakeData.Id()).To(Equal(strconv.Itoa(1234)),
				"it should set the id of the resource",
			)
			t.Run("it set the resource data with the values from the tracker API", func(t *testing.T) {
				Expect(fakeData.Get("account_id")).To(Equal(updatedProject.AccountID), "account_id")
				Expect(fakeData.Get("atom_enabled")).To(Equal(updatedProject.AtomEnabled), "atom_enabled")
				Expect(fakeData.Get("automatic_planning")).To(Equal(updatedProject.AutomaticPlanning), "automatic_planning")
				Expect(fakeData.Get("bugs_and_chores_are_estimatable")).To(Equal(updatedProject.BugsAndChoresAreEstimatable), "bugs_and_chores_are_estimateable")
				Expect(fakeData.Get("description")).To(Equal(updatedProject.Description), "description")
				Expect(fakeData.Get("enable_incoming_emails")).To(Equal(updatedProject.EnableIncomingEmails), "enable_incoming_emails")
				Expect(fakeData.Get("enable_tasks")).To(Equal(updatedProject.EnableTasks), "enable_tasks")
				Expect(fakeData.Get("initial_velocity")).To(Equal(updatedProject.InitialVelocity), "initial_velocity")
				Expect(fakeData.Get("iteration_length")).To(Equal(updatedProject.IterationLength), "iteration_length")
				Expect(fakeData.Get("name")).To(Equal(updatedProject.Name), "name")
				Expect(fakeData.Get("number_of_done_iterations_to_show")).To(Equal(updatedProject.NumberOfDoneIterationsToShow), "number_of_done_iterations_to_show")
				Expect(fakeData.Get("point_scale")).To(Equal(updatedProject.PointScale), "point_scale")
				Expect(fakeData.Get("profile_content")).To(Equal(updatedProject.ProfileContent), "profile_content")
				Expect(fakeData.Get("project_type")).To(Equal(updatedProject.ProjectType), "project_type")
				Expect(fakeData.Get("public")).To(Equal(updatedProject.Public), "public")
				Expect(fakeData.Get("velocity_averaged_over")).To(Equal(updatedProject.VelocityAveragedOver), "velocity_averaged_over")
			})
		})
	})
}

func createControlDataset() (map[string]interface{}, pt.ProjectsRequest, *schema.Resource, *schema.ResourceData) {

	projectResource := projects.NewProjectResource()
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

	fakeData := projectResource.TestResourceData()
	for k, v := range schemaMap {
		fakeData.Set(k, v)
	}
	return schemaMap, controlProjects, projectResource, fakeData
}
