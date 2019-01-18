package trackerprovider_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/trackerprovider"
)

func TestResourceProvider(t *testing.T) {
	RegisterTestingT(t)
	t.Run("should support expected resources", func(t *testing.T) {
		provider := trackerprovider.Create(nil)
		Expect(provider.ResourcesMap).NotTo(BeEmpty(), "there should be some resources")
		for k, v := range provider.ResourcesMap {
			Expect([]string{
				"pivotaltracker_project",
			}).To(ContainElement(k), "resource type is not expected")
			Expect(v).NotTo(BeNil(), "resource value is not valid")
		}
	})
}
