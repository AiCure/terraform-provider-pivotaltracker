package projects_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt/ptfakes"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider/resources/projects"
)

func TestProject(t *testing.T) {
	RegisterTestingT(t)
	projectResource := projects.NewProjectResource()
	t.Run("Schema", func(t *testing.T) {
		t.Run("Should provide all fields supported by the API", func(t *testing.T) {
			for k, v := range projectResource.Schema {
				Expect([]string{
					"no_owner",
					"new_account_name",
					"name",
					"status",
					"iteration_length",
					"week_start_day",
					"point_scale",
					"bugs_and_chores_are_estimatable",
					"automatic_planning",
					"enable_tasks",
					"start_date",
					"time_zone",
					"velocity_averaged_over",
					"number_of_done_iterations_to_show",
					"description",
					"profile_content",
					"enable_incoming_emails",
					"initial_velocity",
					"project_type",
					"public",
					"atom_enabled",
					"account_id",
					"join_as",
				}).To(ContainElement(k), "schema element is not expected")
				Expect(v).NotTo(BeNil(), "resource value is not valid")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when create fails", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).To(HaveOccurred(), "it should error")
		})

		t.Run("when it creates a new project", func(t *testing.T) {
			err := projectResource.Create(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(nil).NotTo(BeNil(), "it should call the tracker api")
			Expect(nil).NotTo(BeNil(), "it should set the schema to the value from tracker")
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
