package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	hautechapi "hautech/api"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StorageResource struct {
	client *hautechapi.ClientWithResponses
}

type StorageResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Key       types.String `tfsdk:"key"`
	Value     types.Map    `tfsdk:"value"`
	CreatorId types.String `tfsdk:"creator_id"`
	Metadata  types.String `tfsdk:"metadata"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func NewStorageResource() resource.Resource {
	return &StorageResource{}
}

func (r *StorageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage"
}

func (r *StorageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a storage record.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the storage record.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"key": schema.StringAttribute{
				Required:    true,
				Description: "The key of the storage record.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"value": schema.MapAttribute{
				Required:    true,
				Description: "The value of the storage record as an object.",
				ElementType: types.StringType,
			},
			"creator_id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the creator of the storage record.",
			},
			"metadata": schema.StringAttribute{
				Computed:    true,
				Description: "The metadata of the storage record as a JSON string.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The creation timestamp of the storage record.",
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The last update timestamp of the storage record.",
			},
		},
	}
}

func (r *StorageResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData != nil {
		r.client = req.ProviderData.(*Context).Client
	}
}

// Helper function to convert a types.Map to a map[string]interface{}
func (r *StorageResource) mapToInterface(valueMap types.Map) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range valueMap.Elements() {
		if strVal, ok := v.(types.String); ok {
			result[k] = strVal.ValueString()
		} else {
			result[k] = v
		}
	}
	return result
}

// Helper function to convert a map[string]interface{} to a types.Map
func (r *StorageResource) interfaceToMap(valueMap map[string]interface{}) (types.Map, error) {
	valueElements := make(map[string]attr.Value)
	for k, v := range valueMap {
		if strVal, ok := v.(string); ok {
			valueElements[k] = types.StringValue(strVal)
		} else {
			// For non-string values, convert to JSON string
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				return types.Map{}, fmt.Errorf("error converting value to JSON: %w", err)
			}
			valueElements[k] = types.StringValue(string(jsonBytes))
		}
	}
	result, diags := types.MapValue(types.StringType, valueElements)
	if diags.HasError() {
		return types.Map{}, fmt.Errorf("error creating map value: %v", diags)
	}
	return result, nil
}

// createStorageRecord creates a new storage record
func (r *StorageResource) createStorageRecord(ctx context.Context, key string, value map[string]interface{}) (*hautechapi.StorageEntity, error) {
	body := hautechapi.CreateStorageRecordParamsDto{
		Key:   key,
		Value: value,
	}

	resp, err := r.client.StorageControllerCreateRecordV1WithResponse(ctx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	if resp.JSON201 == nil {
		return nil, fmt.Errorf("response body is empty")
	}

	return resp.JSON201, nil
}

// readStorageRecord reads a storage record by key
func (r *StorageResource) readStorageRecord(ctx context.Context, key string) (*hautechapi.StorageRecordsResultDto, error) {
	body := hautechapi.GetStorageRecordParamsDto{
		Keys: []string{key},
	}

	resp, err := r.client.StorageControllerGetRecordsV1WithResponse(ctx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	// Check if the record was found
	if resp.JSON200 == nil || len(*resp.JSON200) == 0 {
		return nil, nil // Record not found
	}

	storageRecords := *resp.JSON200
	return &storageRecords[0], nil
}

// updateStorageRecord updates an existing storage record
func (r *StorageResource) updateStorageRecord(ctx context.Context, key string, value map[string]interface{}) (*hautechapi.StorageEntity, error) {
	body := hautechapi.UpdateStorageRecordParamsDto{
		Key:   key,
		Value: value,
	}

	resp, err := r.client.StorageControllerUpdateRecordV1WithResponse(ctx, body)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	if resp.JSON200 == nil {
		return nil, fmt.Errorf("response body is empty")
	}

	return resp.JSON200, nil
}

// deleteStorageRecord deletes a storage record by key
func (r *StorageResource) deleteStorageRecord(ctx context.Context, key string) error {
	body := hautechapi.DeleteStorageParamsDto{
		Key: key,
	}

	resp, err := r.client.StorageControllerDeleteRecordV1WithResponse(ctx, body)
	if err != nil {
		return fmt.Errorf("failed to call API: %w", err)
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected response: %d - body: %s", resp.StatusCode(), string(resp.Body))
	}

	return nil
}

func (r *StorageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StorageResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the value map to a map[string]interface{}
	valueMap := r.mapToInterface(data.Value)

	// Create the storage record
	storageEntity, err := r.createStorageRecord(ctx, data.Key.ValueString(), valueMap)
	if err != nil {
		resp.Diagnostics.AddError("Error creating storage record", err.Error())
		return
	}

	// Update the model with the response data
	data.Id = types.StringValue(storageEntity.Id)
	data.CreatorId = types.StringValue(storageEntity.CreatorId)

	// Convert metadata to JSON string
	metadataJSON, err := mapToJSONString(storageEntity.Metadata)
	if err != nil {
		resp.Diagnostics.AddError("Error converting metadata to JSON", err.Error())
		return
	}
	data.Metadata = types.StringValue(metadataJSON)

	// Convert the value from the API response to a types.Map
	valueData, err := r.interfaceToMap(storageEntity.Value)
	if err != nil {
		resp.Diagnostics.AddError("Error converting value to map", err.Error())
		return
	}
	data.Value = valueData

	data.CreatedAt = types.StringValue(storageEntity.CreatedAt.Format(time.RFC3339))
	data.UpdatedAt = types.StringValue(storageEntity.UpdatedAt.Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StorageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StorageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the storage record
	storageRecord, err := r.readStorageRecord(ctx, data.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading storage record", err.Error())
		return
	}

	// Check if the record was found
	if storageRecord == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	// Convert the value from the API response to a types.Map
	valueData, err := r.interfaceToMap(storageRecord.Value)
	if err != nil {
		resp.Diagnostics.AddError("Error converting value to map", err.Error())
		return
	}
	data.Value = valueData

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StorageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StorageResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert the value map to a map[string]interface{}
	valueMap := r.mapToInterface(data.Value)

	// Update the storage record
	storageEntity, err := r.updateStorageRecord(ctx, data.Key.ValueString(), valueMap)
	if err != nil {
		resp.Diagnostics.AddError("Error updating storage record", err.Error())
		return
	}

	// Update the model with the response data
	data.Id = types.StringValue(storageEntity.Id)
	data.CreatorId = types.StringValue(storageEntity.CreatorId)

	// Convert metadata to JSON string
	metadataJSON, err := mapToJSONString(storageEntity.Metadata)
	if err != nil {
		resp.Diagnostics.AddError("Error converting metadata to JSON", err.Error())
		return
	}
	data.Metadata = types.StringValue(metadataJSON)

	// Convert the value from the API response to a types.Map
	valueData, err := r.interfaceToMap(storageEntity.Value)
	if err != nil {
		resp.Diagnostics.AddError("Error converting value to map", err.Error())
		return
	}
	data.Value = valueData

	data.CreatedAt = types.StringValue(storageEntity.CreatedAt.Format(time.RFC3339))
	data.UpdatedAt = types.StringValue(storageEntity.UpdatedAt.Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StorageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StorageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the storage record
	err := r.deleteStorageRecord(ctx, data.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting storage record", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

// Helper function for JSON conversion (still needed for metadata)
func mapToJSONString(m map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
