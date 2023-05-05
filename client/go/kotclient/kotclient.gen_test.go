package kotclient_test

import (
	"context"
	"io"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/kot-api-docs/kotclient"
)

func TestRefreshAccessToken(t *testing.T) {
	// Example BearerToken
	// See: https://swagger.io/docs/specification/authentication/bearer-authentication/
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClient("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	input := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83"
	got, err := sut.RefreshAccessToken(ctx, input)
	if err != nil {
		t.Fatalf("failed to sut.RefreshAccessToken: %v", err)
	}
	defer got.Body.Close()

	if diff := cmp.Diff(201, got.StatusCode); diff != "" {
		t.Errorf("User value is mismatch (-want +got):\n%s", diff)
	}
	body, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("io.ReadAll: %v", err)
	}
	gotBody := string(body)
	want := `{"token":"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}`
	if diff := cmp.Diff(want, gotBody); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}
