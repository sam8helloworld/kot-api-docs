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
		Gender:         kotclient.Male,
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
		Gender:         kotclient.Male,
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
