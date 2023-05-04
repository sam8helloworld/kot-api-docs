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


// RequestApiService RequestApi service
type RequestApiService service

type ApiGetSchedulesRequest struct {
	ctx context.Context
	ApiService *RequestApiService
	date string
	administratorKey *string
	additionalFields *string
}

// 管理者識別キー（管理者コードが変更されても不変）
func (r ApiGetSchedulesRequest) AdministratorKey(administratorKey string) ApiGetSchedulesRequest {
	r.administratorKey = &administratorKey
	return r
}

// 指定されたプロパティをレスポンスに追加
func (r ApiGetSchedulesRequest) AdditionalFields(additionalFields string) ApiGetSchedulesRequest {
	r.additionalFields = &additionalFields
	return r
}

func (r ApiGetSchedulesRequest) Execute() (*GetSchedules200Response, *http.Response, error) {
	return r.ApiService.GetSchedulesExecute(r)
}

/*
GetSchedules /requests/schedules/{date}{?administratorKey,additionalFields}

指定した年月のスケジュール申請データを取得する。
※対象企業がスケジュール申請を利用できない場合にはエラーとなります。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param date 詳細を取得したい年月 ・過去日は最大3年前まで ・未来日は最大1年後まで
 @return ApiGetSchedulesRequest
*/
func (a *RequestApiService) GetSchedules(ctx context.Context, date string) ApiGetSchedulesRequest {
	return ApiGetSchedulesRequest{
		ApiService: a,
		ctx: ctx,
		date: date,
	}
}

// Execute executes the request
//  @return GetSchedules200Response
func (a *RequestApiService) GetSchedulesExecute(r ApiGetSchedulesRequest) (*GetSchedules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSchedules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestApiService.GetSchedules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests/schedules/{date}"
	localVarPath = strings.Replace(localVarPath, "{"+"date"+"}", url.PathEscape(parameterValueToString(r.date, "date")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.administratorKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "administratorKey", r.administratorKey, "")
	}
	if r.additionalFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "additionalFields", r.additionalFields, "")
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