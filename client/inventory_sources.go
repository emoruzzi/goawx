package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// InventorySourcesService implements awx inventory sources apis.
type InventorySourcesService struct {
	client *Client
}

// ListInventorySourcesResponse represents `ListInventorySources` endpoint response.
type ListInventorySourcesResponse struct {
	Pagination
	Results []*InventorySource `json:"results"`
}

// ListInventorySourcesResponse represents `ListInventorySources` endpoint response.
type SyncInventorySourcesResponse struct {
	InventoryUpdate int     `json:"inventory_update"`
	ID              int     `json:"id"`
	Type            string  `json:"type"`
	URL             string  `json:"url"`
	Created         string  `json:"created"`
	Modified        string  `json:"modified"`
	Name            string  `json:"name"`
	SourcePath      string  `json:"source_path"`
	Timeout         int     `json:"timeout"`
	LaunchType      string  `json:"launch_type"`
	Status          string  `json:"status"`
	Failed          bool    `json:"failed"`
	Started         *string `json:"started"`
	Finished        *string `json:"finished"`
	CanceledOn      *string `json:"canceled_on"`
	Elapsed         float64 `json:"elapsed"`
	Inventory       int     `json:"inventory"`
	InventorySource int     `json:"inventory_source"`
}

const inventorySourcesAPIEndpoint = "/api/v2/inventory_sources/"
const inventoryUpdatesSourcesAPIEndpoint = "/api/v2/inventory_updates/"

// GetInventorySourceByID shows the details of a awx inventroy sources.
func (i *InventorySourcesService) GetInventorySourceByID(id int, params map[string]string) (*InventorySource, error) {
	result := new(InventorySource)
	endpoint := fmt.Sprintf("%s%d/", inventorySourcesAPIEndpoint, id)
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListInventorySources shows list of awx inventories.
func (i *InventorySourcesService) ListInventorySources(params map[string]string) ([]*InventorySource, *ListInventorySourcesResponse, error) {
	result := new(ListInventorySourcesResponse)
	resp, err := i.client.Requester.GetJSON(inventorySourcesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateInventorySource creates an awx InventorySource.
func (i *InventorySourcesService) CreateInventorySource(data map[string]interface{}, params map[string]string) (*InventorySource, error) {
	mandatoryFields = []string{"name", "inventory"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(InventorySource)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if InventorySource exists and return proper error

	resp, err := i.client.Requester.PostJSON(inventorySourcesAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateInventorySource update an awx InventorySource
func (i *InventorySourcesService) UpdateInventorySource(id int, data map[string]interface{}, params map[string]string) (*InventorySource, error) {
	result := new(InventorySource)
	endpoint := fmt.Sprintf("%s%d", inventorySourcesAPIEndpoint, id)
	fmt.Println(endpoint)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := i.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// SyncInventorySource Sync a awx InventorySource by ID
func (i *InventorySourcesService) SyncInventorySource(id int) (*SyncInventorySourcesResponse, error) {
	result := new(SyncInventorySourcesResponse)
	endpoint := fmt.Sprintf("%s%d/update/", inventorySourcesAPIEndpoint, id)
	resp, err := i.client.Requester.PostJSON(endpoint, nil, result, nil)
	if err != nil {
		return result, err
	}
	if err := CheckResponse(resp); err != nil {
		return result, err
	}

	return result, nil
}

// GetSyncInventorySource retrives the InventorySource information from its ID or Name
func (i *InventorySourcesService) GetSyncInventorySource(id int) (*SyncInventorySourcesResponse, error) {
	result := new(SyncInventorySourcesResponse)
	endpoint := fmt.Sprintf("%s%d", inventoryUpdatesSourcesAPIEndpoint, id)
	resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GetInventorySource retrives the InventorySource information from its ID or Name
func (i *InventorySourcesService) GetInventorySource(id int, params map[string]string) (*InventorySource, error) {
	endpoint := fmt.Sprintf("%s%d", inventorySourcesAPIEndpoint, id)
	result := new(InventorySource)
	resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteInventorySource delete an InventorySource from AWX
func (i *InventorySourcesService) DeleteInventorySource(id int) (*InventorySource, error) {
	result := new(InventorySource)
	endpoint := fmt.Sprintf("%s%d", inventorySourcesAPIEndpoint, id)

	resp, err := i.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
