// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: announcement/v1/service.proto

package announcementv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/twin-te/twin-te/back/handler/api/rpcgen/announcement/v1"
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
	// AnnouncementServiceName is the fully-qualified name of the AnnouncementService service.
	AnnouncementServiceName = "announcement.v1.AnnouncementService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AnnouncementServiceListAnnouncementsProcedure is the fully-qualified name of the
	// AnnouncementService's ListAnnouncements RPC.
	AnnouncementServiceListAnnouncementsProcedure = "/announcement.v1.AnnouncementService/ListAnnouncements"
	// AnnouncementServiceReadAnnouncementsProcedure is the fully-qualified name of the
	// AnnouncementService's ReadAnnouncements RPC.
	AnnouncementServiceReadAnnouncementsProcedure = "/announcement.v1.AnnouncementService/ReadAnnouncements"
)

// AnnouncementServiceClient is a client for the announcement.v1.AnnouncementService service.
type AnnouncementServiceClient interface {
	ListAnnouncements(context.Context, *connect_go.Request[v1.ListAnnouncementsRequest]) (*connect_go.Response[v1.ListAnnouncementsResponse], error)
	ReadAnnouncements(context.Context, *connect_go.Request[v1.ReadAnnouncementsRequest]) (*connect_go.Response[v1.ReadAnnouncementsResponse], error)
}

// NewAnnouncementServiceClient constructs a client for the announcement.v1.AnnouncementService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAnnouncementServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AnnouncementServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &announcementServiceClient{
		listAnnouncements: connect_go.NewClient[v1.ListAnnouncementsRequest, v1.ListAnnouncementsResponse](
			httpClient,
			baseURL+AnnouncementServiceListAnnouncementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		readAnnouncements: connect_go.NewClient[v1.ReadAnnouncementsRequest, v1.ReadAnnouncementsResponse](
			httpClient,
			baseURL+AnnouncementServiceReadAnnouncementsProcedure,
			opts...,
		),
	}
}

// announcementServiceClient implements AnnouncementServiceClient.
type announcementServiceClient struct {
	listAnnouncements *connect_go.Client[v1.ListAnnouncementsRequest, v1.ListAnnouncementsResponse]
	readAnnouncements *connect_go.Client[v1.ReadAnnouncementsRequest, v1.ReadAnnouncementsResponse]
}

// ListAnnouncements calls announcement.v1.AnnouncementService.ListAnnouncements.
func (c *announcementServiceClient) ListAnnouncements(ctx context.Context, req *connect_go.Request[v1.ListAnnouncementsRequest]) (*connect_go.Response[v1.ListAnnouncementsResponse], error) {
	return c.listAnnouncements.CallUnary(ctx, req)
}

// ReadAnnouncements calls announcement.v1.AnnouncementService.ReadAnnouncements.
func (c *announcementServiceClient) ReadAnnouncements(ctx context.Context, req *connect_go.Request[v1.ReadAnnouncementsRequest]) (*connect_go.Response[v1.ReadAnnouncementsResponse], error) {
	return c.readAnnouncements.CallUnary(ctx, req)
}

// AnnouncementServiceHandler is an implementation of the announcement.v1.AnnouncementService
// service.
type AnnouncementServiceHandler interface {
	ListAnnouncements(context.Context, *connect_go.Request[v1.ListAnnouncementsRequest]) (*connect_go.Response[v1.ListAnnouncementsResponse], error)
	ReadAnnouncements(context.Context, *connect_go.Request[v1.ReadAnnouncementsRequest]) (*connect_go.Response[v1.ReadAnnouncementsResponse], error)
}

// NewAnnouncementServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAnnouncementServiceHandler(svc AnnouncementServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	announcementServiceListAnnouncementsHandler := connect_go.NewUnaryHandler(
		AnnouncementServiceListAnnouncementsProcedure,
		svc.ListAnnouncements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	announcementServiceReadAnnouncementsHandler := connect_go.NewUnaryHandler(
		AnnouncementServiceReadAnnouncementsProcedure,
		svc.ReadAnnouncements,
		opts...,
	)
	return "/announcement.v1.AnnouncementService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AnnouncementServiceListAnnouncementsProcedure:
			announcementServiceListAnnouncementsHandler.ServeHTTP(w, r)
		case AnnouncementServiceReadAnnouncementsProcedure:
			announcementServiceReadAnnouncementsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAnnouncementServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAnnouncementServiceHandler struct{}

func (UnimplementedAnnouncementServiceHandler) ListAnnouncements(context.Context, *connect_go.Request[v1.ListAnnouncementsRequest]) (*connect_go.Response[v1.ListAnnouncementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("announcement.v1.AnnouncementService.ListAnnouncements is not implemented"))
}

func (UnimplementedAnnouncementServiceHandler) ReadAnnouncements(context.Context, *connect_go.Request[v1.ReadAnnouncementsRequest]) (*connect_go.Response[v1.ReadAnnouncementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("announcement.v1.AnnouncementService.ReadAnnouncements is not implemented"))
}
