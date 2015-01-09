// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53domains provides a client for Amazon Route 53 Domains.
package route53domains

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// Route53Domains is a client for Amazon Route 53 Domains.
type Route53Domains struct {
	client *aws.JSONClient
}

// New returns a new Route53Domains client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *Route53Domains {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("route53domains", region)

	return &Route53Domains{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "Route53Domains_v20140515",
		},
	}
}

// CheckDomainAvailability this operation checks the availability of one
// domain name. You can access this API without authenticating. Note that
// if the availability status of a domain is pending, you must submit
// another request to determine the availability of the domain name.
func (c *Route53Domains) CheckDomainAvailability(req *CheckDomainAvailabilityRequest) (resp *CheckDomainAvailabilityResponse, err error) {
	resp = &CheckDomainAvailabilityResponse{}
	err = c.client.Do("CheckDomainAvailability", "POST", "/", req, resp)
	return
}

// DisableDomainAutoRenew this operation disables automatic renewal of
// domain registration for the specified domain.
func (c *Route53Domains) DisableDomainAutoRenew(req *DisableDomainAutoRenewRequest) (resp *DisableDomainAutoRenewResponse, err error) {
	resp = &DisableDomainAutoRenewResponse{}
	err = c.client.Do("DisableDomainAutoRenew", "POST", "/", req, resp)
	return
}

// DisableDomainTransferLock this operation removes the transfer lock on
// the domain (specifically the clientTransferProhibited status) to allow
// domain transfers. We recommend you refrain from performing this action
// unless you intend to transfer the domain to a different registrar.
// Successful submission returns an operation ID that you can use to track
// the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Route53Domains) DisableDomainTransferLock(req *DisableDomainTransferLockRequest) (resp *DisableDomainTransferLockResponse, err error) {
	resp = &DisableDomainTransferLockResponse{}
	err = c.client.Do("DisableDomainTransferLock", "POST", "/", req, resp)
	return
}

// EnableDomainAutoRenew this operation configures Amazon Route 53 to
// automatically renew the specified domain before the domain registration
// expires. The cost of renewing your domain registration is billed to your
// AWS account. The period during which you can renew a domain name varies
// by For a list of TLDs and their renewal policies, see "Renewal,
// restoration, and deletion times"
// (http://wiki.gandi.net/en/domains/renew#renewal_restoration_and_deletion_times)
// on the website for our registrar partner, Gandi. Route 53 requires that
// you renew before the end of the renewal period that is listed on the
// Gandi website so we can complete processing before the deadline.
func (c *Route53Domains) EnableDomainAutoRenew(req *EnableDomainAutoRenewRequest) (resp *EnableDomainAutoRenewResponse, err error) {
	resp = &EnableDomainAutoRenewResponse{}
	err = c.client.Do("EnableDomainAutoRenew", "POST", "/", req, resp)
	return
}

// EnableDomainTransferLock this operation sets the transfer lock on the
// domain (specifically the clientTransferProhibited status) to prevent
// domain transfers. Successful submission returns an operation ID that you
// can use to track the progress and completion of the action. If the
// request is not completed successfully, the domain registrant will be
// notified by email.
func (c *Route53Domains) EnableDomainTransferLock(req *EnableDomainTransferLockRequest) (resp *EnableDomainTransferLockResponse, err error) {
	resp = &EnableDomainTransferLockResponse{}
	err = c.client.Do("EnableDomainTransferLock", "POST", "/", req, resp)
	return
}

// GetDomainDetail this operation returns detailed information about the
// domain. The domain's contact information is also returned as part of the
// output.
func (c *Route53Domains) GetDomainDetail(req *GetDomainDetailRequest) (resp *GetDomainDetailResponse, err error) {
	resp = &GetDomainDetailResponse{}
	err = c.client.Do("GetDomainDetail", "POST", "/", req, resp)
	return
}

// GetOperationDetail this operation returns the current status of an
// operation that is not completed.
func (c *Route53Domains) GetOperationDetail(req *GetOperationDetailRequest) (resp *GetOperationDetailResponse, err error) {
	resp = &GetOperationDetailResponse{}
	err = c.client.Do("GetOperationDetail", "POST", "/", req, resp)
	return
}

// ListDomains this operation returns all the domain names registered with
// Amazon Route 53 for the current AWS account.
func (c *Route53Domains) ListDomains(req *ListDomainsRequest) (resp *ListDomainsResponse, err error) {
	resp = &ListDomainsResponse{}
	err = c.client.Do("ListDomains", "POST", "/", req, resp)
	return
}

// ListOperations this operation returns the operation IDs of operations
// that are not yet complete.
func (c *Route53Domains) ListOperations(req *ListOperationsRequest) (resp *ListOperationsResponse, err error) {
	resp = &ListOperationsResponse{}
	err = c.client.Do("ListOperations", "POST", "/", req, resp)
	return
}

// RegisterDomain this operation registers a domain. Domains are registered
// by the AWS registrar partner, Gandi. For some top-level domains (TLDs),
// this operation requires extra parameters. When you register a domain,
// Amazon Route 53 does the following: Creates a Amazon Route 53 hosted
// zone that has the same name as the domain. Amazon Route 53 assigns four
// name servers to your hosted zone and automatically updates your domain
// registration with the names of these name servers. Enables autorenew, so
// your domain registration will renew automatically each year. We'll
// notify you in advance of the renewal date so you can choose whether to
// renew the registration. Optionally enables privacy protection, so
// queries return contact information for our registrar partner, Gandi,
// instead of the information you entered for registrant, admin, and tech
// contacts. If registration is successful, returns an operation ID that
// you can use to track the progress and completion of the action. If the
// request is not completed successfully, the domain registrant is notified
// by email. Charges your AWS account an amount based on the top-level
// domain. For more information, see Amazon Route 53 Pricing
func (c *Route53Domains) RegisterDomain(req *RegisterDomainRequest) (resp *RegisterDomainResponse, err error) {
	resp = &RegisterDomainResponse{}
	err = c.client.Do("RegisterDomain", "POST", "/", req, resp)
	return
}

// RetrieveDomainAuthCode this operation returns the AuthCode for the
// domain. To transfer a domain to another registrar, you provide this
// value to the new registrar.
func (c *Route53Domains) RetrieveDomainAuthCode(req *RetrieveDomainAuthCodeRequest) (resp *RetrieveDomainAuthCodeResponse, err error) {
	resp = &RetrieveDomainAuthCodeResponse{}
	err = c.client.Do("RetrieveDomainAuthCode", "POST", "/", req, resp)
	return
}

// TransferDomain this operation transfers a domain from another registrar
// to Amazon Route 53. Domains are registered by the AWS registrar, Gandi
// upon transfer. To transfer a domain, you need to meet all the domain
// transfer criteria, including the following: You must supply nameservers
// to transfer a domain. You must disable the domain transfer lock (if any)
// before transferring the domain. A minimum of 60 days must have elapsed
// since the domain's registration or last transfer. We recommend you use
// the Amazon Route 53 as the DNS service for your domain. You can create a
// hosted zone in Amazon Route 53 for your current domain before
// transferring your domain. Note that upon transfer, the domain duration
// is extended for a year if not otherwise specified. Autorenew is enabled
// by default. If the transfer is successful, this method returns an
// operation ID that you can use to track the progress and completion of
// the action. If the request is not completed successfully, the domain
// registrant will be notified by email. Transferring domains charges your
// AWS account an amount based on the top-level domain. For more
// information, see Amazon Route 53 Pricing .
func (c *Route53Domains) TransferDomain(req *TransferDomainRequest) (resp *TransferDomainResponse, err error) {
	resp = &TransferDomainResponse{}
	err = c.client.Do("TransferDomain", "POST", "/", req, resp)
	return
}

// UpdateDomainContact this operation updates the contact information for a
// particular domain. Information for at least one contact (registrant,
// administrator, or technical) must be supplied for update. If the update
// is successful, this method returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainContact(req *UpdateDomainContactRequest) (resp *UpdateDomainContactResponse, err error) {
	resp = &UpdateDomainContactResponse{}
	err = c.client.Do("UpdateDomainContact", "POST", "/", req, resp)
	return
}

// UpdateDomainContactPrivacy this operation updates the specified domain
// contact's privacy setting. When the privacy option is enabled, personal
// information such as postal or email address is hidden from the results
// of a public query. The privacy services are provided by the AWS
// registrar, Gandi. For more information, see the Gandi privacy features
// This operation only affects the privacy of the specified contact type
// (registrant, administrator, or tech). Successful acceptance returns an
// operation ID that you can use with GetOperationDetail to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainContactPrivacy(req *UpdateDomainContactPrivacyRequest) (resp *UpdateDomainContactPrivacyResponse, err error) {
	resp = &UpdateDomainContactPrivacyResponse{}
	err = c.client.Do("UpdateDomainContactPrivacy", "POST", "/", req, resp)
	return
}

// UpdateDomainNameservers this operation replaces the current set of name
// servers for the domain with the specified set of name servers. If you
// use Amazon Route 53 as your DNS service, specify the four name servers
// in the delegation set for the hosted zone for the domain. If successful,
// this operation returns an operation ID that you can use to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
func (c *Route53Domains) UpdateDomainNameservers(req *UpdateDomainNameserversRequest) (resp *UpdateDomainNameserversResponse, err error) {
	resp = &UpdateDomainNameserversResponse{}
	err = c.client.Do("UpdateDomainNameservers", "POST", "/", req, resp)
	return
}

// CheckDomainAvailabilityRequest is undocumented.
type CheckDomainAvailabilityRequest struct {
	DomainName  aws.StringValue `json:"DomainName"`
	IDNLangCode aws.StringValue `json:"IdnLangCode,omitempty"`
}

// CheckDomainAvailabilityResponse is undocumented.
type CheckDomainAvailabilityResponse struct {
	Availability aws.StringValue `json:"Availability"`
}

// ContactDetail is undocumented.
type ContactDetail struct {
	AddressLine1     aws.StringValue `json:"AddressLine1,omitempty"`
	AddressLine2     aws.StringValue `json:"AddressLine2,omitempty"`
	City             aws.StringValue `json:"City,omitempty"`
	ContactType      aws.StringValue `json:"ContactType,omitempty"`
	CountryCode      aws.StringValue `json:"CountryCode,omitempty"`
	Email            aws.StringValue `json:"Email,omitempty"`
	ExtraParams      []ExtraParam    `json:"ExtraParams,omitempty"`
	Fax              aws.StringValue `json:"Fax,omitempty"`
	FirstName        aws.StringValue `json:"FirstName,omitempty"`
	LastName         aws.StringValue `json:"LastName,omitempty"`
	OrganizationName aws.StringValue `json:"OrganizationName,omitempty"`
	PhoneNumber      aws.StringValue `json:"PhoneNumber,omitempty"`
	State            aws.StringValue `json:"State,omitempty"`
	ZipCode          aws.StringValue `json:"ZipCode,omitempty"`
}

// Possible values for Route53Domains.
const (
	ContactTypeAssociation = "ASSOCIATION"
	ContactTypeCompany     = "COMPANY"
	ContactTypePerson      = "PERSON"
	ContactTypePublicBody  = "PUBLIC_BODY"
	ContactTypeReseller    = "RESELLER"
)

// Possible values for Route53Domains.
const (
	CountryCodeAd = "AD"
	CountryCodeAe = "AE"
	CountryCodeAf = "AF"
	CountryCodeAg = "AG"
	CountryCodeAi = "AI"
	CountryCodeAl = "AL"
	CountryCodeAm = "AM"
	CountryCodeAn = "AN"
	CountryCodeAo = "AO"
	CountryCodeAq = "AQ"
	CountryCodeAr = "AR"
	CountryCodeAs = "AS"
	CountryCodeAt = "AT"
	CountryCodeAu = "AU"
	CountryCodeAw = "AW"
	CountryCodeAz = "AZ"
	CountryCodeBa = "BA"
	CountryCodeBb = "BB"
	CountryCodeBd = "BD"
	CountryCodeBe = "BE"
	CountryCodeBf = "BF"
	CountryCodeBg = "BG"
	CountryCodeBh = "BH"
	CountryCodeBi = "BI"
	CountryCodeBj = "BJ"
	CountryCodeBl = "BL"
	CountryCodeBm = "BM"
	CountryCodeBn = "BN"
	CountryCodeBo = "BO"
	CountryCodeBr = "BR"
	CountryCodeBs = "BS"
	CountryCodeBt = "BT"
	CountryCodeBw = "BW"
	CountryCodeBy = "BY"
	CountryCodeBz = "BZ"
	CountryCodeCa = "CA"
	CountryCodeCc = "CC"
	CountryCodeCd = "CD"
	CountryCodeCf = "CF"
	CountryCodeCg = "CG"
	CountryCodeCh = "CH"
	CountryCodeCi = "CI"
	CountryCodeCk = "CK"
	CountryCodeCl = "CL"
	CountryCodeCm = "CM"
	CountryCodeCn = "CN"
	CountryCodeCo = "CO"
	CountryCodeCr = "CR"
	CountryCodeCu = "CU"
	CountryCodeCv = "CV"
	CountryCodeCx = "CX"
	CountryCodeCy = "CY"
	CountryCodeCz = "CZ"
	CountryCodeDe = "DE"
	CountryCodeDj = "DJ"
	CountryCodeDk = "DK"
	CountryCodeDm = "DM"
	CountryCodeDo = "DO"
	CountryCodeDz = "DZ"
	CountryCodeEc = "EC"
	CountryCodeEe = "EE"
	CountryCodeEg = "EG"
	CountryCodeEr = "ER"
	CountryCodeEs = "ES"
	CountryCodeEt = "ET"
	CountryCodeFi = "FI"
	CountryCodeFj = "FJ"
	CountryCodeFk = "FK"
	CountryCodeFm = "FM"
	CountryCodeFo = "FO"
	CountryCodeFr = "FR"
	CountryCodeGa = "GA"
	CountryCodeGb = "GB"
	CountryCodeGd = "GD"
	CountryCodeGe = "GE"
	CountryCodeGh = "GH"
	CountryCodeGi = "GI"
	CountryCodeGl = "GL"
	CountryCodeGm = "GM"
	CountryCodeGn = "GN"
	CountryCodeGq = "GQ"
	CountryCodeGr = "GR"
	CountryCodeGt = "GT"
	CountryCodeGu = "GU"
	CountryCodeGw = "GW"
	CountryCodeGy = "GY"
	CountryCodeHk = "HK"
	CountryCodeHn = "HN"
	CountryCodeHr = "HR"
	CountryCodeHt = "HT"
	CountryCodeHu = "HU"
	CountryCodeID = "ID"
	CountryCodeIe = "IE"
	CountryCodeIl = "IL"
	CountryCodeIm = "IM"
	CountryCodeIn = "IN"
	CountryCodeIq = "IQ"
	CountryCodeIr = "IR"
	CountryCodeIs = "IS"
	CountryCodeIt = "IT"
	CountryCodeJm = "JM"
	CountryCodeJo = "JO"
	CountryCodeJp = "JP"
	CountryCodeKe = "KE"
	CountryCodeKg = "KG"
	CountryCodeKh = "KH"
	CountryCodeKi = "KI"
	CountryCodeKm = "KM"
	CountryCodeKn = "KN"
	CountryCodeKp = "KP"
	CountryCodeKr = "KR"
	CountryCodeKw = "KW"
	CountryCodeKy = "KY"
	CountryCodeKz = "KZ"
	CountryCodeLa = "LA"
	CountryCodeLb = "LB"
	CountryCodeLc = "LC"
	CountryCodeLi = "LI"
	CountryCodeLk = "LK"
	CountryCodeLr = "LR"
	CountryCodeLs = "LS"
	CountryCodeLt = "LT"
	CountryCodeLu = "LU"
	CountryCodeLv = "LV"
	CountryCodeLy = "LY"
	CountryCodeMa = "MA"
	CountryCodeMc = "MC"
	CountryCodeMd = "MD"
	CountryCodeMe = "ME"
	CountryCodeMf = "MF"
	CountryCodeMg = "MG"
	CountryCodeMh = "MH"
	CountryCodeMk = "MK"
	CountryCodeMl = "ML"
	CountryCodeMm = "MM"
	CountryCodeMn = "MN"
	CountryCodeMo = "MO"
	CountryCodeMp = "MP"
	CountryCodeMr = "MR"
	CountryCodeMs = "MS"
	CountryCodeMt = "MT"
	CountryCodeMu = "MU"
	CountryCodeMv = "MV"
	CountryCodeMw = "MW"
	CountryCodeMx = "MX"
	CountryCodeMy = "MY"
	CountryCodeMz = "MZ"
	CountryCodeNa = "NA"
	CountryCodeNc = "NC"
	CountryCodeNe = "NE"
	CountryCodeNg = "NG"
	CountryCodeNi = "NI"
	CountryCodeNl = "NL"
	CountryCodeNo = "NO"
	CountryCodeNp = "NP"
	CountryCodeNr = "NR"
	CountryCodeNu = "NU"
	CountryCodeNz = "NZ"
	CountryCodeOm = "OM"
	CountryCodePa = "PA"
	CountryCodePe = "PE"
	CountryCodePf = "PF"
	CountryCodePg = "PG"
	CountryCodePh = "PH"
	CountryCodePk = "PK"
	CountryCodePl = "PL"
	CountryCodePm = "PM"
	CountryCodePn = "PN"
	CountryCodePr = "PR"
	CountryCodePt = "PT"
	CountryCodePw = "PW"
	CountryCodePy = "PY"
	CountryCodeQa = "QA"
	CountryCodeRo = "RO"
	CountryCodeRs = "RS"
	CountryCodeRu = "RU"
	CountryCodeRw = "RW"
	CountryCodeSa = "SA"
	CountryCodeSb = "SB"
	CountryCodeSc = "SC"
	CountryCodeSd = "SD"
	CountryCodeSe = "SE"
	CountryCodeSg = "SG"
	CountryCodeSh = "SH"
	CountryCodeSi = "SI"
	CountryCodeSk = "SK"
	CountryCodeSl = "SL"
	CountryCodeSm = "SM"
	CountryCodeSn = "SN"
	CountryCodeSo = "SO"
	CountryCodeSr = "SR"
	CountryCodeSt = "ST"
	CountryCodeSv = "SV"
	CountryCodeSy = "SY"
	CountryCodeSz = "SZ"
	CountryCodeTc = "TC"
	CountryCodeTd = "TD"
	CountryCodeTg = "TG"
	CountryCodeTh = "TH"
	CountryCodeTj = "TJ"
	CountryCodeTk = "TK"
	CountryCodeTl = "TL"
	CountryCodeTm = "TM"
	CountryCodeTn = "TN"
	CountryCodeTo = "TO"
	CountryCodeTr = "TR"
	CountryCodeTt = "TT"
	CountryCodeTv = "TV"
	CountryCodeTw = "TW"
	CountryCodeTz = "TZ"
	CountryCodeUa = "UA"
	CountryCodeUg = "UG"
	CountryCodeUs = "US"
	CountryCodeUy = "UY"
	CountryCodeUz = "UZ"
	CountryCodeVa = "VA"
	CountryCodeVc = "VC"
	CountryCodeVe = "VE"
	CountryCodeVg = "VG"
	CountryCodeVi = "VI"
	CountryCodeVn = "VN"
	CountryCodeVu = "VU"
	CountryCodeWf = "WF"
	CountryCodeWs = "WS"
	CountryCodeYe = "YE"
	CountryCodeYt = "YT"
	CountryCodeZa = "ZA"
	CountryCodeZm = "ZM"
	CountryCodeZw = "ZW"
)

// DisableDomainAutoRenewRequest is undocumented.
type DisableDomainAutoRenewRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// DisableDomainAutoRenewResponse is undocumented.
type DisableDomainAutoRenewResponse struct {
}

// DisableDomainTransferLockRequest is undocumented.
type DisableDomainTransferLockRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// DisableDomainTransferLockResponse is undocumented.
type DisableDomainTransferLockResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// Possible values for Route53Domains.
const (
	DomainAvailabilityAvailable             = "AVAILABLE"
	DomainAvailabilityAvailablePreorder     = "AVAILABLE_PREORDER"
	DomainAvailabilityAvailableReserved     = "AVAILABLE_RESERVED"
	DomainAvailabilityReserved              = "RESERVED"
	DomainAvailabilityUnavailable           = "UNAVAILABLE"
	DomainAvailabilityUnavailablePremium    = "UNAVAILABLE_PREMIUM"
	DomainAvailabilityUnavailableRestricted = "UNAVAILABLE_RESTRICTED"
)

// DomainSummary is undocumented.
type DomainSummary struct {
	AutoRenew    aws.BooleanValue `json:"AutoRenew,omitempty"`
	DomainName   aws.StringValue  `json:"DomainName"`
	Expiry       time.Time        `json:"Expiry,omitempty"`
	TransferLock aws.BooleanValue `json:"TransferLock,omitempty"`
}

// EnableDomainAutoRenewRequest is undocumented.
type EnableDomainAutoRenewRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// EnableDomainAutoRenewResponse is undocumented.
type EnableDomainAutoRenewResponse struct {
}

// EnableDomainTransferLockRequest is undocumented.
type EnableDomainTransferLockRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// EnableDomainTransferLockResponse is undocumented.
type EnableDomainTransferLockResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// ExtraParam is undocumented.
type ExtraParam struct {
	Name  aws.StringValue `json:"Name"`
	Value aws.StringValue `json:"Value"`
}

// Possible values for Route53Domains.
const (
	ExtraParamNameAuIDNumber          = "AU_ID_NUMBER"
	ExtraParamNameAuIDType            = "AU_ID_TYPE"
	ExtraParamNameBirthCity           = "BIRTH_CITY"
	ExtraParamNameBirthCountry        = "BIRTH_COUNTRY"
	ExtraParamNameBirthDateInYyyyMmDd = "BIRTH_DATE_IN_YYYY_MM_DD"
	ExtraParamNameBirthDepartment     = "BIRTH_DEPARTMENT"
	ExtraParamNameBrandNumber         = "BRAND_NUMBER"
	ExtraParamNameCaLegalType         = "CA_LEGAL_TYPE"
	ExtraParamNameDocumentNumber      = "DOCUMENT_NUMBER"
	ExtraParamNameDunsNumber          = "DUNS_NUMBER"
	ExtraParamNameFiBusinessNumber    = "FI_BUSINESS_NUMBER"
	ExtraParamNameFiIDNumber          = "FI_ID_NUMBER"
	ExtraParamNameItPin               = "IT_PIN"
	ExtraParamNameRuPassportData      = "RU_PASSPORT_DATA"
	ExtraParamNameSeIDNumber          = "SE_ID_NUMBER"
	ExtraParamNameSgIDNumber          = "SG_ID_NUMBER"
	ExtraParamNameVatNumber           = "VAT_NUMBER"
)

// GetDomainDetailRequest is undocumented.
type GetDomainDetailRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// GetDomainDetailResponse is undocumented.
type GetDomainDetailResponse struct {
	AbuseContactEmail aws.StringValue  `json:"AbuseContactEmail,omitempty"`
	AbuseContactPhone aws.StringValue  `json:"AbuseContactPhone,omitempty"`
	AdminContact      *ContactDetail   `json:"AdminContact"`
	AdminPrivacy      aws.BooleanValue `json:"AdminPrivacy,omitempty"`
	AutoRenew         aws.BooleanValue `json:"AutoRenew,omitempty"`
	CreationDate      time.Time        `json:"CreationDate,omitempty"`
	DNSSec            aws.StringValue  `json:"DnsSec,omitempty"`
	DomainName        aws.StringValue  `json:"DomainName"`
	ExpirationDate    time.Time        `json:"ExpirationDate,omitempty"`
	Nameservers       []Nameserver     `json:"Nameservers"`
	RegistrantContact *ContactDetail   `json:"RegistrantContact"`
	RegistrantPrivacy aws.BooleanValue `json:"RegistrantPrivacy,omitempty"`
	RegistrarName     aws.StringValue  `json:"RegistrarName,omitempty"`
	RegistrarURL      aws.StringValue  `json:"RegistrarUrl,omitempty"`
	RegistryDomainID  aws.StringValue  `json:"RegistryDomainId,omitempty"`
	Reseller          aws.StringValue  `json:"Reseller,omitempty"`
	StatusList        []string         `json:"StatusList,omitempty"`
	TechContact       *ContactDetail   `json:"TechContact"`
	TechPrivacy       aws.BooleanValue `json:"TechPrivacy,omitempty"`
	UpdatedDate       time.Time        `json:"UpdatedDate,omitempty"`
	WhoIsServer       aws.StringValue  `json:"WhoIsServer,omitempty"`
}

// GetOperationDetailRequest is undocumented.
type GetOperationDetailRequest struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// GetOperationDetailResponse is undocumented.
type GetOperationDetailResponse struct {
	DomainName    aws.StringValue `json:"DomainName,omitempty"`
	Message       aws.StringValue `json:"Message,omitempty"`
	OperationID   aws.StringValue `json:"OperationId,omitempty"`
	Status        aws.StringValue `json:"Status,omitempty"`
	SubmittedDate time.Time       `json:"SubmittedDate,omitempty"`
	Type          aws.StringValue `json:"Type,omitempty"`
}

// ListDomainsRequest is undocumented.
type ListDomainsRequest struct {
	Marker   aws.StringValue  `json:"Marker,omitempty"`
	MaxItems aws.IntegerValue `json:"MaxItems,omitempty"`
}

// ListDomainsResponse is undocumented.
type ListDomainsResponse struct {
	Domains        []DomainSummary `json:"Domains"`
	NextPageMarker aws.StringValue `json:"NextPageMarker,omitempty"`
}

// ListOperationsRequest is undocumented.
type ListOperationsRequest struct {
	Marker   aws.StringValue  `json:"Marker,omitempty"`
	MaxItems aws.IntegerValue `json:"MaxItems,omitempty"`
}

// ListOperationsResponse is undocumented.
type ListOperationsResponse struct {
	NextPageMarker aws.StringValue    `json:"NextPageMarker,omitempty"`
	Operations     []OperationSummary `json:"Operations"`
}

// Nameserver is undocumented.
type Nameserver struct {
	GlueIPs []string        `json:"GlueIps,omitempty"`
	Name    aws.StringValue `json:"Name"`
}

// Possible values for Route53Domains.
const (
	OperationStatusError      = "ERROR"
	OperationStatusFailed     = "FAILED"
	OperationStatusInProgress = "IN_PROGRESS"
	OperationStatusSubmitted  = "SUBMITTED"
	OperationStatusSuccessful = "SUCCESSFUL"
)

// OperationSummary is undocumented.
type OperationSummary struct {
	OperationID   aws.StringValue `json:"OperationId"`
	Status        aws.StringValue `json:"Status"`
	SubmittedDate time.Time       `json:"SubmittedDate"`
	Type          aws.StringValue `json:"Type"`
}

// Possible values for Route53Domains.
const (
	OperationTypeChangePrivacyProtection = "CHANGE_PRIVACY_PROTECTION"
	OperationTypeDeleteDomain            = "DELETE_DOMAIN"
	OperationTypeDomainLock              = "DOMAIN_LOCK"
	OperationTypeRegisterDomain          = "REGISTER_DOMAIN"
	OperationTypeTransferInDomain        = "TRANSFER_IN_DOMAIN"
	OperationTypeUpdateDomainContact     = "UPDATE_DOMAIN_CONTACT"
	OperationTypeUpdateNameserver        = "UPDATE_NAMESERVER"
)

// RegisterDomainRequest is undocumented.
type RegisterDomainRequest struct {
	AdminContact                    *ContactDetail   `json:"AdminContact"`
	AutoRenew                       aws.BooleanValue `json:"AutoRenew,omitempty"`
	DomainName                      aws.StringValue  `json:"DomainName"`
	DurationInYears                 aws.IntegerValue `json:"DurationInYears"`
	IDNLangCode                     aws.StringValue  `json:"IdnLangCode,omitempty"`
	PrivacyProtectAdminContact      aws.BooleanValue `json:"PrivacyProtectAdminContact,omitempty"`
	PrivacyProtectRegistrantContact aws.BooleanValue `json:"PrivacyProtectRegistrantContact,omitempty"`
	PrivacyProtectTechContact       aws.BooleanValue `json:"PrivacyProtectTechContact,omitempty"`
	RegistrantContact               *ContactDetail   `json:"RegistrantContact"`
	TechContact                     *ContactDetail   `json:"TechContact"`
}

// RegisterDomainResponse is undocumented.
type RegisterDomainResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// RetrieveDomainAuthCodeRequest is undocumented.
type RetrieveDomainAuthCodeRequest struct {
	DomainName aws.StringValue `json:"DomainName"`
}

// RetrieveDomainAuthCodeResponse is undocumented.
type RetrieveDomainAuthCodeResponse struct {
	AuthCode aws.StringValue `json:"AuthCode"`
}

// TransferDomainRequest is undocumented.
type TransferDomainRequest struct {
	AdminContact                    *ContactDetail   `json:"AdminContact"`
	AuthCode                        aws.StringValue  `json:"AuthCode,omitempty"`
	AutoRenew                       aws.BooleanValue `json:"AutoRenew,omitempty"`
	DomainName                      aws.StringValue  `json:"DomainName"`
	DurationInYears                 aws.IntegerValue `json:"DurationInYears"`
	IDNLangCode                     aws.StringValue  `json:"IdnLangCode,omitempty"`
	Nameservers                     []Nameserver     `json:"Nameservers"`
	PrivacyProtectAdminContact      aws.BooleanValue `json:"PrivacyProtectAdminContact,omitempty"`
	PrivacyProtectRegistrantContact aws.BooleanValue `json:"PrivacyProtectRegistrantContact,omitempty"`
	PrivacyProtectTechContact       aws.BooleanValue `json:"PrivacyProtectTechContact,omitempty"`
	RegistrantContact               *ContactDetail   `json:"RegistrantContact"`
	TechContact                     *ContactDetail   `json:"TechContact"`
}

// TransferDomainResponse is undocumented.
type TransferDomainResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// UpdateDomainContactPrivacyRequest is undocumented.
type UpdateDomainContactPrivacyRequest struct {
	AdminPrivacy      aws.BooleanValue `json:"AdminPrivacy,omitempty"`
	DomainName        aws.StringValue  `json:"DomainName"`
	RegistrantPrivacy aws.BooleanValue `json:"RegistrantPrivacy,omitempty"`
	TechPrivacy       aws.BooleanValue `json:"TechPrivacy,omitempty"`
}

// UpdateDomainContactPrivacyResponse is undocumented.
type UpdateDomainContactPrivacyResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// UpdateDomainContactRequest is undocumented.
type UpdateDomainContactRequest struct {
	AdminContact      *ContactDetail  `json:"AdminContact,omitempty"`
	DomainName        aws.StringValue `json:"DomainName"`
	RegistrantContact *ContactDetail  `json:"RegistrantContact,omitempty"`
	TechContact       *ContactDetail  `json:"TechContact,omitempty"`
}

// UpdateDomainContactResponse is undocumented.
type UpdateDomainContactResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// UpdateDomainNameserversRequest is undocumented.
type UpdateDomainNameserversRequest struct {
	DomainName  aws.StringValue `json:"DomainName"`
	Nameservers []Nameserver    `json:"Nameservers"`
}

// UpdateDomainNameserversResponse is undocumented.
type UpdateDomainNameserversResponse struct {
	OperationID aws.StringValue `json:"OperationId"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
