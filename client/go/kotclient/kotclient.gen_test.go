package kotclient_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
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
			EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
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
			EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaihanako@h-t.co.jp"}),
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

func TestGetEmployee(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.GetEmployeeParams{
		Date:             kotclient.Ptr(kotclient.Date{Time: time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)}),
		IncludeResigner:  kotclient.Ptr(true),
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsEmployee{"emailAddresses"}),
	}
	employeeCode := "1000"
	got, err := sut.GetEmployeeWithResponse(ctx, employeeCode, kotclient.Ptr(queries))
	if err != nil {
		t.Fatalf("failed to sut.GetEmployeeWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetEmployee{
		DivisionCode:   "1000",
		DivisionName:   "本社",
		Gender:         kotclient.GetEmployeeGenderMale,
		TypeCode:       "1",
		TypeName:       "正社員",
		Code:           "1000",
		LastName:       "勤怠",
		FirstName:      "太郎",
		Key:            "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
		EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
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
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestRegisterEmployee(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	reqBody := kotclient.EmployeeRequest{
		DivisionCode:   "1000",
		Gender:         kotclient.EmployeeRequestGenderMale,
		TypeCode:       "1",
		Code:           "1000",
		LastName:       "勤怠",
		FirstName:      "太郎",
		EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
	}
	got, err := sut.RegisterEmployeeWithResponse(ctx, reqBody)
	if err != nil {
		t.Fatalf("failed to sut.RefreshAccessToken: %v", err)
	}

	if diff := cmp.Diff(201, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.RegisterEmployee{
		DivisionCode:   "1000",
		Gender:         kotclient.RegisterEmployeeGenderMale,
		TypeCode:       "1",
		Code:           "1000",
		LastName:       "勤怠",
		FirstName:      "太郎",
		Key:            "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe27352152c4",
		EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON201); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestUpdateEmployee(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	reqBody := kotclient.EmployeeRequest{
		DivisionCode:   "1000",
		Gender:         kotclient.EmployeeRequestGenderMale,
		TypeCode:       "1",
		Code:           "1000",
		LastName:       "勤怠",
		FirstName:      "太郎",
		EmailAddresses: kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
	}
	employeeKey := "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3"
	queries := kotclient.UpdateEmployeeParams{
		UpdateDate: kotclient.Ptr(kotclient.Date{Time: time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)}),
	}
	got, err := sut.UpdateEmployeeWithResponse(ctx, employeeKey, kotclient.Ptr(queries), reqBody)
	if err != nil {
		t.Fatalf("failed to sut.UpdateEmployeeWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.UpdateEmployee{
		DivisionCode:                "1000",
		DivisionName:                "本社",
		Gender:                      kotclient.UpdateEmployeeGenderMale,
		TypeCode:                    "1",
		TypeName:                    "正社員",
		Code:                        "1000",
		LastName:                    "勤怠",
		FirstName:                   "太郎",
		Key:                         "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe27352152c4",
		LastNamePhonetics:           kotclient.Ptr("キンタイ"),
		FirstNamePhonetics:          kotclient.Ptr("タロウ"),
		BirthDate:                   kotclient.Ptr(kotclient.Date{Time: time.Date(1990, 9, 1, 0, 0, 0, 0, time.UTC)}),
		HiredDate:                   kotclient.Ptr(kotclient.Date{Time: time.Date(2013, 4, 1, 0, 0, 0, 0, time.UTC)}),
		ResignationDate:             kotclient.Ptr(kotclient.Date{Time: time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC)}),
		EmailAddresses:              kotclient.Ptr([]openapi_types.Email{"kintaitarou@h-t.co.jp"}),
		AllDayRegardingWorkInMinute: kotclient.Ptr(480),
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestDeleteEmployee(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}
	employeeKey := "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3"
	got, err := sut.DeleteEmployeeWithResponse(ctx, employeeKey)
	if err != nil {
		t.Fatalf("failed to sut.DeleteEmployeeWithResponse: %v", err)
	}

	if diff := cmp.Diff(204, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDivisions(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	got, err := sut.GetDivisionsWithResponse(ctx)
	if err != nil {
		t.Fatalf("failed to sut.GetDivisionsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.DivisionResponse{
		{
			Code:          "1000",
			Name:          "本社",
			DayBorderTime: "00:00",
		},
		{
			Code:          "2000",
			Name:          "支社",
			DayBorderTime: "05:00",
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetWorkingTypes(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	got, err := sut.GetWorkingTypesWithResponse(ctx)
	if err != nil {
		t.Fatalf("failed to sut.GetWorkingTypesWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.WorkingTypeResponse{
		{
			Code: "1",
			Name: "正社員",
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetEmployeeGroups(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.Ptr(kotclient.GetEmployeeGroupsParams{
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsEmployeeGroups{"category"}),
	})
	got, err := sut.GetEmployeeGroupsWithResponse(ctx, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetEmployeeGroupsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.GetEmployeeGroupsItem{
		{
			Code: "0001",
			Name: "人事部",
			Category: kotclient.Ptr(kotclient.GetEmployeeGroupsItemCategory{
				Code: "KEIKAN",
				Name: "経営管理本部",
			}),
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDailyWorkings(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.Ptr(kotclient.GetDailyWorkingsParams{
		Division:         kotclient.Ptr("1000"),
		Ondivision:       kotclient.Ptr(true),
		Start:            kotclient.Ptr(kotclient.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)}),
		End:              kotclient.Ptr(kotclient.Date{Time: time.Date(2016, 5, 10, 0, 0, 0, 0, time.UTC)}),
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsEmployeeGroups{"currentDateEmployee"}),
	})
	got, err := sut.GetDailyWorkingsWithResponse(ctx, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetDailyWorkingsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetDailyWorkings{
		{
			Date: openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
			DailyWorkings: []kotclient.DailyWorkingResponse{
				{
					Assigned:     480,
					AutoBreakOff: 1,
					BreakTime:    60,
					CurrentDateEmployee: kotclient.Ptr(kotclient.DailyWorkingCurrentDateEmployee{
						Code:         "1000",
						DivisionCode: "1000",
						DivisionName: "本社",
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
						FirstName:          "太郎",
						FirstNamePhonetics: "タロウ",
						Gender:             "male",
						LastName:           "勤怠",
						LastNamePhonetics:  "キンタイ",
						TypeCode:           "1",
						TypeName:           "正社員",
					}),
					CustomDailyWorkings: []kotclient.DailyWorkingCustomDailyWorking{
						{
							CalculationResult:   1,
							CalculationUnitCode: 1,
							Code:                "dCus1",
							Name:                "日別カスタム1",
						},
						{
							CalculationResult:   10,
							CalculationUnitCode: 2,
							Code:                "dCus2",
							Name:                "日別カスタム2",
						},
						{
							CalculationResult:   100,
							CalculationUnitCode: 4,
							Code:                "dCus3",
							Name:                "日別カスタム3",
						},
					},
					Date:                  openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
					DiscretionaryVacation: 0,
					EarlyLeave:            0,
					EmployeeKey:           "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
					HolidaysObtained: kotclient.DailyWorkingHolidaysObtained{
						FulltimeHoliday: kotclient.DailyWorkingFulltimeHoliday{
							Code: 1,
							Name: "有休",
						},
						HalfdayHolidays: []kotclient.DailyWorkingHalfdayHoliday{
							{
								TypeName: "PM休",
								Code:     1,
								Name:     "有休",
							},
						},
						HourHolidays: []kotclient.DailyWorkingHourHoliday{
							{
								Start:   time.Date(2016, 5, 1, 10, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
								End:     time.Date(2016, 5, 1, 11, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
								Minutes: 60,
								Code:    1,
								Name:    "有休",
							},
						},
					},
					IsClosing:             true,
					IsError:               false,
					IsHelp:                false,
					Late:                  0,
					LateNight:             0,
					LateNightOvertime:     0,
					LateNightUnassigned:   0,
					Overtime:              135,
					TotalWork:             615,
					Unassigned:            135,
					WorkPlaceDivisionCode: "1000",
					WorkPlaceDivisionName: kotclient.Ptr("本社"),
					WorkdayTypeName:       "平日",
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDailyWorking(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	date := openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)}
	queries := kotclient.Ptr(kotclient.GetDailyWorkingParams{
		Division:         kotclient.Ptr("1000"),
		Ondivision:       kotclient.Ptr(true),
		AdditionalFields: kotclient.Ptr(kotclient.AdditionalFieldsEmployeeGroups{"currentDateEmployee"}),
	})
	got, err := sut.GetDailyWorkingWithResponse(ctx, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetDailyWorkingWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetDailyWorking{
		Date: openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
		DailyWorkings: []kotclient.DailyWorkingResponse{
			{
				Assigned:     480,
				AutoBreakOff: 1,
				BreakTime:    60,
				CurrentDateEmployee: kotclient.Ptr(kotclient.DailyWorkingCurrentDateEmployee{
					Code:         "1000",
					DivisionCode: "1000",
					DivisionName: "本社",
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
					FirstName:          "太郎",
					FirstNamePhonetics: "タロウ",
					Gender:             "male",
					LastName:           "勤怠",
					LastNamePhonetics:  "キンタイ",
					TypeCode:           "1",
					TypeName:           "正社員",
				}),
				CustomDailyWorkings: []kotclient.DailyWorkingCustomDailyWorking{
					{
						CalculationResult:   1,
						CalculationUnitCode: 1,
						Code:                "dCus1",
						Name:                "日別カスタム1",
					},
					{
						CalculationResult:   10,
						CalculationUnitCode: 2,
						Code:                "dCus2",
						Name:                "日別カスタム2",
					},
					{
						CalculationResult:   100,
						CalculationUnitCode: 4,
						Code:                "dCus3",
						Name:                "日別カスタム3",
					},
				},
				Date:                  openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
				DiscretionaryVacation: 0,
				EarlyLeave:            0,
				EmployeeKey:           "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
				HolidaysObtained: kotclient.DailyWorkingHolidaysObtained{
					FulltimeHoliday: kotclient.DailyWorkingFulltimeHoliday{
						Code: 1,
						Name: "有休",
					},
					HalfdayHolidays: []kotclient.DailyWorkingHalfdayHoliday{
						{
							TypeName: "PM休",
							Code:     1,
							Name:     "有休",
						},
					},
					HourHolidays: []kotclient.DailyWorkingHourHoliday{
						{
							Start:   time.Date(2016, 5, 1, 10, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							End:     time.Date(2016, 5, 1, 11, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							Minutes: 60,
							Code:    1,
							Name:    "有休",
						},
					},
				},
				IsClosing:             true,
				IsError:               false,
				IsHelp:                false,
				Late:                  0,
				LateNight:             0,
				LateNightOvertime:     0,
				LateNightUnassigned:   0,
				Overtime:              135,
				TotalWork:             615,
				Unassigned:            135,
				WorkPlaceDivisionCode: "1000",
				WorkPlaceDivisionName: kotclient.Ptr("本社"),
				WorkdayTypeName:       "平日",
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestRegisterDailyWorkingTimerecord(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	employeeKey := "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3"
	body := kotclient.DailyWorkingTimerecordRequest{
		Date:      openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
		Time:      time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC),
		Latitude:  kotclient.Ptr(35.6672237),
		Longitude: kotclient.Ptr(139.7422207),
	}
	got, err := sut.RegisterDailyWorkingTimerecordWithResponse(ctx, employeeKey, body)
	if err != nil {
		t.Fatalf("failed to sut.RegisterDailyWorkingTimerecordWithResponse: %v", err)
	}

	if diff := cmp.Diff(201, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.RegisterDailyWorkingTimerecord{
		Date:        openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
		EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
		TimeRecord: kotclient.DailyWorkingTimerecord{
			Time:         time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC),
			Code:         "1",
			Name:         "出勤",
			DivisionName: "正社員",
			DivisionCode: "1000",
			Latitude:     35.6672237,
			Longitude:    139.7422207,
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON201); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDailyWorkingTimerecords(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.Ptr(kotclient.GetDailyWorkingTimerecordsParams{
		EmployeeKeys: kotclient.Ptr([]string{"8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3"}),
		Division:     kotclient.Ptr("1000"),
		Ondivision:   kotclient.Ptr(true),
		Start:        kotclient.Ptr(openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)}),
		End:          kotclient.Ptr(openapi_types.Date{Time: time.Date(2016, 5, 10, 0, 0, 0, 0, time.UTC)}),
	})
	got, err := sut.GetDailyWorkingTimerecordsWithResponse(ctx, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetDailyWorkingTimerecordsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetDailyWorkingTimerecords{
		{
			Date: openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
			DailyWorkings: []kotclient.DailyWorkingTimerecordResponse{
				{
					Date:        openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
					EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
					CurrentDateEmployee: kotclient.Ptr(kotclient.DailyWorkingCurrentDateEmployee{
						Code:               "1000",
						DivisionCode:       "1000",
						DivisionName:       "本社",
						Gender:             kotclient.DailyWorkingCurrentDateEmployeeGenderMale,
						TypeCode:           "1",
						TypeName:           "正社員",
						LastName:           "勤怠",
						FirstName:          "太郎",
						LastNamePhonetics:  "キンタイ",
						FirstNamePhonetics: "タロウ",
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
					}),
					TimeRecord: []kotclient.DailyWorkingTimerecord{
						{
							Time:         time.Date(2016, 5, 1, 9, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							Code:         "1",
							Name:         "出勤",
							DivisionName: "本社",
							DivisionCode: "1000",
							Latitude:     35.6672237,
							Longitude:    139.7422207,
						},
						{
							Time:           time.Date(2015, 5, 1, 18, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							Code:           "2",
							Name:           "退勤",
							DivisionName:   "本社",
							CredentialCode: kotclient.Ptr(300),
							CredentialName: kotclient.Ptr("KOTSL"),
							DivisionCode:   "1000",
							Latitude:       35.6672237,
							Longitude:      139.7422207,
						},
						{
							Time:         time.Date(2016, 5, 1, 10, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							Code:         "3",
							Name:         "休憩開始",
							DivisionName: "本社",
							DivisionCode: "1000",
						},
						{
							Time:         time.Date(2016, 5, 1, 11, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
							Code:         "4",
							Name:         "休憩終了",
							DivisionName: "本社",
							DivisionCode: "1000",
						},
					},
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDailyWorkingCosts(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	queries := kotclient.Ptr(kotclient.GetDailyWorkingCostsParams{
		Division:         kotclient.Ptr("1000"),
		Ondivision:       kotclient.Ptr(true),
		Start:            kotclient.Ptr(openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)}),
		End:              kotclient.Ptr(openapi_types.Date{Time: time.Date(2016, 5, 10, 0, 0, 0, 0, time.UTC)}),
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetDailyWorkingCostsWithResponse(ctx, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetDailyWorkingCostsWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetDailyWorkingCosts{
		{
			Date: openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
			DailyWorkings: []kotclient.DailyWorkingCostResponse{
				{
					Date:        openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
					EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
					CurrentDateEmployee: kotclient.Ptr(kotclient.DailyWorkingCurrentDateEmployee{
						Code:               "1000",
						DivisionCode:       "1000",
						DivisionName:       "本社",
						Gender:             kotclient.DailyWorkingCurrentDateEmployeeGenderMale,
						TypeCode:           "1",
						TypeName:           "正社員",
						LastName:           "勤怠",
						FirstName:          "太郎",
						LastNamePhonetics:  "キンタイ",
						FirstNamePhonetics: "タロウ",
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
					}),
					LaborCostEstimate:     8000,
					TransportationExpense: 1000,
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDailyWorkingCost(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	date := openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)}
	queries := kotclient.Ptr(kotclient.GetDailyWorkingCostParams{
		Division:         kotclient.Ptr("1000"),
		Ondivision:       kotclient.Ptr(true),
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetDailyWorkingCostWithResponse(ctx, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetDailyWorkingCostWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetDailyWorkingCost{
		Date: openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
		DailyWorkings: []kotclient.DailyWorkingCostResponse{
			{
				Date:        openapi_types.Date{Time: time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)},
				EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
				CurrentDateEmployee: kotclient.Ptr(kotclient.DailyWorkingCurrentDateEmployee{
					Code:               "1000",
					DivisionCode:       "1000",
					DivisionName:       "本社",
					Gender:             kotclient.DailyWorkingCurrentDateEmployeeGenderMale,
					TypeCode:           "1",
					TypeName:           "正社員",
					LastName:           "勤怠",
					FirstName:          "太郎",
					LastNamePhonetics:  "キンタイ",
					FirstNamePhonetics: "タロウ",
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
				}),
				LaborCostEstimate:     8000,
				TransportationExpense: 1000,
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetMonthlyWorking(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	date := openapi_types.Date{Time: time.Date(2016, 6, 1, 0, 0, 0, 0, time.UTC)}
	queries := kotclient.Ptr(kotclient.GetMonthlyWorkingParams{
		Division:         kotclient.Ptr("1000"),
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetMonthlyWorkingWithResponse(ctx, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetMonthlyWorkingWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetMonthlyWorking{
		{
			Year:        2016,
			Month:       6,
			EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
			CurrentDateEmployee: kotclient.Ptr(kotclient.MonthlyWorkingCurrentDateEmployee{
				DivisionCode:       "1000",
				DivisionName:       "本社",
				Gender:             kotclient.MonthlyWorkingCurrentDateEmployeeGenderMale,
				TypeCode:           "1",
				TypeName:           "正社員",
				Code:               "1000",
				LastName:           "勤怠",
				FirstName:          "太郎",
				LastNamePhonetics:  "キンタイ",
				FirstNamePhonetics: "タロウ",
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
			}),
			StartDate:              openapi_types.Date{Time: time.Date(2016, 6, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:                openapi_types.Date{Time: time.Date(2016, 6, 21, 0, 0, 0, 0, time.UTC)},
			WorkingCount:           14,
			WeekdayWorkingCount:    14,
			LateCount:              0,
			EarlyLeaveCount:        0,
			WorkingDayCount:        14,
			WeekdayWorkingdayCount: 14,
			AbsentdayCount:         0,
			HolidaysObtained: []kotclient.MonthlyWorkingHolidayObtained{
				{
					DayCount: 0,
					Minutes:  0,
					Code:     1,
					Name:     "有休",
				},
			},
			Assigned:      0,
			Unassigned:    0,
			Overtime:      0,
			Night:         0,
			NightOvertime: 0,
			BreakSum:      840,
			Late:          0,
			EarlyLeave:    0,
			HolidayWork: kotclient.MonthlyWorkingHolidayWork{
				Normal:        0,
				Night:         0,
				Overtime:      0,
				NightOvertime: 0,
				Extra:         0,
			},
			PremiumWork: kotclient.MonthlyWorkingPremiumWork{
				Overtime1:      0,
				NightOvertime1: 0,
				Overtime2:      0,
				NightOvertime2: 0,
			},
			LegalHolidayWork: kotclient.MonthlyWorkingLegalHolidayWork{
				Normal:        0,
				Night:         0,
				Overtime:      0,
				NightOvertime: 0,
				Extra:         0,
				Count:         0,
				DayCount:      0,
			},
			GeneralHolidayWork: kotclient.MonthlyWorkingGeneralHolidayWork{
				Normal:        0,
				Night:         0,
				Overtime:      0,
				NightOvertime: 0,
				Extra:         0,
				Count:         0,
				DayCount:      0,
			},
			Bind:      0,
			Regarding: 0,
			IsClosing: false,
			CustomMonthlyWorkings: []kotclient.MonthlyWorkingCustomMonthlyWorking{
				{
					Code:                "mCus3",
					Name:                "月別カスタム1",
					CalculationUnitCode: 1,
					CalculationUnitName: "日",
					CalculationResult:   1,
				},
				{
					Code:                "mCus2",
					Name:                "月別カスタム2",
					CalculationUnitCode: 2,
					CalculationUnitName: "時間",
					CalculationResult:   10,
				},
				{
					Code:                "mCus3",
					Name:                "月別カスタム3",
					CalculationUnitCode: 4,
					CalculationUnitName: "数値",
					CalculationResult:   100,
				},
			},
			IntervalShortageCount: 1,
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetMonthlyWorkingCost(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	date := "2016-06"
	queries := kotclient.Ptr(kotclient.GetMonthlyWorkingCostParams{
		Division:         kotclient.Ptr("1000"),
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetMonthlyWorkingCostWithResponse(ctx, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetMonthlyWorkingCostWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.MonthlyWorkingCostResponse{
		{
			Year:        2016,
			Month:       6,
			EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
			CurrentDateEmployee: kotclient.Ptr(kotclient.MonthlyWorkingCurrentDateEmployee{
				DivisionCode:       "1000",
				DivisionName:       "本社",
				Gender:             kotclient.MonthlyWorkingCurrentDateEmployeeGenderMale,
				TypeCode:           "1",
				TypeName:           "正社員",
				Code:               "1000",
				LastName:           "勤怠",
				FirstName:          "太郎",
				LastNamePhonetics:  "キンタイ",
				FirstNamePhonetics: "タロウ",
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
			}),
			StartDate:             openapi_types.Date{Time: time.Date(2016, 6, 1, 0, 0, 0, 0, time.UTC)},
			EndDate:               openapi_types.Date{Time: time.Date(2016, 6, 30, 0, 0, 0, 0, time.UTC)},
			LaborCostEstimate:     200000,
			TransportationExpense: 20000,
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetMonthlyWorkingHolidayRemained(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	employeeTypeCode := 1000
	date := "2016-06"
	queries := kotclient.Ptr(kotclient.GetMonthlyWorkingHolidayRemainedParams{
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetMonthlyWorkingHolidayRemainedWithResponse(ctx, employeeTypeCode, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetMonthlyWorkingHolidayRemainedWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := []kotclient.MonthlyWorkingHolidayRemainedResponse{
		{
			Date:        openapi_types.Date{Time: time.Date(2016, 6, 30, 0, 0, 0, 0, time.UTC)},
			CloseDate:   30,
			EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
			CurrentDateEmployee: kotclient.Ptr(kotclient.MonthlyWorkingCurrentDateEmployee{
				DivisionCode:       "1000",
				DivisionName:       "本社",
				Gender:             kotclient.MonthlyWorkingCurrentDateEmployeeGenderMale,
				TypeCode:           "1",
				TypeName:           "正社員",
				Code:               "1000",
				LastName:           "勤怠",
				FirstName:          "太郎",
				LastNamePhonetics:  "キンタイ",
				FirstNamePhonetics: "タロウ",
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
			}),
			HolidayRemained: []kotclient.MonthlyWorkingHolidayRemainedHolidayRemained{
				{
					Day:     0,
					Minutes: 0,
					Code:    1,
					Name:    "有休",
				},
				{
					Day:     0,
					Minutes: 0,
					Code:    2,
					Name:    "代休",
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetYearlyWorkingHoliday(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	employeeTypeCode := 1000
	year := 2019
	queries := kotclient.Ptr(kotclient.GetYearlyWorkingHolidayParams{
		AdditionalFields: kotclient.Ptr([]string{"currentDateEmployee"}),
	})
	got, err := sut.GetYearlyWorkingHolidayWithResponse(ctx, employeeTypeCode, year, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetYearlyWorkingHolidayWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetYearlyWorkingHoliday{
		Year:      "2019",
		StartDate: openapi_types.Date{Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		EndDate:   openapi_types.Date{Time: time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)},
		CloseDate: 30,
		Employees: []kotclient.YearlyWorkingEmployee{
			{
				EmployeeKey: "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3",
				CurrentDateEmployee: kotclient.Ptr(kotclient.YearlyWorkingCurrentDateEmployee{
					DivisionCode:       "1000",
					DivisionName:       "本社",
					Gender:             kotclient.YearlyWorkingCurrentDateEmployeeGenderMale,
					TypeCode:           "1",
					TypeName:           "正社員",
					Code:               "1000",
					LastName:           "勤怠",
					FirstName:          "太郎",
					LastNamePhonetics:  "キンタイ",
					FirstNamePhonetics: "タロウ",
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
				}),
				Holidays: []kotclient.YearlyWorkingEmployeeHoliday{
					{
						Code: 1,
						Name: "有休",
						Granted: []kotclient.YearlyWorkingEmployeeHolidayGranted{
							{
								Date:            openapi_types.Date{Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
								Days:            5,
								Minutes:         120,
								EffectivePeriod: openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
							},
							{
								Date:            openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
								Minutes:         180,
								EffectivePeriod: openapi_types.Date{Time: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC)},
							},
						},
						Obtained: []kotclient.YearlyWorkingEmployeeHolidayObtained{
							{
								Date: openapi_types.Date{Time: time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC)},
								Days: 1,
							},
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC)},
								Minutes: 180,
							},
						},
						Remained: []kotclient.YearlyWorkingEmployeeHolidayRemained{
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
								Days:    2,
								Minutes: 120,
							},
						},
						Expired: []kotclient.YearlyWorkingEmployeeHolidayExpired{
							{
								Date: openapi_types.Date{Time: time.Date(2019, 6, 1, 0, 0, 0, 0, time.UTC)},
								Days: 2,
							},
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 3, 1, 0, 0, 0, 0, time.UTC)},
								Minutes: 120,
							},
						},
					},
					{
						Code: 2,
						Name: "代休",
						Granted: []kotclient.YearlyWorkingEmployeeHolidayGranted{
							{
								Date:            openapi_types.Date{Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
								Days:            5,
								Minutes:         120,
								EffectivePeriod: openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
							},
							{
								Date:            openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
								Minutes:         180,
								EffectivePeriod: openapi_types.Date{Time: time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC)},
							},
						},
						Obtained: []kotclient.YearlyWorkingEmployeeHolidayObtained{
							{
								Date: openapi_types.Date{Time: time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC)},
								Days: 1,
							},
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC)},
								Minutes: 180,
							},
						},
						Remained: []kotclient.YearlyWorkingEmployeeHolidayRemained{
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
								Days:    2,
								Minutes: 120,
							},
						},
						Expired: []kotclient.YearlyWorkingEmployeeHolidayExpired{
							{
								Date: openapi_types.Date{Time: time.Date(2019, 6, 1, 0, 0, 0, 0, time.UTC)},
								Days: 2,
							},
							{
								Date:    openapi_types.Date{Time: time.Date(2019, 3, 1, 0, 0, 0, 0, time.UTC)},
								Minutes: 120,
							},
						},
					},
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}

func TestGetOvertime(t *testing.T) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken("8j9f7v4893y58rvt7nyfq2893n75tr78937n83")
	if err != nil {
		t.Fatalf("failed to securityprovider.NewSecurityProviderBearerToken: %v", err)
	}
	ctx := context.Background()
	sut, err := kotclient.NewClientWithResponses("http://localhost:8001", kotclient.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		t.Fatalf("failed to kotclient.NewClient: %v", err)
	}

	date := "2018-08"
	queries := kotclient.Ptr(kotclient.GetOvertimeParams{
		AdministratorKey: kotclient.Ptr("c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac222"),
		AdditionalFields: kotclient.Ptr([]kotclient.GetOvertimeParamsAdditionalFields{"flow"}),
	})
	got, err := sut.GetOvertimeWithResponse(ctx, date, queries)
	if err != nil {
		t.Fatalf("failed to sut.GetOvertimeWithResponse: %v", err)
	}

	if diff := cmp.Diff(200, got.StatusCode()); diff != "" {
		t.Errorf("value is mismatch (-want +got):\n%s", diff)
	}

	want := kotclient.GetOvertime{
		Year:  2018,
		Month: 8,
		OvertimeRequests: []kotclient.RequestOvertimeResponse{
			{
				Date:          openapi_types.Date{Time: time.Date(2018, 8, 1, 0, 0, 0, 0, time.UTC)},
				RequestedDate: openapi_types.Date{Time: time.Date(2018, 8, 10, 0, 0, 0, 0, time.UTC)},
				EmployeeKey:   "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe2735215201",
				RequestKey:    "59d25f49d4dce8b6e6658cc6b5c3b89b34c617916f90f6e27e9a9fa6cca576a4",
				Applicant: kotclient.RequestOvertimeApplicant{
					Type: "employee",
					Key:  "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe2735215201",
				},
				Status:      "applying",
				CurrentFlow: 2,
				Flow: kotclient.Ptr([]kotclient.RequestFlow{
					{
						Level: 1,
						AdministratorKeys: []string{
							"c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac6ff",
						},
					},
					{
						Level: 2,
						AdministratorKeys: []string{
							"c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac222",
						},
					},
					{
						Level: 3,
						AdministratorKeys: []string{
							"c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac111",
						},
					},
				}),
				Message:                      "申請いたします",
				AdminComment:                 "第1承認者承認済",
				LastModifiedAdministratorKey: "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac6ff",
				Requested: kotclient.RequestOvertimeRequested{
					IsBeforeSchedule: true,
					Start:            kotclient.KotDate{time.Date(2018, 8, 1, 8, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))},
					End:              kotclient.KotDate{time.Date(2018, 8, 1, 10, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))},
				},
				Current: kotclient.RequestOvertimeCurrent{
					IsBeforeSchedule: true,
					Start:            kotclient.KotDate{time.Date(2018, 8, 1, 9, 0, 0, 0, time.FixedZone("+09:00", 9*60*60))},
					End:              kotclient.KotDate{time.Date(2018, 8, 1, 10, 0, 0, 0, time.FixedZone("+09:00", 9*60*60))},
				},
			},
		},
	}
	if diff := cmp.Diff(kotclient.Ptr(want), got.JSON200); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}
