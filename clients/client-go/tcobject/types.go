// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcobject

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster/v38/clients/client-go"
)

type (
	Details struct {
		URL string `json:"url"`
	}

	// Download object request. See [Download Methods](https://docs.taskcluster.net/docs/reference/platform/object/upload-download-methods#download-methods) for more detail.
	DownloadObjectRequest struct {

		// Array items:
		// Supported download methods.
		//
		// Possible values:
		//   * "HTTP:GET"
		AcceptDownloadMethods []string `json:"acceptDownloadMethods"`
	}

	// Download object response.
	//
	// One of:
	//   * DownloadObjectResponse1
	DownloadObjectResponse json.RawMessage

	// Download object response.
	DownloadObjectResponse1 struct {
		Details Details `json:"details"`

		// Constant value: "HTTP:GET"
		Protocol string `json:"protocol"`
	}

	// Representation of the object entry to insert.
	UploadObjectRequest struct {

		// Date at which this entry expires from the object table.
		Expires tcclient.Time `json:"expires"`

		// Project identifier.
		ProjectID string `json:"projectId"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// DownloadObjectResponse is of type json.RawMessage...
func (this *DownloadObjectResponse) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *DownloadObjectResponse) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("DownloadObjectResponse: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
