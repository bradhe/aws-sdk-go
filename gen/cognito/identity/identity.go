// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentity provides a client for Amazon Cognito Identity.
package cognitoidentity

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// CognitoIdentity is a client for Amazon Cognito Identity.
type CognitoIdentity struct {
	client *aws.JSONClient
}

// New returns a new CognitoIdentity client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *CognitoIdentity {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("cognito-identity", region)

	return &CognitoIdentity{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "AWSCognitoIdentityService",
		},
	}
}

// CreateIdentityPool creates a new identity pool. The identity pool is a
// store of user identity information that is specific to your AWS account.
// The limit on identity pools is 60 per account.
func (c *CognitoIdentity) CreateIdentityPool(req *CreateIdentityPoolInput) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("CreateIdentityPool", "POST", "/", req, resp)
	return
}

// DeleteIdentityPool deletes a user pool. Once a pool is deleted, users
// will not be able to authenticate with the pool.
func (c *CognitoIdentity) DeleteIdentityPool(req *DeleteIdentityPoolInput) (err error) {
	// NRE
	err = c.client.Do("DeleteIdentityPool", "POST", "/", req, nil)
	return
}

// DescribeIdentityPool gets details about a particular identity pool,
// including the pool name, ID description, creation date, and current
// number of users.
func (c *CognitoIdentity) DescribeIdentityPool(req *DescribeIdentityPoolInput) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("DescribeIdentityPool", "POST", "/", req, resp)
	return
}

// GetID generates (or retrieves) a Cognito ID. Supplying multiple logins
// will create an implicit linked account.
func (c *CognitoIdentity) GetID(req *GetIDInput) (resp *GetIDResponse, err error) {
	resp = &GetIDResponse{}
	err = c.client.Do("GetId", "POST", "/", req, resp)
	return
}

// GetOpenIDToken gets an OpenID token, using a known Cognito ID. This
// known Cognito ID is returned by GetId . You can optionally add
// additional logins for the identity. Supplying multiple logins creates an
// implicit link. The OpenId token is valid for 15 minutes.
func (c *CognitoIdentity) GetOpenIDToken(req *GetOpenIDTokenInput) (resp *GetOpenIDTokenResponse, err error) {
	resp = &GetOpenIDTokenResponse{}
	err = c.client.Do("GetOpenIdToken", "POST", "/", req, resp)
	return
}

// GetOpenIDTokenForDeveloperIdentity registers (or retrieves) a Cognito
// IdentityId and an OpenID Connect token for a user authenticated by your
// backend authentication process. Supplying multiple logins will create an
// implicit linked account. You can only specify one developer provider as
// part of the Logins map, which is linked to the identity pool. The
// developer provider is the "domain" by which Cognito will refer to your
// users. You can use GetOpenIdTokenForDeveloperIdentity to create a new
// identity and to link new logins (that is, user credentials issued by a
// public provider or developer provider) to an existing identity. When you
// want to create a new identity, the IdentityId should be null. When you
// want to associate a new login with an existing
// authenticated/unauthenticated identity, you can do so by providing the
// existing IdentityId . This API will create the identity in the specified
// IdentityPoolId
func (c *CognitoIdentity) GetOpenIDTokenForDeveloperIdentity(req *GetOpenIDTokenForDeveloperIdentityInput) (resp *GetOpenIDTokenForDeveloperIdentityResponse, err error) {
	resp = &GetOpenIDTokenForDeveloperIdentityResponse{}
	err = c.client.Do("GetOpenIdTokenForDeveloperIdentity", "POST", "/", req, resp)
	return
}

// ListIdentities is undocumented.
func (c *CognitoIdentity) ListIdentities(req *ListIdentitiesInput) (resp *ListIdentitiesResponse, err error) {
	resp = &ListIdentitiesResponse{}
	err = c.client.Do("ListIdentities", "POST", "/", req, resp)
	return
}

// ListIdentityPools lists all of the Cognito identity pools registered for
// your account.
func (c *CognitoIdentity) ListIdentityPools(req *ListIdentityPoolsInput) (resp *ListIdentityPoolsResponse, err error) {
	resp = &ListIdentityPoolsResponse{}
	err = c.client.Do("ListIdentityPools", "POST", "/", req, resp)
	return
}

// LookupDeveloperIdentity retrieves the IdentityID associated with a
// DeveloperUserIdentifier or the list of DeveloperUserIdentifier s
// associated with an IdentityId for an existing identity. Either
// IdentityID or DeveloperUserIdentifier must not be null. If you supply
// only one of these values, the other value will be searched in the
// database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID . If the
// values are verified against the database, the response returns both
// values and is the same as the request. Otherwise a
// ResourceConflictException is thrown.
func (c *CognitoIdentity) LookupDeveloperIdentity(req *LookupDeveloperIdentityInput) (resp *LookupDeveloperIdentityResponse, err error) {
	resp = &LookupDeveloperIdentityResponse{}
	err = c.client.Do("LookupDeveloperIdentity", "POST", "/", req, resp)
	return
}

// MergeDeveloperIdentities merges two users having different IdentityId s,
// existing in the same identity pool, and identified by the same developer
// provider. You can use this action to request that discrete users be
// merged and identified as a single user in the Cognito environment.
// Cognito associates the given source user SourceUserIdentifier ) with the
// IdentityId of the DestinationUserIdentifier . Only
// developer-authenticated users can be merged. If the users to be merged
// are associated with the same public provider, but as two different
// users, an exception will be thrown.
func (c *CognitoIdentity) MergeDeveloperIdentities(req *MergeDeveloperIdentitiesInput) (resp *MergeDeveloperIdentitiesResponse, err error) {
	resp = &MergeDeveloperIdentitiesResponse{}
	err = c.client.Do("MergeDeveloperIdentities", "POST", "/", req, resp)
	return
}

// UnlinkDeveloperIdentity unlinks a DeveloperUserIdentifier from an
// existing identity. Unlinked developer users will be considered new
// identities next time they are seen. If, for a given Cognito identity,
// you remove all federated identities as well as the developer user
// identifier, the Cognito identity becomes inaccessible.
func (c *CognitoIdentity) UnlinkDeveloperIdentity(req *UnlinkDeveloperIdentityInput) (err error) {
	// NRE
	err = c.client.Do("UnlinkDeveloperIdentity", "POST", "/", req, nil)
	return
}

// UnlinkIdentity unlinks a federated identity from an existing account.
// Unlinked logins will be considered new identities next time they are
// seen. Removing the last linked login will make this identity
// inaccessible.
func (c *CognitoIdentity) UnlinkIdentity(req *UnlinkIdentityInput) (err error) {
	// NRE
	err = c.client.Do("UnlinkIdentity", "POST", "/", req, nil)
	return
}

// UpdateIdentityPool is undocumented.
func (c *CognitoIdentity) UpdateIdentityPool(req *IdentityPool) (resp *IdentityPool, err error) {
	resp = &IdentityPool{}
	err = c.client.Do("UpdateIdentityPool", "POST", "/", req, resp)
	return
}

// CreateIdentityPoolInput is undocumented.
type CreateIdentityPoolInput struct {
	AllowUnauthenticatedIdentities aws.BooleanValue  `json:"AllowUnauthenticatedIdentities"`
	DeveloperProviderName          aws.StringValue   `json:"DeveloperProviderName,omitempty"`
	IdentityPoolName               aws.StringValue   `json:"IdentityPoolName"`
	OpenIDConnectProviderARNs      []string          `json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// DeleteIdentityPoolInput is undocumented.
type DeleteIdentityPoolInput struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// DescribeIdentityPoolInput is undocumented.
type DescribeIdentityPoolInput struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// GetIDInput is undocumented.
type GetIDInput struct {
	AccountID      aws.StringValue   `json:"AccountId"`
	IdentityPoolID aws.StringValue   `json:"IdentityPoolId"`
	Logins         map[string]string `json:"Logins,omitempty"`
}

// GetIDResponse is undocumented.
type GetIDResponse struct {
	IdentityID aws.StringValue `json:"IdentityId,omitempty"`
}

// GetOpenIDTokenForDeveloperIdentityInput is undocumented.
type GetOpenIDTokenForDeveloperIdentityInput struct {
	IdentityID     aws.StringValue   `json:"IdentityId,omitempty"`
	IdentityPoolID aws.StringValue   `json:"IdentityPoolId"`
	Logins         map[string]string `json:"Logins"`
	TokenDuration  aws.LongValue     `json:"TokenDuration,omitempty"`
}

// GetOpenIDTokenForDeveloperIdentityResponse is undocumented.
type GetOpenIDTokenForDeveloperIdentityResponse struct {
	IdentityID aws.StringValue `json:"IdentityId,omitempty"`
	Token      aws.StringValue `json:"Token,omitempty"`
}

// GetOpenIDTokenInput is undocumented.
type GetOpenIDTokenInput struct {
	IdentityID aws.StringValue   `json:"IdentityId"`
	Logins     map[string]string `json:"Logins,omitempty"`
}

// GetOpenIDTokenResponse is undocumented.
type GetOpenIDTokenResponse struct {
	IdentityID aws.StringValue `json:"IdentityId,omitempty"`
	Token      aws.StringValue `json:"Token,omitempty"`
}

// IdentityDescription is undocumented.
type IdentityDescription struct {
	IdentityID aws.StringValue `json:"IdentityId,omitempty"`
	Logins     []string        `json:"Logins,omitempty"`
}

// IdentityPool is undocumented.
type IdentityPool struct {
	AllowUnauthenticatedIdentities aws.BooleanValue  `json:"AllowUnauthenticatedIdentities"`
	DeveloperProviderName          aws.StringValue   `json:"DeveloperProviderName,omitempty"`
	IdentityPoolID                 aws.StringValue   `json:"IdentityPoolId"`
	IdentityPoolName               aws.StringValue   `json:"IdentityPoolName"`
	OpenIDConnectProviderARNs      []string          `json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// IdentityPoolShortDescription is undocumented.
type IdentityPoolShortDescription struct {
	IdentityPoolID   aws.StringValue `json:"IdentityPoolId,omitempty"`
	IdentityPoolName aws.StringValue `json:"IdentityPoolName,omitempty"`
}

// ListIdentitiesInput is undocumented.
type ListIdentitiesInput struct {
	IdentityPoolID aws.StringValue  `json:"IdentityPoolId"`
	MaxResults     aws.IntegerValue `json:"MaxResults"`
	NextToken      aws.StringValue  `json:"NextToken,omitempty"`
}

// ListIdentitiesResponse is undocumented.
type ListIdentitiesResponse struct {
	Identities     []IdentityDescription `json:"Identities,omitempty"`
	IdentityPoolID aws.StringValue       `json:"IdentityPoolId,omitempty"`
	NextToken      aws.StringValue       `json:"NextToken,omitempty"`
}

// ListIdentityPoolsInput is undocumented.
type ListIdentityPoolsInput struct {
	MaxResults aws.IntegerValue `json:"MaxResults"`
	NextToken  aws.StringValue  `json:"NextToken,omitempty"`
}

// ListIdentityPoolsResponse is undocumented.
type ListIdentityPoolsResponse struct {
	IdentityPools []IdentityPoolShortDescription `json:"IdentityPools,omitempty"`
	NextToken     aws.StringValue                `json:"NextToken,omitempty"`
}

// LookupDeveloperIdentityInput is undocumented.
type LookupDeveloperIdentityInput struct {
	DeveloperUserIdentifier aws.StringValue  `json:"DeveloperUserIdentifier,omitempty"`
	IdentityID              aws.StringValue  `json:"IdentityId,omitempty"`
	IdentityPoolID          aws.StringValue  `json:"IdentityPoolId"`
	MaxResults              aws.IntegerValue `json:"MaxResults,omitempty"`
	NextToken               aws.StringValue  `json:"NextToken,omitempty"`
}

// LookupDeveloperIdentityResponse is undocumented.
type LookupDeveloperIdentityResponse struct {
	DeveloperUserIdentifierList []string        `json:"DeveloperUserIdentifierList,omitempty"`
	IdentityID                  aws.StringValue `json:"IdentityId,omitempty"`
	NextToken                   aws.StringValue `json:"NextToken,omitempty"`
}

// MergeDeveloperIdentitiesInput is undocumented.
type MergeDeveloperIdentitiesInput struct {
	DestinationUserIdentifier aws.StringValue `json:"DestinationUserIdentifier"`
	DeveloperProviderName     aws.StringValue `json:"DeveloperProviderName"`
	IdentityPoolID            aws.StringValue `json:"IdentityPoolId"`
	SourceUserIdentifier      aws.StringValue `json:"SourceUserIdentifier"`
}

// MergeDeveloperIdentitiesResponse is undocumented.
type MergeDeveloperIdentitiesResponse struct {
	IdentityID aws.StringValue `json:"IdentityId,omitempty"`
}

// UnlinkDeveloperIdentityInput is undocumented.
type UnlinkDeveloperIdentityInput struct {
	DeveloperProviderName   aws.StringValue `json:"DeveloperProviderName"`
	DeveloperUserIdentifier aws.StringValue `json:"DeveloperUserIdentifier"`
	IdentityID              aws.StringValue `json:"IdentityId"`
	IdentityPoolID          aws.StringValue `json:"IdentityPoolId"`
}

// UnlinkIdentityInput is undocumented.
type UnlinkIdentityInput struct {
	IdentityID     aws.StringValue   `json:"IdentityId"`
	Logins         map[string]string `json:"Logins"`
	LoginsToRemove []string          `json:"LoginsToRemove"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
