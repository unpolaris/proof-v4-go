# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAutoProofsPost**](ProofApi.md#ApiAutoProofsPost) | **Post** /api/auto_proofs | 添加批量自动存证
[**V1ProofGetProofsPost**](ProofApi.md#V1ProofGetProofsPost) | **Post** /v1/proof/GetProofs | 获取多个指定hash的存证信息
[**V1ProofListPost**](ProofApi.md#V1ProofListPost) | **Post** /v1/proof/List | 获取存证列表

# **ApiAutoProofsPost**
> InlineResponse200 ApiAutoProofsPost(ctx, body, optional)
添加批量自动存证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelReqAutoProof**](model.ReqAutoProof.md)| 添加批量自动存证 | 
 **optional** | ***ProofApiApiAutoProofsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProofApiApiAutoProofsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fZMCaAppId** | **optional.**| 应用id | 
 **fZMCaAppKey** | **optional.**| 应用公钥 | 
 **fZMCaSign** | **optional.**| 签名 | 
 **fZMCaNonce** | **optional.**| 随机字符串 | 
 **fZMCaTimestamp** | **optional.**| 时间戳 | 
 **fZMCaSignType** | **optional.**| 签名方式，HMAC-SHA256（默认）或MD5 | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProofGetProofsPost**
> SwaggerListProofResult V1ProofGetProofsPost(ctx, body)
获取多个指定hash的存证信息

get proof by hashes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProofGetProofsBody**](ProofGetProofsBody.md)| INPUT | 

### Return type

[**SwaggerListProofResult**](swagger.ListProofResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProofListPost**
> SwaggerListProofResult V1ProofListPost(ctx, body)
获取存证列表

list proof of organization/sender

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProofListBody**](ProofListBody.md)| INPUT | 

### Return type

[**SwaggerListProofResult**](swagger.ListProofResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

