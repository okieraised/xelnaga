package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type apiServer struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewAPIServer(engine *gin.Engine, logger *zap.Logger) {
	app := &apiServer{
		engine: engine,
		logger: logger,
	}
	app.init()
}

func (r *apiServer) init() {

}

//Register("/instances", ListResources<ResourceType_Instance>);
//Register("/patients", ListResources<ResourceType_Patient>);
//Register("/series", ListResources<ResourceType_Series>);
//Register("/studies", ListResources<ResourceType_Study>);
//
//Register("/instances/{id}", DeleteSingleResource<ResourceType_Instance>);
//Register("/instances/{id}", GetSingleResource<ResourceType_Instance>);
//Register("/patients/{id}", DeleteSingleResource<ResourceType_Patient>);
//Register("/patients/{id}", GetSingleResource<ResourceType_Patient>);
//Register("/series/{id}", DeleteSingleResource<ResourceType_Series>);
//Register("/series/{id}", GetSingleResource<ResourceType_Series>);
//Register("/studies/{id}", DeleteSingleResource<ResourceType_Study>);
//Register("/studies/{id}", GetSingleResource<ResourceType_Study>);
//
//Register("/instances/{id}/statistics", GetResourceStatistics);
//Register("/patients/{id}/statistics", GetResourceStatistics);
//Register("/studies/{id}/statistics", GetResourceStatistics);
//Register("/series/{id}/statistics", GetResourceStatistics);
//
//Register("/patients/{id}/shared-tags", GetSharedTags);
//Register("/series/{id}/shared-tags", GetSharedTags);
//Register("/studies/{id}/shared-tags", GetSharedTags);
//
//Register("/instances/{id}/module", GetModule<ResourceType_Instance, DicomModule_Instance>);
//Register("/patients/{id}/module", GetModule<ResourceType_Patient, DicomModule_Patient>);
//Register("/series/{id}/module", GetModule<ResourceType_Series, DicomModule_Series>);
//Register("/studies/{id}/module", GetModule<ResourceType_Study, DicomModule_Study>);
//Register("/studies/{id}/module-patient", GetModule<ResourceType_Study, DicomModule_Patient>);
//
//Register("/instances/{id}/file", GetInstanceFile);
//Register("/instances/{id}/export", ExportInstanceFile);
//Register("/instances/{id}/tags", GetInstanceTags);
//Register("/instances/{id}/simplified-tags", GetInstanceSimplifiedTags);
//Register("/instances/{id}/frames", ListFrames);
//
//Register("/instances/{id}/frames/{frame}", RestApi::AutoListChildren);
//Register("/instances/{id}/frames/{frame}/preview", GetImage<ImageExtractionMode_Preview>);
//Register("/instances/{id}/frames/{frame}/rendered", GetRenderedFrame);
//Register("/instances/{id}/frames/{frame}/image-uint8", GetImage<ImageExtractionMode_UInt8>);
//Register("/instances/{id}/frames/{frame}/image-uint16", GetImage<ImageExtractionMode_UInt16>);
//Register("/instances/{id}/frames/{frame}/image-int16", GetImage<ImageExtractionMode_Int16>);
//Register("/instances/{id}/frames/{frame}/matlab", GetMatlabImage);
//Register("/instances/{id}/frames/{frame}/raw", GetRawFrame<false>);
//Register("/instances/{id}/frames/{frame}/raw.gz", GetRawFrame<true>);
//Register("/instances/{id}/frames/{frame}/numpy", GetNumpyFrame);  // New in Orthanc 1.10.0
//Register("/instances/{id}/pdf", ExtractPdf);
//Register("/instances/{id}/preview", GetImage<ImageExtractionMode_Preview>);
//Register("/instances/{id}/rendered", GetRenderedFrame);
//Register("/instances/{id}/image-uint8", GetImage<ImageExtractionMode_UInt8>);
//Register("/instances/{id}/image-uint16", GetImage<ImageExtractionMode_UInt16>);
//Register("/instances/{id}/image-int16", GetImage<ImageExtractionMode_Int16>);
//Register("/instances/{id}/matlab", GetMatlabImage);
//Register("/instances/{id}/header", GetInstanceHeader);
//Register("/instances/{id}/numpy", GetNumpyInstance);  // New in Orthanc 1.10.0
//
//Register("/patients/{id}/protected", IsProtectedPatient);
//Register("/patients/{id}/protected", SetPatientProtection);
//
//std::vector<std::string> resourceTypes;
//resourceTypes.push_back("patients");
//resourceTypes.push_back("studies");
//resourceTypes.push_back("series");
//resourceTypes.push_back("instances");
//
//for (size_t i = 0; i < resourceTypes.size(); i++)
//{
//Register("/" + resourceTypes[i] + "/{id}/metadata", ListMetadata);
//Register("/" + resourceTypes[i] + "/{id}/metadata/{name}", DeleteMetadata);
//Register("/" + resourceTypes[i] + "/{id}/metadata/{name}", GetMetadata);
//Register("/" + resourceTypes[i] + "/{id}/metadata/{name}", SetMetadata);
//
//Register("/" + resourceTypes[i] + "/{id}/attachments", ListAttachments);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}", DeleteAttachment);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}", GetAttachmentOperations);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}", UploadAttachment);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/compress", ChangeAttachmentCompression<CompressionType_ZlibWithSize>);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/compressed-data", GetAttachmentData<0>);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/compressed-md5", GetAttachmentCompressedMD5);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/compressed-size", GetAttachmentCompressedSize);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/data", GetAttachmentData<1>);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/is-compressed", IsAttachmentCompressed);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/md5", GetAttachmentMD5);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/size", GetAttachmentSize);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/uncompress", ChangeAttachmentCompression<CompressionType_None>);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/info", GetAttachmentInfo);
//Register("/" + resourceTypes[i] + "/{id}/attachments/{name}/verify-md5", VerifyAttachment);
//}
//
//Register("/tools/invalidate-tags", InvalidateTags);
//Register("/tools/lookup", Lookup);
//Register("/tools/find", Find);
//
//Register("/patients/{id}/studies", GetChildResources<ResourceType_Patient, ResourceType_Study>);
//Register("/patients/{id}/series", GetChildResources<ResourceType_Patient, ResourceType_Series>);
//Register("/patients/{id}/instances", GetChildResources<ResourceType_Patient, ResourceType_Instance>);
//Register("/studies/{id}/series", GetChildResources<ResourceType_Study, ResourceType_Series>);
//Register("/studies/{id}/instances", GetChildResources<ResourceType_Study, ResourceType_Instance>);
//Register("/series/{id}/instances", GetChildResources<ResourceType_Series, ResourceType_Instance>);
//
//Register("/studies/{id}/patient", GetParentResource<ResourceType_Study, ResourceType_Patient>);
//Register("/series/{id}/patient", GetParentResource<ResourceType_Series, ResourceType_Patient>);
//Register("/series/{id}/study", GetParentResource<ResourceType_Series, ResourceType_Study>);
//Register("/instances/{id}/patient", GetParentResource<ResourceType_Instance, ResourceType_Patient>);
//Register("/instances/{id}/study", GetParentResource<ResourceType_Instance, ResourceType_Study>);
//Register("/instances/{id}/series", GetParentResource<ResourceType_Instance, ResourceType_Series>);
//
//Register("/patients/{id}/instances-tags", GetChildInstancesTags);
//Register("/studies/{id}/instances-tags", GetChildInstancesTags);
//Register("/series/{id}/instances-tags", GetChildInstancesTags);
//
//Register("/instances/{id}/content/*", GetRawContent);
//
//Register("/series/{id}/ordered-slices", OrderSlices);
//Register("/series/{id}/numpy", GetNumpySeries);  // New in Orthanc 1.10.0
//
//Register("/patients/{id}/reconstruct", ReconstructResource<ResourceType_Patient>);
//Register("/studies/{id}/reconstruct", ReconstructResource<ResourceType_Study>);
//Register("/series/{id}/reconstruct", ReconstructResource<ResourceType_Series>);
//Register("/instances/{id}/reconstruct", ReconstructResource<ResourceType_Instance>);
//Register("/tools/reconstruct", ReconstructAllResources);
//
//Register("/tools/bulk-content", BulkContent);
//Register("/tools/bulk-delete", BulkDelete);
