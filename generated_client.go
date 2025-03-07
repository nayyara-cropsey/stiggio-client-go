// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package stigg

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type StiggClient interface {
	GetPlans(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetPlans, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) StiggClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type GetPlans_Plans_Edges_Node struct {
	CreatedAt     *string       "json:\"createdAt,omitempty\" graphql:\"createdAt\""
	Description   *string       "json:\"description,omitempty\" graphql:\"description\""
	DisplayName   string        "json:\"displayName\" graphql:\"displayName\""
	ID            string        "json:\"id\" graphql:\"id\""
	IsLatest      *bool         "json:\"isLatest,omitempty\" graphql:\"isLatest\""
	ProductID     *string       "json:\"productId,omitempty\" graphql:\"productId\""
	RefID         string        "json:\"refId\" graphql:\"refId\""
	Status        PackageStatus "json:\"status\" graphql:\"status\""
	Type          string        "json:\"type\" graphql:\"type\""
	UpdatedAt     *string       "json:\"updatedAt,omitempty\" graphql:\"updatedAt\""
	VersionNumber int64         "json:\"versionNumber\" graphql:\"versionNumber\""
}

func (t *GetPlans_Plans_Edges_Node) GetCreatedAt() *string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.CreatedAt
}
func (t *GetPlans_Plans_Edges_Node) GetDescription() *string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.Description
}
func (t *GetPlans_Plans_Edges_Node) GetDisplayName() string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.DisplayName
}
func (t *GetPlans_Plans_Edges_Node) GetID() string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.ID
}
func (t *GetPlans_Plans_Edges_Node) GetIsLatest() *bool {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.IsLatest
}
func (t *GetPlans_Plans_Edges_Node) GetProductID() *string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.ProductID
}
func (t *GetPlans_Plans_Edges_Node) GetRefID() string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.RefID
}
func (t *GetPlans_Plans_Edges_Node) GetStatus() *PackageStatus {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return &t.Status
}
func (t *GetPlans_Plans_Edges_Node) GetType() string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.Type
}
func (t *GetPlans_Plans_Edges_Node) GetUpdatedAt() *string {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.UpdatedAt
}
func (t *GetPlans_Plans_Edges_Node) GetVersionNumber() int64 {
	if t == nil {
		t = &GetPlans_Plans_Edges_Node{}
	}
	return t.VersionNumber
}

type GetPlans_Plans_Edges struct {
	Cursor string                    "json:\"cursor\" graphql:\"cursor\""
	Node   GetPlans_Plans_Edges_Node "json:\"node\" graphql:\"node\""
}

func (t *GetPlans_Plans_Edges) GetCursor() string {
	if t == nil {
		t = &GetPlans_Plans_Edges{}
	}
	return t.Cursor
}
func (t *GetPlans_Plans_Edges) GetNode() *GetPlans_Plans_Edges_Node {
	if t == nil {
		t = &GetPlans_Plans_Edges{}
	}
	return &t.Node
}

type GetPlans_Plans_PageInfo struct {
	EndCursor       *string "json:\"endCursor,omitempty\" graphql:\"endCursor\""
	HasNextPage     *bool   "json:\"hasNextPage,omitempty\" graphql:\"hasNextPage\""
	HasPreviousPage *bool   "json:\"hasPreviousPage,omitempty\" graphql:\"hasPreviousPage\""
	StartCursor     *string "json:\"startCursor,omitempty\" graphql:\"startCursor\""
}

func (t *GetPlans_Plans_PageInfo) GetEndCursor() *string {
	if t == nil {
		t = &GetPlans_Plans_PageInfo{}
	}
	return t.EndCursor
}
func (t *GetPlans_Plans_PageInfo) GetHasNextPage() *bool {
	if t == nil {
		t = &GetPlans_Plans_PageInfo{}
	}
	return t.HasNextPage
}
func (t *GetPlans_Plans_PageInfo) GetHasPreviousPage() *bool {
	if t == nil {
		t = &GetPlans_Plans_PageInfo{}
	}
	return t.HasPreviousPage
}
func (t *GetPlans_Plans_PageInfo) GetStartCursor() *string {
	if t == nil {
		t = &GetPlans_Plans_PageInfo{}
	}
	return t.StartCursor
}

type GetPlans_Plans struct {
	Edges      []*GetPlans_Plans_Edges "json:\"edges\" graphql:\"edges\""
	PageInfo   GetPlans_Plans_PageInfo "json:\"pageInfo\" graphql:\"pageInfo\""
	TotalCount int64                   "json:\"totalCount\" graphql:\"totalCount\""
}

func (t *GetPlans_Plans) GetEdges() []*GetPlans_Plans_Edges {
	if t == nil {
		t = &GetPlans_Plans{}
	}
	return t.Edges
}
func (t *GetPlans_Plans) GetPageInfo() *GetPlans_Plans_PageInfo {
	if t == nil {
		t = &GetPlans_Plans{}
	}
	return &t.PageInfo
}
func (t *GetPlans_Plans) GetTotalCount() int64 {
	if t == nil {
		t = &GetPlans_Plans{}
	}
	return t.TotalCount
}

type GetPlans struct {
	Plans GetPlans_Plans "json:\"plans\" graphql:\"plans\""
}

func (t *GetPlans) GetPlans() *GetPlans_Plans {
	if t == nil {
		t = &GetPlans{}
	}
	return &t.Plans
}

const GetPlansDocument = `query GetPlans {
	plans(filter: {status:{eq:PUBLISHED},and:{isLatest:{is:true}}}) {
		totalCount
		edges {
			cursor
			node {
				id
				refId
				displayName
				description
				isLatest
				productId
				status
				type
				createdAt
				updatedAt
				versionNumber
			}
		}
		pageInfo {
			hasNextPage
			hasPreviousPage
			startCursor
			endCursor
		}
	}
}
`

func (c *Client) GetPlans(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetPlans, error) {
	vars := map[string]any{}

	var res GetPlans
	if err := c.Client.Post(ctx, "GetPlans", GetPlansDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	GetPlansDocument: "GetPlans",
}
