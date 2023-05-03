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


// AttendanceApiService AttendanceApi service
type AttendanceApiService service

type ApiGetDailyWorkingRequest struct {
	ctx context.Context
	ApiService *AttendanceApiService
	date string
	division *string
	ondivision *bool
	additionalFields *[]string
}

// 所属コード
func (r ApiGetDailyWorkingRequest) Division(division string) ApiGetDailyWorkingRequest {
	r.division = &division
	return r
}

// ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能
func (r ApiGetDailyWorkingRequest) Ondivision(ondivision bool) ApiGetDailyWorkingRequest {
	r.ondivision = &ondivision
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetDailyWorkingRequest) AdditionalFields(additionalFields []string) ApiGetDailyWorkingRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetDailyWorkingRequest) Execute() (*DailyWorkingResponse, *http.Response, error) {
	return r.ApiService.GetDailyWorkingExecute(r)
}

/*
GetDailyWorking Method for GetDailyWorking

指定した年月日の日別勤務データを取得する。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param date 詳細を取得したい年月日 ・過去日は最大3年前まで ・未来日は最大1年後まで
 @return ApiGetDailyWorkingRequest
*/
func (a *AttendanceApiService) GetDailyWorking(ctx context.Context, date string) ApiGetDailyWorkingRequest {
	return ApiGetDailyWorkingRequest{
		ApiService: a,
		ctx: ctx,
		date: date,
	}
}

// Execute executes the request
//  @return DailyWorkingResponse
func (a *AttendanceApiService) GetDailyWorkingExecute(r ApiGetDailyWorkingRequest) (*DailyWorkingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DailyWorkingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceApiService.GetDailyWorking")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/daily-workings/{date}"
	localVarPath = strings.Replace(localVarPath, "{"+"date"+"}", url.PathEscape(parameterValueToString(r.date, "date")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "")
	}
	if r.ondivision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ondivision", r.ondivision, "")
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

type ApiGetDailyWorkingTimerecordsRequest struct {
	ctx context.Context
	ApiService *AttendanceApiService
	employeeKeys *[]string
	division *string
	ondivision *bool
	start *string
	end *string
	additionalFields *[]string
}

// 従業員識別キー（従業員コードが変更されても不変） ,区切りで複数従業員の指定可能 divisionが指定されていない場合のみ使用可能
func (r ApiGetDailyWorkingTimerecordsRequest) EmployeeKeys(employeeKeys []string) ApiGetDailyWorkingTimerecordsRequest {
	r.employeeKeys = &employeeKeys
	return r
}

// 所属コード
func (r ApiGetDailyWorkingTimerecordsRequest) Division(division string) ApiGetDailyWorkingTimerecordsRequest {
	r.division = &division
	return r
}

// ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能
func (r ApiGetDailyWorkingTimerecordsRequest) Ondivision(ondivision bool) ApiGetDailyWorkingTimerecordsRequest {
	r.ondivision = &ondivision
	return r
}

// 取得したい期間の開始年月日 ・過去日は最大3年前まで
func (r ApiGetDailyWorkingTimerecordsRequest) Start(start string) ApiGetDailyWorkingTimerecordsRequest {
	r.start = &start
	return r
}

// 取得したい期間の終了年月日 ・startとendは同時に指定 ・期間は最大62日 ・未来日は最大1年後まで
func (r ApiGetDailyWorkingTimerecordsRequest) End(end string) ApiGetDailyWorkingTimerecordsRequest {
	r.end = &end
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetDailyWorkingTimerecordsRequest) AdditionalFields(additionalFields []string) ApiGetDailyWorkingTimerecordsRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetDailyWorkingTimerecordsRequest) Execute() ([]GetDailyWorkingTimerecords200ResponseInner, *http.Response, error) {
	return r.ApiService.GetDailyWorkingTimerecordsExecute(r)
}

/*
GetDailyWorkingTimerecords Method for GetDailyWorkingTimerecords

日別打刻データの一覧を取得する。 従業員を指定した場合、対象従業員の日別打刻データを取得する。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDailyWorkingTimerecordsRequest
*/
func (a *AttendanceApiService) GetDailyWorkingTimerecords(ctx context.Context) ApiGetDailyWorkingTimerecordsRequest {
	return ApiGetDailyWorkingTimerecordsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetDailyWorkingTimerecords200ResponseInner
func (a *AttendanceApiService) GetDailyWorkingTimerecordsExecute(r ApiGetDailyWorkingTimerecordsRequest) ([]GetDailyWorkingTimerecords200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetDailyWorkingTimerecords200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceApiService.GetDailyWorkingTimerecords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/daily-workings/timerecord/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeKeys != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeKeys", r.employeeKeys, "csv")
	}
	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "")
	}
	if r.ondivision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ondivision", r.ondivision, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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

type ApiGetDailyWorkingsRequest struct {
	ctx context.Context
	ApiService *AttendanceApiService
	division *string
	ondivision *bool
	start *string
	end *string
	additionalFields *[]string
}

// 所属コード
func (r ApiGetDailyWorkingsRequest) Division(division string) ApiGetDailyWorkingsRequest {
	r.division = &division
	return r
}

// ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能
func (r ApiGetDailyWorkingsRequest) Ondivision(ondivision bool) ApiGetDailyWorkingsRequest {
	r.ondivision = &ondivision
	return r
}

// 取得したい期間の開始年月日 ・過去日は最大3年前まで
func (r ApiGetDailyWorkingsRequest) Start(start string) ApiGetDailyWorkingsRequest {
	r.start = &start
	return r
}

// 取得したい期間の終了年月日 ・startとendは同時に指定 ・期間は最大62日 ・未来日は最大1年後まで
func (r ApiGetDailyWorkingsRequest) End(end string) ApiGetDailyWorkingsRequest {
	r.end = &end
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetDailyWorkingsRequest) AdditionalFields(additionalFields []string) ApiGetDailyWorkingsRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetDailyWorkingsRequest) Execute() ([]DailyWorkingResponse, *http.Response, error) {
	return r.ApiService.GetDailyWorkingsExecute(r)
}

/*
GetDailyWorkings Method for GetDailyWorkings

日別勤務データの一覧を取得する。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDailyWorkingsRequest
*/
func (a *AttendanceApiService) GetDailyWorkings(ctx context.Context) ApiGetDailyWorkingsRequest {
	return ApiGetDailyWorkingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DailyWorkingResponse
func (a *AttendanceApiService) GetDailyWorkingsExecute(r ApiGetDailyWorkingsRequest) ([]DailyWorkingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DailyWorkingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceApiService.GetDailyWorkings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/daily-workings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "")
	}
	if r.ondivision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ondivision", r.ondivision, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
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

type ApiRegisterDailyWorkingTimerecordRequest struct {
	ctx context.Context
	ApiService *AttendanceApiService
	employeeKey string
	dailyWorkingTimerecordRequest *DailyWorkingTimerecordRequest
}

func (r ApiRegisterDailyWorkingTimerecordRequest) DailyWorkingTimerecordRequest(dailyWorkingTimerecordRequest DailyWorkingTimerecordRequest) ApiRegisterDailyWorkingTimerecordRequest {
	r.dailyWorkingTimerecordRequest = &dailyWorkingTimerecordRequest
	return r
}

func (r ApiRegisterDailyWorkingTimerecordRequest) Execute() (*RegisterDailyWorkingTimerecord201Response, *http.Response, error) {
	return r.ApiService.RegisterDailyWorkingTimerecordExecute(r)
}

/*
RegisterDailyWorkingTimerecord Method for RegisterDailyWorkingTimerecord

指定した従業員の打刻データを登録する。(1件)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param employeeKey 従業員識別キー（従業員コードが変更されても不変）
 @return ApiRegisterDailyWorkingTimerecordRequest
*/
func (a *AttendanceApiService) RegisterDailyWorkingTimerecord(ctx context.Context, employeeKey string) ApiRegisterDailyWorkingTimerecordRequest {
	return ApiRegisterDailyWorkingTimerecordRequest{
		ApiService: a,
		ctx: ctx,
		employeeKey: employeeKey,
	}
}

// Execute executes the request
//  @return RegisterDailyWorkingTimerecord201Response
func (a *AttendanceApiService) RegisterDailyWorkingTimerecordExecute(r ApiRegisterDailyWorkingTimerecordRequest) (*RegisterDailyWorkingTimerecord201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegisterDailyWorkingTimerecord201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttendanceApiService.RegisterDailyWorkingTimerecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/daily-workings/timerecord/{employeeKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"employeeKey"+"}", url.PathEscape(parameterValueToString(r.employeeKey, "employeeKey")), -1)

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
	localVarPostBody = r.dailyWorkingTimerecordRequest
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
