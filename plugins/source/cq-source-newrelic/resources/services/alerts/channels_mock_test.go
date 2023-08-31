package alerts

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/newrelic/cq-source-newrelic/client"
	"testing"
)

func buildDashboardListsMock(t *testing.T, ctrl *gomock.Controller) client.DatadogServices {
	m := mocks.NewMockChannelsAPIClient(ctrl)
	services := client.Services.Alert.ListChannels()

	var d datadogV1.DashboardListListResponse
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListDashboardLists(gomock.Any()).Return(d, nil, nil)

	return services
}

func TestDashboardLists(t *testing.T) {
	client.DatadogMockTestHelper(t, DashboardLists(), buildDashboardListsMock, client.TestOptions{})
}
