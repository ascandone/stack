/*
Formance Stack API

Testing LogsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package formance

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	client "github.com/formancehq/formance-sdk-go"
)

func Test_formance_LogsApiService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test LogsApiService ListLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string

		resp, httpRes, err := apiClient.LogsApi.ListLogs(context.Background(), ledger).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}