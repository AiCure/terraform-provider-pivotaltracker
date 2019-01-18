package projects_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt/ptfakes"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider/resources/projects"
)

func TestProject(t *testing.T) {
	RegisterTestingT(t)
	t.Run("Create", func(t *testing.T) {
		t.Skip("this is not yet implemented")
	})

	t.Run("Read", func(t *testing.T) {
		t.Skip("this is not yet implemented")
	})

	t.Run("Delete", func(t *testing.T) {
		t.Skip("this is not yet implemented")
	})

	t.Run("Exists", func(t *testing.T) {
		t.Skip("this is not yet implemented")
		t.Run("when project doesnt exist", func(t *testing.T) {
			exists, err := projects.Exists(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeFalse(), "it should return false")
		})

		t.Run("when project exists", func(t *testing.T) {
			exists, err := projects.Exists(nil, &ptfakes.FakeClientCaller{})
			Expect(err).NotTo(HaveOccurred(), "it should not error")
			Expect(exists).To(BeTrue(), "it should return true")
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Skip("this is not yet implemented")
	})
}
