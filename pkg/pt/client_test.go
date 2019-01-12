package pt_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt/ptfakes"
)

func TestClient(t *testing.T) {
	RegisterTestingT(t)
	t.Run("ProjectCaller", func(t *testing.T) {
		t.Run("API Call Structure", func(t *testing.T) {
			var client *pt.Client
			fakeRequestDoer := &ptfakes.FakeRequestDoer{}
			client = &pt.Client{RequestDoer: fakeRequestDoer}
			table := []struct {
				name          string
				path          string
				controlMethod string
				hasData       bool
				call          func()
			}{
				{"DeleteProject", "projects/1234", "DELETE", false, func() {
					client.DeleteProject(1234)
				}},
				{"UpdateProject", "projects/1234", "PUT", true, func() {
					client.UpdateProject(1234, pt.ProjectRequest{})
				}},
				{"NewProject", "projects", "POST", true, func() {
					client.NewProject(pt.ProjectsRequest{})
				}},
				{"ListProjects", "projects", "GET", false, func() {
					client.ListProjects()
				}},
				{"GetProject", "projects/1234", "GET", false, func() {
					client.GetProject(1234)
				}},
			}

			for i, record := range table {
				t.Run(record.name, func(t *testing.T) {
					record.call()
					method, path, data := fakeRequestDoer.NewRequestArgsForCall(i)
					Expect(data == nil).NotTo(
						Equal(record.hasData),
						fmt.Sprintf("when true we should have data when false we should not (%v: %v)",
							record.hasData,
							data,
						),
					)
					Expect(path).To(
						Equal(record.path),
						"path for api call is not correct",
					)
					Expect(method).To(
						Equal(record.controlMethod),
						"method for api call is not correct",
					)
				})
			}
		})

		t.Run("with errors", func(t *testing.T) {
			var client *pt.Client
			table := []struct {
				name string
				call func() error
			}{
				{"DeleteProject", func() error {
					_, err := client.DeleteProject(1234)
					return err
				}},
				{"UpdateProject", func() error {
					_, _, err := client.UpdateProject(1234, pt.ProjectRequest{})
					return err
				}},
				{"NewProject", func() error {
					_, _, err := client.NewProject(pt.ProjectsRequest{})
					return err
				}},
				{"ListProjects", func() error {
					_, _, err := client.ListProjects()
					return err
				}},
				{"GetProject", func() error {
					_, _, err := client.GetProject(1234)
					return err
				}},
			}

			for _, record := range table {
				t.Run(record.name, func(t *testing.T) {
					fakeRequestDoer := &ptfakes.FakeRequestDoer{}
					client = &pt.Client{RequestDoer: fakeRequestDoer}
					doError := fmt.Errorf("fake do error")
					requestError := fmt.Errorf("fake request error")

					fakeRequestDoer.NewRequestReturns(nil, nil)
					fakeRequestDoer.DoReturns(nil, doError)
					Expect(record.call()).To(
						Equal(doError),
						"error should be returned when Do fails in client",
					)

					fakeRequestDoer.NewRequestReturns(nil, requestError)
					fakeRequestDoer.DoReturns(nil, nil)
					Expect(record.call()).To(
						Equal(requestError),
						"error should be returned when NewRequest fails in client",
					)
				})
			}
		})
	})
}
