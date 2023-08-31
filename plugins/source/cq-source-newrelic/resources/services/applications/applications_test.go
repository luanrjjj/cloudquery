package applications

import (
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/newrelic/cq-source-newrelic/client"
	"github.com/newrelic/cq-source-newrelic/client/mocks"
)

func listApplications(t *testing.T, ctrl *gomock.Controller) client.Services {
	fmt.Printf("M VAIASDLKASLDKASLDKLASKDLADKLAS: %v+\n", ctrl)

	m := mocks.NewMockApplicationService(ctrl)
	fmt.Printf("M VAIASDLKASLDKASLDKLASKDLADKLAS: %v+\n", m)
	services := client.Services{}
	fmt.Printf("services: %v+\n", services)
	var d apm.ListApplicationsParams
	fmt.Printf("d: %v+\n", d)
	err := faker.FakeObject(&d)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("d1: %v+\n", d)
	m.EXPECT().ListApplications(gomock.Any()).Return(d, nil, nil)
	fmt.Printf("d2: %v+\n", m.EXPECT().ListApplications(gomock.Any()).Return(d, nil, nil))

	return services

}

func TestApplicationList(t *testing.T) {
	client.MockTestHelper(t, Applications(), listApplications, client.TestOptions{})
}
