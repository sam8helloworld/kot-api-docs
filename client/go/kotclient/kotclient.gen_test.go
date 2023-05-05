package kotclient_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/deepmap/oapi-codegen/pkg/types"
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

func TestGetCompany(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	got, err := sut.GetCompanyWithResponse(ctx)
	if err != nil {
		t.Fatalf("failed to sut.GetCompany: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetCompany{
		BusinessYearStartDate: "04/01",
		Code:                  "cmp",
		Host:                  "s2.kingtime.jp",
		Key:                   "1qazxsw23edc....",
		Name:                  "企業名",
		Settings: kotclient.GetCompanySettings{
			DecimalTreatType:  kotclient.RoundUp,
			TimeDisplayFormat: kotclient.Decimal,
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetAdministrators(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	queries := kotclient.GetAdministratorsParams{
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsAdministrator{"associatedEmployees"}),
	}
	got, err := sut.GetAdministratorsWithResponse(ctx, kotclient.Ptr(queries))
	if err != nil {
		t.Fatalf("failed to sut.GetAdministratorsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.GetAdministratorsItem{
		{
			Code: "admin",
			Key:  "4888a97dd7bbc6a26a18743f4697a1de4b38b6ee646a9620b286499c3df6918c",
			Name: "全権管理者",
		},
		{
			Code: "2000",
			Key:  "6713cd21ca98689efec9fe27352152c4c77a34b32f5de30b6335d141ad714baf",
			Name: "エリアマネージャー",
			AssociatedEmployees: kotclient.Ptr([]kotclient.GetAdministratorsAssociatedEmployee{
				{
					Code:      "1000",
					Key:       "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
					LastName:  "勤怠",
					FirstName: "太郎",
				},
				{
					Code:      "2000",
					Key:       "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe27352152c4",
					LastName:  "勤怠",
					FirstName: "花子",
				},
			}),
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetEmployees(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.GetEmployeesParams{
		Date:             kotclient.Ptr(kotclient.Date{Time: time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)}),
		Division:         kotclient.Ptr("1000"),
		IncludeResigner:  kotclient.Ptr(true),
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsEmployee{"emailAddresses"}),
	}
	got, err := sut.GetEmployeesWithResponse(ctx, kotclient.Ptr(queries))
	if err != nil {
		t.Fatalf("failed to sut.GetEmployeesWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.GetEmployeesItem{
		{
			DivisionCode:   "1000",
			DivisionName:   "本社",
			Gender:         kotclient.GetEmployeesItemGenderMale,
			TypeCode:       "1",
			TypeName:       "正社員",
			Code:           "1000",
			LastName:       "勤怠",
			FirstName:      "太郎",
			Key:            "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
			EmailAddresses: kotclient.Ptr([]types.Email{"kintaitarou@h-t.co.jp"}),
			EmployeeGroups: []kotclient.EmployeeGroup{
				{
					Code: "0001",
					Name: "人事部",
				},
				{
					Code: "0002",
					Name: "総務部",
				},
			},
		},
		{
			DivisionCode:   "1000",
			DivisionName:   "本社",
			Gender:         kotclient.GetEmployeesItemGenderFemale,
			TypeCode:       "1",
			TypeName:       "正社員",
			Code:           "2000",
			LastName:       "勤怠",
			FirstName:      "花子",
			Key:            "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe27352152c4",
			EmailAddresses: kotclient.Ptr([]types.Email{"kintaihanako@h-t.co.jp"}),
			EmployeeGroups: []kotclient.EmployeeGroup{
				{
					Code: "0003",
					Name: "営業部",
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}
