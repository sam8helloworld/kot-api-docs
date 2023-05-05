package kotclient_test

import (
	"context"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/kot-api-docs/kotclient"
)

func TestRefreshAccessToken(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	input := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83"
	got, err := sut.RefreshAccessTokenWithResponse(ctx, input)
	if err != nil {
		t.Fatalf("failed to sut.RefreshAccessToken: %v", err)
	}

	if diff := cmp.Diff(201, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.RefreshAccessToken{
		Token: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON201); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetAccessTokenAvailability(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	input := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83"
	got, err := sut.GetAccessTokenAvailabilityWithResponse(ctx, input)
	if err != nil {
		t.Fatalf("failed to sut.GetAccessTokenAvailabilityWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetAccessTokenAvailability{
		Available: true,
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}
