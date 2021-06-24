// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	ent "entgo.io/contrib/entproto/internal/todo/ent"
	user "entgo.io/contrib/entproto/internal/todo/ent/user"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	strings "strings"
)

// UserService implements UserServiceServer
type UserService struct {
	client *ent.Client
	UnimplementedUserServiceServer
}

// NewUserService returns a new UserService
func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func toProtoUser_Status(e user.Status) User_Status {
	if v, ok := User_Status_value[strings.ToUpper(string(e))]; ok {
		return User_Status(v)
	}
	return User_Status(0)
}

func toEntUser_Status(e User_Status) user.Status {
	if v, ok := User_Status_name[int32(e)]; ok {
		return user.Status(strings.ToLower(v))
	}
	return ""
}

// toProtoUser transforms the ent type to the pb type
func toProtoUser(e *ent.User) *User {
	v := &User{
		Banned:     e.Banned,
		CrmId:      runtime.MustExtractUUIDBytes(e.CrmID),
		CustomPb:   uint64(e.CustomPb),
		Exp:        e.Exp,
		ExternalId: int32(e.ExternalID),
		Id:         int32(e.ID),
		Joined:     timestamppb.New(e.Joined),
		OptBool:    wrapperspb.Bool(e.OptBool),
		OptNum:     wrapperspb.Int32(int32(e.OptNum)),
		OptStr:     wrapperspb.String(e.OptStr),
		Points:     uint32(e.Points),
		Status:     toProtoUser_Status(e.Status),
		UserName:   e.UserName,
	}

	if edg := e.Edges.Attachment; edg != nil {
		v.Attachment = &Attachment{
			Id: runtime.MustExtractUUIDBytes(edg.ID),
		}
	}

	if edg := e.Edges.Group; edg != nil {
		v.Group = &Group{
			Id: int32(edg.ID),
		}
	}

	for _, edg := range e.Edges.Received {
		v.Received = append(v.Received, &Attachment{
			Id: runtime.MustExtractUUIDBytes(edg.ID),
		})
	}
	return v
}

// validateUser validates that all fields are encoded properly and are safe to pass
// to the ent entity builder.
func validateUser(x *User, checkId bool) error {
	if err := runtime.ValidateUUID(x.GetCrmId()); err != nil {
		return err
	}
	if err := runtime.ValidateUUID(x.GetAttachment().GetId()); err != nil {
		return err
	}
	for _, item := range x.GetReceived() {
		if err := runtime.ValidateUUID(item.GetId()); err != nil {
			return err
		}
	}

	return nil
}

// Create implements UserServiceServer.Create
func (svc *UserService) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	user := req.GetUser()
	if err := validateUser(user, false); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.User.Create()
	m.SetBanned(user.GetBanned())
	m.SetCrmID(runtime.MustBytesToUUID(user.GetCrmId()))
	m.SetCustomPb(uint8(user.GetCustomPb()))
	m.SetExp(uint64(user.GetExp()))
	m.SetExternalID(int(user.GetExternalId()))
	m.SetJoined(runtime.ExtractTime(user.GetJoined()))
	if user.GetOptBool() != nil {
		m.SetOptBool(user.GetOptBool().GetValue())
	}
	if user.GetOptNum() != nil {
		m.SetOptNum(int(user.GetOptNum().GetValue()))
	}
	if user.GetOptStr() != nil {
		m.SetOptStr(user.GetOptStr().GetValue())
	}
	m.SetPoints(uint(user.GetPoints()))
	m.SetStatus(toEntUser_Status(user.GetStatus()))
	m.SetUserName(user.GetUserName())
	m.SetAttachmentID(runtime.MustBytesToUUID(user.GetAttachment().GetId()))
	m.SetGroupID(int(user.GetGroup().GetId()))
	for _, item := range user.GetReceived() {
		m.AddReceivedIDs(runtime.MustBytesToUUID(item.GetId()))
	}

	res, err := m.Save(ctx)

	switch {
	case err == nil:
		return toProtoUser(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	var (
		err error
		get *ent.User
	)
	switch req.GetView() {
	case GetUserRequest_VIEW_UNSPECIFIED, GetUserRequest_BASIC:
		get, err = svc.client.User.Get(ctx, int(req.GetId()))
	case GetUserRequest_WITH_EDGE_IDS:
		get, err = svc.client.User.Query().
			WithAttachment(func(query *ent.AttachmentQuery) {
				query.Select("id")
			}).
			WithGroup(func(query *ent.GroupQuery) {
				query.Select("id")
			}).
			WithReceived(func(query *ent.AttachmentQuery) {
				query.Select("id")
			}).
			First(ctx)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUser(get), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements UserServiceServer.Update
func (svc *UserService) Update(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	user := req.GetUser()
	if err := validateUser(user, true); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.User.UpdateOneID(int(user.GetId()))
	m.SetBanned(user.GetBanned())
	m.SetCrmID(runtime.MustBytesToUUID(user.GetCrmId()))
	m.SetCustomPb(uint8(user.GetCustomPb()))
	m.SetExp(uint64(user.GetExp()))
	m.SetExternalID(int(user.GetExternalId()))
	if user.GetOptBool() != nil {
		m.SetOptBool(user.GetOptBool().GetValue())
	}
	if user.GetOptNum() != nil {
		m.SetOptNum(int(user.GetOptNum().GetValue()))
	}
	if user.GetOptStr() != nil {
		m.SetOptStr(user.GetOptStr().GetValue())
	}
	m.SetPoints(uint(user.GetPoints()))
	m.SetStatus(toEntUser_Status(user.GetStatus()))
	m.SetUserName(user.GetUserName())
	m.SetAttachmentID(runtime.MustBytesToUUID(user.GetAttachment().GetId()))
	m.SetGroupID(int(user.GetGroup().GetId()))
	for _, item := range user.GetReceived() {
		m.AddReceivedIDs(runtime.MustBytesToUUID(item.GetId()))
	}

	res, err := m.Save(ctx)

	switch {
	case err == nil:
		return toProtoUser(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

// Delete implements UserServiceServer.Delete
func (svc *UserService) Delete(ctx context.Context, req *DeleteUserRequest) (*emptypb.Empty, error) {
	err := svc.client.User.DeleteOneID(int(req.GetId())).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
