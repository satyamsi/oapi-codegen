// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /json)
	JSONExample(c *gin.Context)

	// (POST /multipart)
	MultipartExample(c *gin.Context)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(c *gin.Context)

	// (POST /reusable-responses)
	ReusableResponses(c *gin.Context)

	// (POST /text)
	TextExample(c *gin.Context)

	// (POST /unknown)
	UnknownExample(c *gin.Context)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(c *gin.Context)

	// (POST /urlencoded)
	URLEncodedExample(c *gin.Context)

	// (POST /with-headers)
	HeadersExample(c *gin.Context, params HeadersExampleParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// JSONExample operation middleware
func (siw *ServerInterfaceWrapper) JSONExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.JSONExample(c)
}

// MultipartExample operation middleware
func (siw *ServerInterfaceWrapper) MultipartExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.MultipartExample(c)
}

// MultipleRequestAndResponseTypes operation middleware
func (siw *ServerInterfaceWrapper) MultipleRequestAndResponseTypes(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.MultipleRequestAndResponseTypes(c)
}

// ReusableResponses operation middleware
func (siw *ServerInterfaceWrapper) ReusableResponses(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReusableResponses(c)
}

// TextExample operation middleware
func (siw *ServerInterfaceWrapper) TextExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.TextExample(c)
}

// UnknownExample operation middleware
func (siw *ServerInterfaceWrapper) UnknownExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UnknownExample(c)
}

// UnspecifiedContentType operation middleware
func (siw *ServerInterfaceWrapper) UnspecifiedContentType(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UnspecifiedContentType(c)
}

// URLEncodedExample operation middleware
func (siw *ServerInterfaceWrapper) URLEncodedExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.URLEncodedExample(c)
}

// HeadersExample operation middleware
func (siw *ServerInterfaceWrapper) HeadersExample(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params HeadersExampleParams

	headers := c.Request.Header

	// ------------- Required header parameter "header1" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header1")]; found {
		var Header1 string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for header1, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header1", runtime.ParamLocationHeader, valueList[0], &Header1)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter header1: %s", err), http.StatusBadRequest)
			return
		}

		params.Header1 = Header1

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Header parameter header1 is required, but not found: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional header parameter "header2" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header2")]; found {
		var Header2 int
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for header2, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header2", runtime.ParamLocationHeader, valueList[0], &Header2)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter header2: %s", err), http.StatusBadRequest)
			return
		}

		params.Header2 = &Header2

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.HeadersExample(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/json", wrapper.JSONExample)

	router.POST(options.BaseURL+"/multipart", wrapper.MultipartExample)

	router.POST(options.BaseURL+"/multiple", wrapper.MultipleRequestAndResponseTypes)

	router.POST(options.BaseURL+"/reusable-responses", wrapper.ReusableResponses)

	router.POST(options.BaseURL+"/text", wrapper.TextExample)

	router.POST(options.BaseURL+"/unknown", wrapper.UnknownExample)

	router.POST(options.BaseURL+"/unspecified-content-type", wrapper.UnspecifiedContentType)

	router.POST(options.BaseURL+"/urlencoded", wrapper.URLEncodedExample)

	router.POST(options.BaseURL+"/with-headers", wrapper.HeadersExample)

	return router
}

type BadrequestResponse struct {
}

type ReusableresponseResponseHeaders struct {
	Header1 string
	Header2 int
}
type ReusableresponseJSONResponse struct {
	Body Example

	Headers ReusableresponseResponseHeaders
}

type JSONExampleRequestObject struct {
	Body *JSONExampleJSONRequestBody
}

type JSONExampleResponseObject interface {
	VisitJSONExampleResponse(w http.ResponseWriter) error
}

type JSONExample200JSONResponse Example

func (response JSONExample200JSONResponse) VisitJSONExampleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type JSONExample400Response = BadrequestResponse

func (response JSONExample400Response) VisitJSONExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type JSONExampledefaultResponse struct {
	StatusCode int
}

func (response JSONExampledefaultResponse) VisitJSONExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type MultipartExampleRequestObject struct {
	Body *multipart.Reader
}

type MultipartExampleResponseObject interface {
	VisitMultipartExampleResponse(w http.ResponseWriter) error
}

type MultipartExample200MultipartResponse func(writer *multipart.Writer) error

func (response MultipartExample200MultipartResponse) VisitMultipartExampleResponse(w http.ResponseWriter) error {
	writer := multipart.NewWriter(w)
	w.Header().Set("Content-Type", writer.FormDataContentType())
	w.WriteHeader(200)

	defer writer.Close()
	return response(writer)
}

type MultipartExample400Response = BadrequestResponse

func (response MultipartExample400Response) VisitMultipartExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type MultipartExampledefaultResponse struct {
	StatusCode int
}

func (response MultipartExampledefaultResponse) VisitMultipartExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type MultipleRequestAndResponseTypesRequestObject struct {
	JSONBody      *MultipleRequestAndResponseTypesJSONRequestBody
	FormdataBody  *MultipleRequestAndResponseTypesFormdataRequestBody
	Body          io.Reader
	MultipartBody *multipart.Reader
	TextBody      *MultipleRequestAndResponseTypesTextRequestBody
}

type MultipleRequestAndResponseTypesResponseObject interface {
	VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error
}

type MultipleRequestAndResponseTypes200JSONResponse Example

func (response MultipleRequestAndResponseTypes200JSONResponse) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type MultipleRequestAndResponseTypes200FormdataResponse Example

func (response MultipleRequestAndResponseTypes200FormdataResponse) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.WriteHeader(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := w.Write([]byte(form.Encode()))
		return err
	}
}

type MultipleRequestAndResponseTypes200ImagepngResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response MultipleRequestAndResponseTypes200ImagepngResponse) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "image/png")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type MultipleRequestAndResponseTypes200MultipartResponse func(writer *multipart.Writer) error

func (response MultipleRequestAndResponseTypes200MultipartResponse) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	writer := multipart.NewWriter(w)
	w.Header().Set("Content-Type", writer.FormDataContentType())
	w.WriteHeader(200)

	defer writer.Close()
	return response(writer)
}

type MultipleRequestAndResponseTypes200TextResponse string

func (response MultipleRequestAndResponseTypes200TextResponse) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type MultipleRequestAndResponseTypes400Response = BadrequestResponse

func (response MultipleRequestAndResponseTypes400Response) VisitMultipleRequestAndResponseTypesResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ReusableResponsesRequestObject struct {
	Body *ReusableResponsesJSONRequestBody
}

type ReusableResponsesResponseObject interface {
	VisitReusableResponsesResponse(w http.ResponseWriter) error
}

type ReusableResponses200JSONResponse struct{ ReusableresponseJSONResponse }

func (response ReusableResponses200JSONResponse) VisitReusableResponsesResponse(w http.ResponseWriter) error {
	w.Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	w.Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type ReusableResponses400Response = BadrequestResponse

func (response ReusableResponses400Response) VisitReusableResponsesResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type ReusableResponsesdefaultResponse struct {
	StatusCode int
}

func (response ReusableResponsesdefaultResponse) VisitReusableResponsesResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type TextExampleRequestObject struct {
	Body *TextExampleTextRequestBody
}

type TextExampleResponseObject interface {
	VisitTextExampleResponse(w http.ResponseWriter) error
}

type TextExample200TextResponse string

func (response TextExample200TextResponse) VisitTextExampleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type TextExample400Response = BadrequestResponse

func (response TextExample400Response) VisitTextExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type TextExampledefaultResponse struct {
	StatusCode int
}

func (response TextExampledefaultResponse) VisitTextExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type UnknownExampleRequestObject struct {
	Body io.Reader
}

type UnknownExampleResponseObject interface {
	VisitUnknownExampleResponse(w http.ResponseWriter) error
}

type UnknownExample200Videomp4Response struct {
	Body          io.Reader
	ContentLength int64
}

func (response UnknownExample200Videomp4Response) VisitUnknownExampleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "video/mp4")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type UnknownExample400Response = BadrequestResponse

func (response UnknownExample400Response) VisitUnknownExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UnknownExampledefaultResponse struct {
	StatusCode int
}

func (response UnknownExampledefaultResponse) VisitUnknownExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type UnspecifiedContentTypeRequestObject struct {
	ContentType string
	Body        io.Reader
}

type UnspecifiedContentTypeResponseObject interface {
	VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error
}

type UnspecifiedContentType200VideoResponse struct {
	Body          io.Reader
	ContentType   string
	ContentLength int64
}

func (response UnspecifiedContentType200VideoResponse) VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", response.ContentType)
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type UnspecifiedContentType400Response = BadrequestResponse

func (response UnspecifiedContentType400Response) VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UnspecifiedContentType401Response struct {
}

func (response UnspecifiedContentType401Response) VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type UnspecifiedContentType403Response struct {
}

func (response UnspecifiedContentType403Response) VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type UnspecifiedContentTypedefaultResponse struct {
	StatusCode int
}

func (response UnspecifiedContentTypedefaultResponse) VisitUnspecifiedContentTypeResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type URLEncodedExampleRequestObject struct {
	Body *URLEncodedExampleFormdataRequestBody
}

type URLEncodedExampleResponseObject interface {
	VisitURLEncodedExampleResponse(w http.ResponseWriter) error
}

type URLEncodedExample200FormdataResponse Example

func (response URLEncodedExample200FormdataResponse) VisitURLEncodedExampleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.WriteHeader(200)

	if form, err := runtime.MarshalForm(response, nil); err != nil {
		return err
	} else {
		_, err := w.Write([]byte(form.Encode()))
		return err
	}
}

type URLEncodedExample400Response = BadrequestResponse

func (response URLEncodedExample400Response) VisitURLEncodedExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type URLEncodedExampledefaultResponse struct {
	StatusCode int
}

func (response URLEncodedExampledefaultResponse) VisitURLEncodedExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type HeadersExampleRequestObject struct {
	Params HeadersExampleParams
	Body   *HeadersExampleJSONRequestBody
}

type HeadersExampleResponseObject interface {
	VisitHeadersExampleResponse(w http.ResponseWriter) error
}

type HeadersExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type HeadersExample200JSONResponse struct {
	Body    Example
	Headers HeadersExample200ResponseHeaders
}

func (response HeadersExample200JSONResponse) VisitHeadersExampleResponse(w http.ResponseWriter) error {
	w.Header().Set("header1", fmt.Sprint(response.Headers.Header1))
	w.Header().Set("header2", fmt.Sprint(response.Headers.Header2))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type HeadersExample400Response = BadrequestResponse

func (response HeadersExample400Response) VisitHeadersExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type HeadersExampledefaultResponse struct {
	StatusCode int
}

func (response HeadersExampledefaultResponse) VisitHeadersExampleResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /json)
	JSONExample(ctx context.Context, request JSONExampleRequestObject) (JSONExampleResponseObject, error)

	// (POST /multipart)
	MultipartExample(ctx context.Context, request MultipartExampleRequestObject) (MultipartExampleResponseObject, error)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx context.Context, request MultipleRequestAndResponseTypesRequestObject) (MultipleRequestAndResponseTypesResponseObject, error)

	// (POST /reusable-responses)
	ReusableResponses(ctx context.Context, request ReusableResponsesRequestObject) (ReusableResponsesResponseObject, error)

	// (POST /text)
	TextExample(ctx context.Context, request TextExampleRequestObject) (TextExampleResponseObject, error)

	// (POST /unknown)
	UnknownExample(ctx context.Context, request UnknownExampleRequestObject) (UnknownExampleResponseObject, error)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(ctx context.Context, request UnspecifiedContentTypeRequestObject) (UnspecifiedContentTypeResponseObject, error)

	// (POST /urlencoded)
	URLEncodedExample(ctx context.Context, request URLEncodedExampleRequestObject) (URLEncodedExampleResponseObject, error)

	// (POST /with-headers)
	HeadersExample(ctx context.Context, request HeadersExampleRequestObject) (HeadersExampleResponseObject, error)
}

type StrictHandlerFunc func(ctx *gin.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// JSONExample operation middleware
func (sh *strictHandler) JSONExample(ctx *gin.Context) {
	var request JSONExampleRequestObject

	var body JSONExampleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.JSONExample(ctx, request.(JSONExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JSONExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(JSONExampleResponseObject); ok {
		if err := validResponse.VisitJSONExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// MultipartExample operation middleware
func (sh *strictHandler) MultipartExample(ctx *gin.Context) {
	var request MultipartExampleRequestObject

	if reader, err := ctx.Request.MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.Error(err)
		return
	}

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MultipartExample(ctx, request.(MultipartExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipartExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(MultipartExampleResponseObject); ok {
		if err := validResponse.VisitMultipartExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// MultipleRequestAndResponseTypes operation middleware
func (sh *strictHandler) MultipleRequestAndResponseTypes(ctx *gin.Context) {
	var request MultipleRequestAndResponseTypesRequestObject

	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/json") {
		var body MultipleRequestAndResponseTypesJSONRequestBody
		if err := ctx.Bind(&body); err != nil {
			ctx.Error(err)
			return
		}
		request.JSONBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/x-www-form-urlencoded") {
		if err := ctx.Request.ParseForm(); err != nil {
			ctx.Error(err)
			return
		}
		var body MultipleRequestAndResponseTypesFormdataRequestBody
		if err := runtime.BindForm(&body, ctx.Request.Form, nil, nil); err != nil {
			ctx.Error(err)
			return
		}
		request.FormdataBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "image/png") {
		request.Body = ctx.Request.Body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "multipart/form-data") {
		if reader, err := ctx.Request.MultipartReader(); err == nil {
			request.MultipartBody = reader
		} else {
			ctx.Error(err)
			return
		}
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "text/plain") {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Error(err)
			return
		}
		body := MultipleRequestAndResponseTypesTextRequestBody(data)
		request.TextBody = &body
	}

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MultipleRequestAndResponseTypes(ctx, request.(MultipleRequestAndResponseTypesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipleRequestAndResponseTypes")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(MultipleRequestAndResponseTypesResponseObject); ok {
		if err := validResponse.VisitMultipleRequestAndResponseTypesResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// ReusableResponses operation middleware
func (sh *strictHandler) ReusableResponses(ctx *gin.Context) {
	var request ReusableResponsesRequestObject

	var body ReusableResponsesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ReusableResponses(ctx, request.(ReusableResponsesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReusableResponses")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(ReusableResponsesResponseObject); ok {
		if err := validResponse.VisitReusableResponsesResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// TextExample operation middleware
func (sh *strictHandler) TextExample(ctx *gin.Context) {
	var request TextExampleRequestObject

	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error(err)
		return
	}
	body := TextExampleTextRequestBody(data)
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TextExample(ctx, request.(TextExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TextExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(TextExampleResponseObject); ok {
		if err := validResponse.VisitTextExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// UnknownExample operation middleware
func (sh *strictHandler) UnknownExample(ctx *gin.Context) {
	var request UnknownExampleRequestObject

	request.Body = ctx.Request.Body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnknownExample(ctx, request.(UnknownExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnknownExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(UnknownExampleResponseObject); ok {
		if err := validResponse.VisitUnknownExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// UnspecifiedContentType operation middleware
func (sh *strictHandler) UnspecifiedContentType(ctx *gin.Context) {
	var request UnspecifiedContentTypeRequestObject

	request.ContentType = ctx.ContentType()

	request.Body = ctx.Request.Body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UnspecifiedContentType(ctx, request.(UnspecifiedContentTypeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnspecifiedContentType")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(UnspecifiedContentTypeResponseObject); ok {
		if err := validResponse.VisitUnspecifiedContentTypeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// URLEncodedExample operation middleware
func (sh *strictHandler) URLEncodedExample(ctx *gin.Context) {
	var request URLEncodedExampleRequestObject

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.Error(err)
		return
	}
	var body URLEncodedExampleFormdataRequestBody
	if err := runtime.BindForm(&body, ctx.Request.Form, nil, nil); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.URLEncodedExample(ctx, request.(URLEncodedExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "URLEncodedExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(URLEncodedExampleResponseObject); ok {
		if err := validResponse.VisitURLEncodedExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// HeadersExample operation middleware
func (sh *strictHandler) HeadersExample(ctx *gin.Context, params HeadersExampleParams) {
	var request HeadersExampleRequestObject

	request.Params = params

	var body HeadersExampleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.HeadersExample(ctx, request.(HeadersExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HeadersExample")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(HeadersExampleResponseObject); ok {
		if err := validResponse.VisitHeadersExampleResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYS4/iOBD+K1btnkaB0D194rbTGmnfI9Ezp9UcirgAzya2166QRoj/vnJsaBjSCFo8",
	"pNXeEqde/qq+KsdLKExljSbNHoZLcOSt0Z7alzFKR//U5Dm8SfKFU5aV0TCEDyhH6dsqA0e1x3FJa/Ug",
	"XxjNpFtVtLZUBQbV/JsP+kvwxYwqDE8/OprAEH7IX0LJ41ef0zNWtiRYrVbZdxF8+g0ymBFKcm208fFu",
	"1zYvLMEQPDulpxCMRLH7TjGlmabkgrcgmoIIAus4hkuwzlhyrCJGcyxr6vaUVsz4GxUcd6D0xOxj+Wg0",
	"o9JeSDWZkCPNIoEngg0vfG2tcUxSjBcieChYeHJzcpABKw6BwdP2ukgBe8hgTs5HR3f9QX8Q8mUsabQK",
	"hvC+XcrAIs/aDW0SZE1X3n99+vSnUF5gzaZCVgWW5UJU6PwMy5KkUJpNiLEu2PehdeXazP8ik/rHhGUo",
	"m7aCPhi5uETFtIW5Vc/3g8GVCnOVwUN01mVjE1S+xbDWzATrsgP0L/pvbRotyDnj0s7yqi5ZWXS8naxd",
	"tP9YixwD+cZePjGu6klkvBDq5/J0U+BTM+gkydPMNF7MTCPYCElYikbxTKwVv2O30gKFV3paklgHlXVm",
	"sqTUc3/ScpT28jnYuDiXsh0rz72maXpt8mpXki6MJPk2s6rCKeVWT3fVg21kGMJ4waFs97vrmYooA6Zn",
	"zm2JSh8eHVdqJ/8jfTZiR7quzya9neR1E3dNKi8K1GIc+DjxgcRdvvZIOkqeRlsStxlxhzHaO61do2uG",
	"5L8+qT7T81FD6oxkvXY1ngpYHRdfxyxpHQPbG7l/BIpzJcnklX040fLNQPWWCjVRJHtpF70Y22st4dHo",
	"whHvDu1wAtaGxcZYOJjzjEREIBPeiIZEVXsWFr0XitsuUqp4uJe01zy+vET2GD2FyX5EVt9dKKfvbpXR",
	"h8Hd6SrvL1w3O8P3FT6Ofv8YZU79wznblD/xjHI+vzeiczhW97buALop/HMUeJnpBak5SYFaCkdcO01S",
	"zBWuf1v3uJkMvKTVosOKuPX61xLCCEkXC5CBxoo273epCJQLyLKrKTt0PXHQ1j1kh+4svv6Hf6gvedNz",
	"6TpdZRAvZWKx1K4MGWW2wzyPlzl93+B0Sq6vTI5Wwerr6t8AAAD//ygSomqZEwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
