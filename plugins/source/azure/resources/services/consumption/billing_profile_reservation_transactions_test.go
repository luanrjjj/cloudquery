package consumption

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createBillingProfileReservationTransactions(router *mux.Router) error {
	var item armconsumption.ReservationTransactionsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.NextLink = to.Ptr("")

	router.HandleFunc("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{profileId}/providers/Microsoft.Consumption/reservationTransactions", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestBillingProfileReservationTransactions(t *testing.T) {
	client.MockTestHelper(t, BillingProfileReservationTransactions(), createBillingProfileReservationTransactions)
}
