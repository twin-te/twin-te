// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: schoolcalendar/v1/service.proto

package schoolcalendarv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// SchoolCalendarServiceName is the fully-qualified name of the SchoolCalendarService service.
	SchoolCalendarServiceName = "schoolcalendar.v1.SchoolCalendarService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SchoolCalendarServiceGetEventsByDateProcedure is the fully-qualified name of the
	// SchoolCalendarService's GetEventsByDate RPC.
	SchoolCalendarServiceGetEventsByDateProcedure = "/schoolcalendar.v1.SchoolCalendarService/GetEventsByDate"
	// SchoolCalendarServiceGetModuleByDateProcedure is the fully-qualified name of the
	// SchoolCalendarService's GetModuleByDate RPC.
	SchoolCalendarServiceGetModuleByDateProcedure = "/schoolcalendar.v1.SchoolCalendarService/GetModuleByDate"
)

// SchoolCalendarServiceClient is a client for the schoolcalendar.v1.SchoolCalendarService service.
type SchoolCalendarServiceClient interface {
	GetEventsByDate(context.Context, *connect_go.Request[v1.GetEventsByDateRequest]) (*connect_go.Response[v1.GetEventsByDateResponse], error)
	GetModuleByDate(context.Context, *connect_go.Request[v1.GetModuleByDateRequest]) (*connect_go.Response[v1.GetModuleByDateResponse], error)
}

// NewSchoolCalendarServiceClient constructs a client for the
// schoolcalendar.v1.SchoolCalendarService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchoolCalendarServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SchoolCalendarServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &schoolCalendarServiceClient{
		getEventsByDate: connect_go.NewClient[v1.GetEventsByDateRequest, v1.GetEventsByDateResponse](
			httpClient,
			baseURL+SchoolCalendarServiceGetEventsByDateProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getModuleByDate: connect_go.NewClient[v1.GetModuleByDateRequest, v1.GetModuleByDateResponse](
			httpClient,
			baseURL+SchoolCalendarServiceGetModuleByDateProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// schoolCalendarServiceClient implements SchoolCalendarServiceClient.
type schoolCalendarServiceClient struct {
	getEventsByDate *connect_go.Client[v1.GetEventsByDateRequest, v1.GetEventsByDateResponse]
	getModuleByDate *connect_go.Client[v1.GetModuleByDateRequest, v1.GetModuleByDateResponse]
}

// GetEventsByDate calls schoolcalendar.v1.SchoolCalendarService.GetEventsByDate.
func (c *schoolCalendarServiceClient) GetEventsByDate(ctx context.Context, req *connect_go.Request[v1.GetEventsByDateRequest]) (*connect_go.Response[v1.GetEventsByDateResponse], error) {
	return c.getEventsByDate.CallUnary(ctx, req)
}

// GetModuleByDate calls schoolcalendar.v1.SchoolCalendarService.GetModuleByDate.
func (c *schoolCalendarServiceClient) GetModuleByDate(ctx context.Context, req *connect_go.Request[v1.GetModuleByDateRequest]) (*connect_go.Response[v1.GetModuleByDateResponse], error) {
	return c.getModuleByDate.CallUnary(ctx, req)
}

// SchoolCalendarServiceHandler is an implementation of the schoolcalendar.v1.SchoolCalendarService
// service.
type SchoolCalendarServiceHandler interface {
	GetEventsByDate(context.Context, *connect_go.Request[v1.GetEventsByDateRequest]) (*connect_go.Response[v1.GetEventsByDateResponse], error)
	GetModuleByDate(context.Context, *connect_go.Request[v1.GetModuleByDateRequest]) (*connect_go.Response[v1.GetModuleByDateResponse], error)
}

// NewSchoolCalendarServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchoolCalendarServiceHandler(svc SchoolCalendarServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	schoolCalendarServiceGetEventsByDateHandler := connect_go.NewUnaryHandler(
		SchoolCalendarServiceGetEventsByDateProcedure,
		svc.GetEventsByDate,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	schoolCalendarServiceGetModuleByDateHandler := connect_go.NewUnaryHandler(
		SchoolCalendarServiceGetModuleByDateProcedure,
		svc.GetModuleByDate,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/schoolcalendar.v1.SchoolCalendarService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SchoolCalendarServiceGetEventsByDateProcedure:
			schoolCalendarServiceGetEventsByDateHandler.ServeHTTP(w, r)
		case SchoolCalendarServiceGetModuleByDateProcedure:
			schoolCalendarServiceGetModuleByDateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSchoolCalendarServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSchoolCalendarServiceHandler struct{}

func (UnimplementedSchoolCalendarServiceHandler) GetEventsByDate(context.Context, *connect_go.Request[v1.GetEventsByDateRequest]) (*connect_go.Response[v1.GetEventsByDateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("schoolcalendar.v1.SchoolCalendarService.GetEventsByDate is not implemented"))
}

func (UnimplementedSchoolCalendarServiceHandler) GetModuleByDate(context.Context, *connect_go.Request[v1.GetModuleByDateRequest]) (*connect_go.Response[v1.GetModuleByDateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("schoolcalendar.v1.SchoolCalendarService.GetModuleByDate is not implemented"))
}