package artifacts

import (
	"fmt"
	"time"

	tcclient "github.com/taskcluster/taskcluster/v44/clients/client-go"
	"github.com/taskcluster/taskcluster/v44/clients/client-go/tcobject"
	"github.com/taskcluster/taskcluster/v44/clients/client-go/tcqueue"
	"github.com/taskcluster/taskcluster/v44/workers/generic-worker/gwconfig"
)

type ObjectArtifact struct {
	*BaseArtifact
	// Path is the task-directory-relative path to the file (as given in
	// the task description, for example)
	Path string
	// RawContentFile is the filename of the file containing the data
	// for this artifact.
	RawContentFile string
	// ContentType is used in the Content-Type header.
	ContentType string
}

func (a *ObjectArtifact) RequestObject() interface{} {
	return &tcqueue.ObjectArtifactRequest{
		ContentType: a.ContentType,
		Expires:     a.Expires,
		StorageType: "object",
	}
}

func (a *ObjectArtifact) ResponseObject() interface{} {
	return new(tcqueue.ObjectArtifactResponse)
}

func (a *ObjectArtifact) ProcessResponse(resp interface{}, logger Logger, config *gwconfig.Config) (err error) {
	response := resp.(*tcqueue.ObjectArtifactResponse)
	logger.Infof("Uploading artifact %v from file %v with content type %q and expiry %v", a.Name, a.Path, a.ContentType, a.Expires)
	creds := tcclient.Credentials{
		ClientID:    response.Credentials.ClientID,
		AccessToken: response.Credentials.AccessToken,
		Certificate: response.Credentials.Certificate,
	}
	objsvc := tcobject.New(&creds, config.RootURL)
	return objsvc.UploadFromFile(
		response.ProjectID,
		response.Name,
		a.ContentType,
		time.Time(a.Expires),
		response.UploadID,
		a.RawContentFile,
	)
}

func (a *ObjectArtifact) String() string {
	return fmt.Sprintf("Object Artifact - Name: '%v', Path: '%v', RawContentFile: '%v', Expires: %v, Content-Type: '%v'",
		a.Name,
		a.Path,
		a.RawContentFile,
		a.Expires,
		a.ContentType,
	)
}
