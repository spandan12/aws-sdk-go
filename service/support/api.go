package support

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AddAttachmentsToSetRequest generates a request for the AddAttachmentsToSet operation.
func (c *Support) AddAttachmentsToSetRequest(input *AddAttachmentsToSetInput) (req *aws.Request, output *AddAttachmentsToSetOutput) {
	if opAddAttachmentsToSet == nil {
		opAddAttachmentsToSet = &aws.Operation{
			Name:       "AddAttachmentsToSet",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAddAttachmentsToSet, input, output)
	output = &AddAttachmentsToSetOutput{}
	req.Data = output
	return
}

func (c *Support) AddAttachmentsToSet(input *AddAttachmentsToSetInput) (output *AddAttachmentsToSetOutput, err error) {
	req, out := c.AddAttachmentsToSetRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddAttachmentsToSet *aws.Operation

// AddCommunicationToCaseRequest generates a request for the AddCommunicationToCase operation.
func (c *Support) AddCommunicationToCaseRequest(input *AddCommunicationToCaseInput) (req *aws.Request, output *AddCommunicationToCaseOutput) {
	if opAddCommunicationToCase == nil {
		opAddCommunicationToCase = &aws.Operation{
			Name:       "AddCommunicationToCase",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAddCommunicationToCase, input, output)
	output = &AddCommunicationToCaseOutput{}
	req.Data = output
	return
}

func (c *Support) AddCommunicationToCase(input *AddCommunicationToCaseInput) (output *AddCommunicationToCaseOutput, err error) {
	req, out := c.AddCommunicationToCaseRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddCommunicationToCase *aws.Operation

// CreateCaseRequest generates a request for the CreateCase operation.
func (c *Support) CreateCaseRequest(input *CreateCaseInput) (req *aws.Request, output *CreateCaseOutput) {
	if opCreateCase == nil {
		opCreateCase = &aws.Operation{
			Name:       "CreateCase",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateCase, input, output)
	output = &CreateCaseOutput{}
	req.Data = output
	return
}

func (c *Support) CreateCase(input *CreateCaseInput) (output *CreateCaseOutput, err error) {
	req, out := c.CreateCaseRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateCase *aws.Operation

// DescribeAttachmentRequest generates a request for the DescribeAttachment operation.
func (c *Support) DescribeAttachmentRequest(input *DescribeAttachmentInput) (req *aws.Request, output *DescribeAttachmentOutput) {
	if opDescribeAttachment == nil {
		opDescribeAttachment = &aws.Operation{
			Name:       "DescribeAttachment",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAttachment, input, output)
	output = &DescribeAttachmentOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeAttachment(input *DescribeAttachmentInput) (output *DescribeAttachmentOutput, err error) {
	req, out := c.DescribeAttachmentRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAttachment *aws.Operation

// DescribeCasesRequest generates a request for the DescribeCases operation.
func (c *Support) DescribeCasesRequest(input *DescribeCasesInput) (req *aws.Request, output *DescribeCasesOutput) {
	if opDescribeCases == nil {
		opDescribeCases = &aws.Operation{
			Name:       "DescribeCases",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeCases, input, output)
	output = &DescribeCasesOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeCases(input *DescribeCasesInput) (output *DescribeCasesOutput, err error) {
	req, out := c.DescribeCasesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeCases *aws.Operation

// DescribeCommunicationsRequest generates a request for the DescribeCommunications operation.
func (c *Support) DescribeCommunicationsRequest(input *DescribeCommunicationsInput) (req *aws.Request, output *DescribeCommunicationsOutput) {
	if opDescribeCommunications == nil {
		opDescribeCommunications = &aws.Operation{
			Name:       "DescribeCommunications",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeCommunications, input, output)
	output = &DescribeCommunicationsOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeCommunications(input *DescribeCommunicationsInput) (output *DescribeCommunicationsOutput, err error) {
	req, out := c.DescribeCommunicationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeCommunications *aws.Operation

// DescribeServicesRequest generates a request for the DescribeServices operation.
func (c *Support) DescribeServicesRequest(input *DescribeServicesInput) (req *aws.Request, output *DescribeServicesOutput) {
	if opDescribeServices == nil {
		opDescribeServices = &aws.Operation{
			Name:       "DescribeServices",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeServices, input, output)
	output = &DescribeServicesOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeServices(input *DescribeServicesInput) (output *DescribeServicesOutput, err error) {
	req, out := c.DescribeServicesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeServices *aws.Operation

// DescribeSeverityLevelsRequest generates a request for the DescribeSeverityLevels operation.
func (c *Support) DescribeSeverityLevelsRequest(input *DescribeSeverityLevelsInput) (req *aws.Request, output *DescribeSeverityLevelsOutput) {
	if opDescribeSeverityLevels == nil {
		opDescribeSeverityLevels = &aws.Operation{
			Name:       "DescribeSeverityLevels",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeSeverityLevels, input, output)
	output = &DescribeSeverityLevelsOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeSeverityLevels(input *DescribeSeverityLevelsInput) (output *DescribeSeverityLevelsOutput, err error) {
	req, out := c.DescribeSeverityLevelsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeSeverityLevels *aws.Operation

// DescribeTrustedAdvisorCheckRefreshStatusesRequest generates a request for the DescribeTrustedAdvisorCheckRefreshStatuses operation.
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatusesRequest(input *DescribeTrustedAdvisorCheckRefreshStatusesInput) (req *aws.Request, output *DescribeTrustedAdvisorCheckRefreshStatusesOutput) {
	if opDescribeTrustedAdvisorCheckRefreshStatuses == nil {
		opDescribeTrustedAdvisorCheckRefreshStatuses = &aws.Operation{
			Name:       "DescribeTrustedAdvisorCheckRefreshStatuses",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTrustedAdvisorCheckRefreshStatuses, input, output)
	output = &DescribeTrustedAdvisorCheckRefreshStatusesOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeTrustedAdvisorCheckRefreshStatuses(input *DescribeTrustedAdvisorCheckRefreshStatusesInput) (output *DescribeTrustedAdvisorCheckRefreshStatusesOutput, err error) {
	req, out := c.DescribeTrustedAdvisorCheckRefreshStatusesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrustedAdvisorCheckRefreshStatuses *aws.Operation

// DescribeTrustedAdvisorCheckResultRequest generates a request for the DescribeTrustedAdvisorCheckResult operation.
func (c *Support) DescribeTrustedAdvisorCheckResultRequest(input *DescribeTrustedAdvisorCheckResultInput) (req *aws.Request, output *DescribeTrustedAdvisorCheckResultOutput) {
	if opDescribeTrustedAdvisorCheckResult == nil {
		opDescribeTrustedAdvisorCheckResult = &aws.Operation{
			Name:       "DescribeTrustedAdvisorCheckResult",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTrustedAdvisorCheckResult, input, output)
	output = &DescribeTrustedAdvisorCheckResultOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeTrustedAdvisorCheckResult(input *DescribeTrustedAdvisorCheckResultInput) (output *DescribeTrustedAdvisorCheckResultOutput, err error) {
	req, out := c.DescribeTrustedAdvisorCheckResultRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrustedAdvisorCheckResult *aws.Operation

// DescribeTrustedAdvisorCheckSummariesRequest generates a request for the DescribeTrustedAdvisorCheckSummaries operation.
func (c *Support) DescribeTrustedAdvisorCheckSummariesRequest(input *DescribeTrustedAdvisorCheckSummariesInput) (req *aws.Request, output *DescribeTrustedAdvisorCheckSummariesOutput) {
	if opDescribeTrustedAdvisorCheckSummaries == nil {
		opDescribeTrustedAdvisorCheckSummaries = &aws.Operation{
			Name:       "DescribeTrustedAdvisorCheckSummaries",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTrustedAdvisorCheckSummaries, input, output)
	output = &DescribeTrustedAdvisorCheckSummariesOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeTrustedAdvisorCheckSummaries(input *DescribeTrustedAdvisorCheckSummariesInput) (output *DescribeTrustedAdvisorCheckSummariesOutput, err error) {
	req, out := c.DescribeTrustedAdvisorCheckSummariesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrustedAdvisorCheckSummaries *aws.Operation

// DescribeTrustedAdvisorChecksRequest generates a request for the DescribeTrustedAdvisorChecks operation.
func (c *Support) DescribeTrustedAdvisorChecksRequest(input *DescribeTrustedAdvisorChecksInput) (req *aws.Request, output *DescribeTrustedAdvisorChecksOutput) {
	if opDescribeTrustedAdvisorChecks == nil {
		opDescribeTrustedAdvisorChecks = &aws.Operation{
			Name:       "DescribeTrustedAdvisorChecks",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTrustedAdvisorChecks, input, output)
	output = &DescribeTrustedAdvisorChecksOutput{}
	req.Data = output
	return
}

func (c *Support) DescribeTrustedAdvisorChecks(input *DescribeTrustedAdvisorChecksInput) (output *DescribeTrustedAdvisorChecksOutput, err error) {
	req, out := c.DescribeTrustedAdvisorChecksRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrustedAdvisorChecks *aws.Operation

// RefreshTrustedAdvisorCheckRequest generates a request for the RefreshTrustedAdvisorCheck operation.
func (c *Support) RefreshTrustedAdvisorCheckRequest(input *RefreshTrustedAdvisorCheckInput) (req *aws.Request, output *RefreshTrustedAdvisorCheckOutput) {
	if opRefreshTrustedAdvisorCheck == nil {
		opRefreshTrustedAdvisorCheck = &aws.Operation{
			Name:       "RefreshTrustedAdvisorCheck",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRefreshTrustedAdvisorCheck, input, output)
	output = &RefreshTrustedAdvisorCheckOutput{}
	req.Data = output
	return
}

func (c *Support) RefreshTrustedAdvisorCheck(input *RefreshTrustedAdvisorCheckInput) (output *RefreshTrustedAdvisorCheckOutput, err error) {
	req, out := c.RefreshTrustedAdvisorCheckRequest(input)
	output = out
	err = req.Send()
	return
}

var opRefreshTrustedAdvisorCheck *aws.Operation

// ResolveCaseRequest generates a request for the ResolveCase operation.
func (c *Support) ResolveCaseRequest(input *ResolveCaseInput) (req *aws.Request, output *ResolveCaseOutput) {
	if opResolveCase == nil {
		opResolveCase = &aws.Operation{
			Name:       "ResolveCase",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opResolveCase, input, output)
	output = &ResolveCaseOutput{}
	req.Data = output
	return
}

func (c *Support) ResolveCase(input *ResolveCaseInput) (output *ResolveCaseOutput, err error) {
	req, out := c.ResolveCaseRequest(input)
	output = out
	err = req.Send()
	return
}

var opResolveCase *aws.Operation

type AddAttachmentsToSetInput struct {
	AttachmentSetID *string       `locationName:"attachmentSetId" type:"string" json:",omitempty"`
	Attachments     []*Attachment `locationName:"attachments" type:"list" json:",omitempty"`

	metadataAddAttachmentsToSetInput `json:"-", xml:"-"`
}

type metadataAddAttachmentsToSetInput struct {
	SDKShapeTraits bool `type:"structure" required:"attachments" json:",omitempty"`
}

type AddAttachmentsToSetOutput struct {
	AttachmentSetID *string `locationName:"attachmentSetId" type:"string" json:",omitempty"`
	ExpiryTime      *string `locationName:"expiryTime" type:"string" json:",omitempty"`

	metadataAddAttachmentsToSetOutput `json:"-", xml:"-"`
}

type metadataAddAttachmentsToSetOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AddCommunicationToCaseInput struct {
	AttachmentSetID   *string   `locationName:"attachmentSetId" type:"string" json:",omitempty"`
	CCEmailAddresses  []*string `locationName:"ccEmailAddresses" type:"list" json:",omitempty"`
	CaseID            *string   `locationName:"caseId" type:"string" json:",omitempty"`
	CommunicationBody *string   `locationName:"communicationBody" type:"string" json:",omitempty"`

	metadataAddCommunicationToCaseInput `json:"-", xml:"-"`
}

type metadataAddCommunicationToCaseInput struct {
	SDKShapeTraits bool `type:"structure" required:"communicationBody" json:",omitempty"`
}

type AddCommunicationToCaseOutput struct {
	Result *bool `locationName:"result" type:"boolean" json:",omitempty"`

	metadataAddCommunicationToCaseOutput `json:"-", xml:"-"`
}

type metadataAddCommunicationToCaseOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Attachment struct {
	Data     []byte  `locationName:"data" type:"blob" json:",omitempty"`
	FileName *string `locationName:"fileName" type:"string" json:",omitempty"`

	metadataAttachment `json:"-", xml:"-"`
}

type metadataAttachment struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentDetails struct {
	AttachmentID *string `locationName:"attachmentId" type:"string" json:",omitempty"`
	FileName     *string `locationName:"fileName" type:"string" json:",omitempty"`

	metadataAttachmentDetails `json:"-", xml:"-"`
}

type metadataAttachmentDetails struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentIDNotFound struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataAttachmentIDNotFound `json:"-", xml:"-"`
}

type metadataAttachmentIDNotFound struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataAttachmentLimitExceeded `json:"-", xml:"-"`
}

type metadataAttachmentLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentSetExpired struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataAttachmentSetExpired `json:"-", xml:"-"`
}

type metadataAttachmentSetExpired struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentSetIDNotFound struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataAttachmentSetIDNotFound `json:"-", xml:"-"`
}

type metadataAttachmentSetIDNotFound struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type AttachmentSetSizeLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataAttachmentSetSizeLimitExceeded `json:"-", xml:"-"`
}

type metadataAttachmentSetSizeLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CaseCreationLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataCaseCreationLimitExceeded `json:"-", xml:"-"`
}

type metadataCaseCreationLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CaseDetails struct {
	CCEmailAddresses     []*string                 `locationName:"ccEmailAddresses" type:"list" json:",omitempty"`
	CaseID               *string                   `locationName:"caseId" type:"string" json:",omitempty"`
	CategoryCode         *string                   `locationName:"categoryCode" type:"string" json:",omitempty"`
	DisplayID            *string                   `locationName:"displayId" type:"string" json:",omitempty"`
	Language             *string                   `locationName:"language" type:"string" json:",omitempty"`
	RecentCommunications *RecentCaseCommunications `locationName:"recentCommunications" type:"structure" json:",omitempty"`
	ServiceCode          *string                   `locationName:"serviceCode" type:"string" json:",omitempty"`
	SeverityCode         *string                   `locationName:"severityCode" type:"string" json:",omitempty"`
	Status               *string                   `locationName:"status" type:"string" json:",omitempty"`
	Subject              *string                   `locationName:"subject" type:"string" json:",omitempty"`
	SubmittedBy          *string                   `locationName:"submittedBy" type:"string" json:",omitempty"`
	TimeCreated          *string                   `locationName:"timeCreated" type:"string" json:",omitempty"`

	metadataCaseDetails `json:"-", xml:"-"`
}

type metadataCaseDetails struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CaseIDNotFound struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataCaseIDNotFound `json:"-", xml:"-"`
}

type metadataCaseIDNotFound struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Category struct {
	Code *string `locationName:"code" type:"string" json:",omitempty"`
	Name *string `locationName:"name" type:"string" json:",omitempty"`

	metadataCategory `json:"-", xml:"-"`
}

type metadataCategory struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Communication struct {
	AttachmentSet []*AttachmentDetails `locationName:"attachmentSet" type:"list" json:",omitempty"`
	Body          *string              `locationName:"body" type:"string" json:",omitempty"`
	CaseID        *string              `locationName:"caseId" type:"string" json:",omitempty"`
	SubmittedBy   *string              `locationName:"submittedBy" type:"string" json:",omitempty"`
	TimeCreated   *string              `locationName:"timeCreated" type:"string" json:",omitempty"`

	metadataCommunication `json:"-", xml:"-"`
}

type metadataCommunication struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreateCaseInput struct {
	AttachmentSetID   *string   `locationName:"attachmentSetId" type:"string" json:",omitempty"`
	CCEmailAddresses  []*string `locationName:"ccEmailAddresses" type:"list" json:",omitempty"`
	CategoryCode      *string   `locationName:"categoryCode" type:"string" json:",omitempty"`
	CommunicationBody *string   `locationName:"communicationBody" type:"string" json:",omitempty"`
	IssueType         *string   `locationName:"issueType" type:"string" json:",omitempty"`
	Language          *string   `locationName:"language" type:"string" json:",omitempty"`
	ServiceCode       *string   `locationName:"serviceCode" type:"string" json:",omitempty"`
	SeverityCode      *string   `locationName:"severityCode" type:"string" json:",omitempty"`
	Subject           *string   `locationName:"subject" type:"string" json:",omitempty"`

	metadataCreateCaseInput `json:"-", xml:"-"`
}

type metadataCreateCaseInput struct {
	SDKShapeTraits bool `type:"structure" required:"subject,communicationBody" json:",omitempty"`
}

type CreateCaseOutput struct {
	CaseID *string `locationName:"caseId" type:"string" json:",omitempty"`

	metadataCreateCaseOutput `json:"-", xml:"-"`
}

type metadataCreateCaseOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeAttachmentInput struct {
	AttachmentID *string `locationName:"attachmentId" type:"string" json:",omitempty"`

	metadataDescribeAttachmentInput `json:"-", xml:"-"`
}

type metadataDescribeAttachmentInput struct {
	SDKShapeTraits bool `type:"structure" required:"attachmentId" json:",omitempty"`
}

type DescribeAttachmentLimitExceeded struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataDescribeAttachmentLimitExceeded `json:"-", xml:"-"`
}

type metadataDescribeAttachmentLimitExceeded struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeAttachmentOutput struct {
	Attachment *Attachment `locationName:"attachment" type:"structure" json:",omitempty"`

	metadataDescribeAttachmentOutput `json:"-", xml:"-"`
}

type metadataDescribeAttachmentOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeCasesInput struct {
	AfterTime             *string   `locationName:"afterTime" type:"string" json:",omitempty"`
	BeforeTime            *string   `locationName:"beforeTime" type:"string" json:",omitempty"`
	CaseIDList            []*string `locationName:"caseIdList" type:"list" json:",omitempty"`
	DisplayID             *string   `locationName:"displayId" type:"string" json:",omitempty"`
	IncludeCommunications *bool     `locationName:"includeCommunications" type:"boolean" json:",omitempty"`
	IncludeResolvedCases  *bool     `locationName:"includeResolvedCases" type:"boolean" json:",omitempty"`
	Language              *string   `locationName:"language" type:"string" json:",omitempty"`
	MaxResults            *int      `locationName:"maxResults" type:"integer" json:",omitempty"`
	NextToken             *string   `locationName:"nextToken" type:"string" json:",omitempty"`

	metadataDescribeCasesInput `json:"-", xml:"-"`
}

type metadataDescribeCasesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeCasesOutput struct {
	Cases     []*CaseDetails `locationName:"cases" type:"list" json:",omitempty"`
	NextToken *string        `locationName:"nextToken" type:"string" json:",omitempty"`

	metadataDescribeCasesOutput `json:"-", xml:"-"`
}

type metadataDescribeCasesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeCommunicationsInput struct {
	AfterTime  *string `locationName:"afterTime" type:"string" json:",omitempty"`
	BeforeTime *string `locationName:"beforeTime" type:"string" json:",omitempty"`
	CaseID     *string `locationName:"caseId" type:"string" json:",omitempty"`
	MaxResults *int    `locationName:"maxResults" type:"integer" json:",omitempty"`
	NextToken  *string `locationName:"nextToken" type:"string" json:",omitempty"`

	metadataDescribeCommunicationsInput `json:"-", xml:"-"`
}

type metadataDescribeCommunicationsInput struct {
	SDKShapeTraits bool `type:"structure" required:"caseId" json:",omitempty"`
}

type DescribeCommunicationsOutput struct {
	Communications []*Communication `locationName:"communications" type:"list" json:",omitempty"`
	NextToken      *string          `locationName:"nextToken" type:"string" json:",omitempty"`

	metadataDescribeCommunicationsOutput `json:"-", xml:"-"`
}

type metadataDescribeCommunicationsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeServicesInput struct {
	Language        *string   `locationName:"language" type:"string" json:",omitempty"`
	ServiceCodeList []*string `locationName:"serviceCodeList" type:"list" json:",omitempty"`

	metadataDescribeServicesInput `json:"-", xml:"-"`
}

type metadataDescribeServicesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeServicesOutput struct {
	Services []*Service `locationName:"services" type:"list" json:",omitempty"`

	metadataDescribeServicesOutput `json:"-", xml:"-"`
}

type metadataDescribeServicesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeSeverityLevelsInput struct {
	Language *string `locationName:"language" type:"string" json:",omitempty"`

	metadataDescribeSeverityLevelsInput `json:"-", xml:"-"`
}

type metadataDescribeSeverityLevelsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeSeverityLevelsOutput struct {
	SeverityLevels []*SeverityLevel `locationName:"severityLevels" type:"list" json:",omitempty"`

	metadataDescribeSeverityLevelsOutput `json:"-", xml:"-"`
}

type metadataDescribeSeverityLevelsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckRefreshStatusesInput struct {
	CheckIDs []*string `locationName:"checkIds" type:"list" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckRefreshStatusesInput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckRefreshStatusesInput struct {
	SDKShapeTraits bool `type:"structure" required:"checkIds" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckRefreshStatusesOutput struct {
	Statuses []*TrustedAdvisorCheckRefreshStatus `locationName:"statuses" type:"list" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckRefreshStatusesOutput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckRefreshStatusesOutput struct {
	SDKShapeTraits bool `type:"structure" required:"statuses" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckResultInput struct {
	CheckID  *string `locationName:"checkId" type:"string" json:",omitempty"`
	Language *string `locationName:"language" type:"string" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckResultInput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckResultInput struct {
	SDKShapeTraits bool `type:"structure" required:"checkId" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckResultOutput struct {
	Result *TrustedAdvisorCheckResult `locationName:"result" type:"structure" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckResultOutput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckResultOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckSummariesInput struct {
	CheckIDs []*string `locationName:"checkIds" type:"list" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckSummariesInput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckSummariesInput struct {
	SDKShapeTraits bool `type:"structure" required:"checkIds" json:",omitempty"`
}

type DescribeTrustedAdvisorCheckSummariesOutput struct {
	Summaries []*TrustedAdvisorCheckSummary `locationName:"summaries" type:"list" json:",omitempty"`

	metadataDescribeTrustedAdvisorCheckSummariesOutput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorCheckSummariesOutput struct {
	SDKShapeTraits bool `type:"structure" required:"summaries" json:",omitempty"`
}

type DescribeTrustedAdvisorChecksInput struct {
	Language *string `locationName:"language" type:"string" json:",omitempty"`

	metadataDescribeTrustedAdvisorChecksInput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorChecksInput struct {
	SDKShapeTraits bool `type:"structure" required:"language" json:",omitempty"`
}

type DescribeTrustedAdvisorChecksOutput struct {
	Checks []*TrustedAdvisorCheckDescription `locationName:"checks" type:"list" json:",omitempty"`

	metadataDescribeTrustedAdvisorChecksOutput `json:"-", xml:"-"`
}

type metadataDescribeTrustedAdvisorChecksOutput struct {
	SDKShapeTraits bool `type:"structure" required:"checks" json:",omitempty"`
}

type InternalServerError struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataInternalServerError `json:"-", xml:"-"`
}

type metadataInternalServerError struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RecentCaseCommunications struct {
	Communications []*Communication `locationName:"communications" type:"list" json:",omitempty"`
	NextToken      *string          `locationName:"nextToken" type:"string" json:",omitempty"`

	metadataRecentCaseCommunications `json:"-", xml:"-"`
}

type metadataRecentCaseCommunications struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RefreshTrustedAdvisorCheckInput struct {
	CheckID *string `locationName:"checkId" type:"string" json:",omitempty"`

	metadataRefreshTrustedAdvisorCheckInput `json:"-", xml:"-"`
}

type metadataRefreshTrustedAdvisorCheckInput struct {
	SDKShapeTraits bool `type:"structure" required:"checkId" json:",omitempty"`
}

type RefreshTrustedAdvisorCheckOutput struct {
	Status *TrustedAdvisorCheckRefreshStatus `locationName:"status" type:"structure" json:",omitempty"`

	metadataRefreshTrustedAdvisorCheckOutput `json:"-", xml:"-"`
}

type metadataRefreshTrustedAdvisorCheckOutput struct {
	SDKShapeTraits bool `type:"structure" required:"status" json:",omitempty"`
}

type ResolveCaseInput struct {
	CaseID *string `locationName:"caseId" type:"string" json:",omitempty"`

	metadataResolveCaseInput `json:"-", xml:"-"`
}

type metadataResolveCaseInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ResolveCaseOutput struct {
	FinalCaseStatus   *string `locationName:"finalCaseStatus" type:"string" json:",omitempty"`
	InitialCaseStatus *string `locationName:"initialCaseStatus" type:"string" json:",omitempty"`

	metadataResolveCaseOutput `json:"-", xml:"-"`
}

type metadataResolveCaseOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Service struct {
	Categories []*Category `locationName:"categories" type:"list" json:",omitempty"`
	Code       *string     `locationName:"code" type:"string" json:",omitempty"`
	Name       *string     `locationName:"name" type:"string" json:",omitempty"`

	metadataService `json:"-", xml:"-"`
}

type metadataService struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SeverityLevel struct {
	Code *string `locationName:"code" type:"string" json:",omitempty"`
	Name *string `locationName:"name" type:"string" json:",omitempty"`

	metadataSeverityLevel `json:"-", xml:"-"`
}

type metadataSeverityLevel struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TrustedAdvisorCategorySpecificSummary struct {
	CostOptimizing *TrustedAdvisorCostOptimizingSummary `locationName:"costOptimizing" type:"structure" json:",omitempty"`

	metadataTrustedAdvisorCategorySpecificSummary `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCategorySpecificSummary struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TrustedAdvisorCheckDescription struct {
	Category    *string   `locationName:"category" type:"string" json:",omitempty"`
	Description *string   `locationName:"description" type:"string" json:",omitempty"`
	ID          *string   `locationName:"id" type:"string" json:",omitempty"`
	Metadata    []*string `locationName:"metadata" type:"list" json:",omitempty"`
	Name        *string   `locationName:"name" type:"string" json:",omitempty"`

	metadataTrustedAdvisorCheckDescription `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCheckDescription struct {
	SDKShapeTraits bool `type:"structure" required:"id,name,description,category,metadata" json:",omitempty"`
}

type TrustedAdvisorCheckRefreshStatus struct {
	CheckID                    *string `locationName:"checkId" type:"string" json:",omitempty"`
	MillisUntilNextRefreshable *int64  `locationName:"millisUntilNextRefreshable" type:"long" json:",omitempty"`
	Status                     *string `locationName:"status" type:"string" json:",omitempty"`

	metadataTrustedAdvisorCheckRefreshStatus `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCheckRefreshStatus struct {
	SDKShapeTraits bool `type:"structure" required:"checkId,status,millisUntilNextRefreshable" json:",omitempty"`
}

type TrustedAdvisorCheckResult struct {
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `locationName:"categorySpecificSummary" type:"structure" json:",omitempty"`
	CheckID                 *string                                `locationName:"checkId" type:"string" json:",omitempty"`
	FlaggedResources        []*TrustedAdvisorResourceDetail        `locationName:"flaggedResources" type:"list" json:",omitempty"`
	ResourcesSummary        *TrustedAdvisorResourcesSummary        `locationName:"resourcesSummary" type:"structure" json:",omitempty"`
	Status                  *string                                `locationName:"status" type:"string" json:",omitempty"`
	Timestamp               *string                                `locationName:"timestamp" type:"string" json:",omitempty"`

	metadataTrustedAdvisorCheckResult `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCheckResult struct {
	SDKShapeTraits bool `type:"structure" required:"checkId,timestamp,status,resourcesSummary,categorySpecificSummary,flaggedResources" json:",omitempty"`
}

type TrustedAdvisorCheckSummary struct {
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `locationName:"categorySpecificSummary" type:"structure" json:",omitempty"`
	CheckID                 *string                                `locationName:"checkId" type:"string" json:",omitempty"`
	HasFlaggedResources     *bool                                  `locationName:"hasFlaggedResources" type:"boolean" json:",omitempty"`
	ResourcesSummary        *TrustedAdvisorResourcesSummary        `locationName:"resourcesSummary" type:"structure" json:",omitempty"`
	Status                  *string                                `locationName:"status" type:"string" json:",omitempty"`
	Timestamp               *string                                `locationName:"timestamp" type:"string" json:",omitempty"`

	metadataTrustedAdvisorCheckSummary `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCheckSummary struct {
	SDKShapeTraits bool `type:"structure" required:"checkId,timestamp,status,resourcesSummary,categorySpecificSummary" json:",omitempty"`
}

type TrustedAdvisorCostOptimizingSummary struct {
	EstimatedMonthlySavings        *float64 `locationName:"estimatedMonthlySavings" type:"double" json:",omitempty"`
	EstimatedPercentMonthlySavings *float64 `locationName:"estimatedPercentMonthlySavings" type:"double" json:",omitempty"`

	metadataTrustedAdvisorCostOptimizingSummary `json:"-", xml:"-"`
}

type metadataTrustedAdvisorCostOptimizingSummary struct {
	SDKShapeTraits bool `type:"structure" required:"estimatedMonthlySavings,estimatedPercentMonthlySavings" json:",omitempty"`
}

type TrustedAdvisorResourceDetail struct {
	IsSuppressed *bool     `locationName:"isSuppressed" type:"boolean" json:",omitempty"`
	Metadata     []*string `locationName:"metadata" type:"list" json:",omitempty"`
	Region       *string   `locationName:"region" type:"string" json:",omitempty"`
	ResourceID   *string   `locationName:"resourceId" type:"string" json:",omitempty"`
	Status       *string   `locationName:"status" type:"string" json:",omitempty"`

	metadataTrustedAdvisorResourceDetail `json:"-", xml:"-"`
}

type metadataTrustedAdvisorResourceDetail struct {
	SDKShapeTraits bool `type:"structure" required:"status,region,resourceId,metadata" json:",omitempty"`
}

type TrustedAdvisorResourcesSummary struct {
	ResourcesFlagged    *int64 `locationName:"resourcesFlagged" type:"long" json:",omitempty"`
	ResourcesIgnored    *int64 `locationName:"resourcesIgnored" type:"long" json:",omitempty"`
	ResourcesProcessed  *int64 `locationName:"resourcesProcessed" type:"long" json:",omitempty"`
	ResourcesSuppressed *int64 `locationName:"resourcesSuppressed" type:"long" json:",omitempty"`

	metadataTrustedAdvisorResourcesSummary `json:"-", xml:"-"`
}

type metadataTrustedAdvisorResourcesSummary struct {
	SDKShapeTraits bool `type:"structure" required:"resourcesProcessed,resourcesFlagged,resourcesIgnored,resourcesSuppressed" json:",omitempty"`
}