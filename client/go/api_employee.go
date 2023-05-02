/*
KING OF TIME WebAPI

KING OF TIME WebAPI specification https://developer.kingtime.jp/

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EmployeeApiService EmployeeApi service
type EmployeeApiService service

type ApiGetEmployeeRequest struct {
	ctx context.Context
	ApiService *EmployeeApiService
	employeeCode string
	date *string
	includeResigner *bool
	additionalFields *[]string
}

// 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで
func (r ApiGetEmployeeRequest) Date(date string) ApiGetEmployeeRequest {
	r.date = &date
	return r
}

// 指定された年月日時点で退職済みの従業員を含む場合 True
func (r ApiGetEmployeeRequest) IncludeResigner(includeResigner bool) ApiGetEmployeeRequest {
	r.includeResigner = &includeResigner
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetEmployeeRequest) AdditionalFields(additionalFields []string) ApiGetEmployeeRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetEmployeeRequest) Execute() (*EmployeeResponse, *http.Response, error) {
	return r.ApiService.GetEmployeeExecute(r)
}

/*
GetEmployee Method for GetEmployee

指定した従業員のデータを取得する。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param employeeCode 詳細を取得したい従業員コード
 @return ApiGetEmployeeRequest
*/
func (a *EmployeeApiService) GetEmployee(ctx context.Context, employeeCode string) ApiGetEmployeeRequest {
	return ApiGetEmployeeRequest{
		ApiService: a,
		ctx: ctx,
		employeeCode: employeeCode,
	}
}

// Execute executes the request
//  @return EmployeeResponse
func (a *EmployeeApiService) GetEmployeeExecute(r ApiGetEmployeeRequest) (*EmployeeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeeApiService.GetEmployee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees/{employeeCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"employeeCode"+"}", url.PathEscape(parameterValueToString(r.employeeCode, "employeeCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "")
	}
	if r.includeResigner != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeResigner", r.includeResigner, "")
	}
	if r.additionalFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalFields", r.additionalFields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEmployeesRequest struct {
	ctx context.Context
	ApiService *EmployeeApiService
	date *string
	division *string
	includeResigner *bool
	additionalFields *[]string
}

// 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで
func (r ApiGetEmployeesRequest) Date(date string) ApiGetEmployeesRequest {
	r.date = &date
	return r
}

// 所属コード
func (r ApiGetEmployeesRequest) Division(division string) ApiGetEmployeesRequest {
	r.division = &division
	return r
}

// 指定された年月日時点で退職済みの従業員を含む場合 True
func (r ApiGetEmployeesRequest) IncludeResigner(includeResigner bool) ApiGetEmployeesRequest {
	r.includeResigner = &includeResigner
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetEmployeesRequest) AdditionalFields(additionalFields []string) ApiGetEmployeesRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetEmployeesRequest) Execute() ([]EmployeeResponse, *http.Response, error) {
	return r.ApiService.GetEmployeesExecute(r)
}

/*
GetEmployees Method for GetEmployees

従業員データの一覧を取得する。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEmployeesRequest
*/
func (a *EmployeeApiService) GetEmployees(ctx context.Context) ApiGetEmployeesRequest {
	return ApiGetEmployeesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EmployeeResponse
func (a *EmployeeApiService) GetEmployeesExecute(r ApiGetEmployeesRequest) ([]EmployeeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EmployeeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeeApiService.GetEmployees")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "")
	}
	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "")
	}
	if r.includeResigner != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeResigner", r.includeResigner, "")
	}
	if r.additionalFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalFields", r.additionalFields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRegisterEmployeeRequest struct {
	ctx context.Context
	ApiService *EmployeeApiService
	employeeRequest *EmployeeRequest
}

func (r ApiRegisterEmployeeRequest) EmployeeRequest(employeeRequest EmployeeRequest) ApiRegisterEmployeeRequest {
	r.employeeRequest = &employeeRequest
	return r
}

func (r ApiRegisterEmployeeRequest) Execute() (*RegisterEmployee201Response, *http.Response, error) {
	return r.ApiService.RegisterEmployeeExecute(r)
}

/*
RegisterEmployee Method for RegisterEmployee

従業員のデータを登録する。(1件)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterEmployeeRequest
*/
func (a *EmployeeApiService) RegisterEmployee(ctx context.Context) ApiRegisterEmployeeRequest {
	return ApiRegisterEmployeeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RegisterEmployee201Response
func (a *EmployeeApiService) RegisterEmployeeExecute(r ApiRegisterEmployeeRequest) (*RegisterEmployee201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegisterEmployee201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeeApiService.RegisterEmployee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.employeeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
