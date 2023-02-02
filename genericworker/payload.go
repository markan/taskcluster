package genericworker

import (
	"encoding/json"

	"github.com/taskcluster/d2g/tctime"
)

type (
	Artifact struct {

		// Content-Encoding for the artifact. If not provided, `gzip` will be used, except for the
		// following file extensions, where `identity` will be used, since they are already
		// compressed:
		//
		// * 7z
		// * bz2
		// * deb
		// * dmg
		// * flv
		// * gif
		// * gz
		// * jpeg
		// * jpg
		// * png
		// * swf
		// * tbz
		// * tgz
		// * webp
		// * whl
		// * woff
		// * woff2
		// * xz
		// * zip
		// * zst
		//
		// Note, setting `contentEncoding` on a directory artifact will apply the same content
		// encoding to all the files contained in the directory.
		//
		// Since: generic-worker 16.2.0
		//
		// Possible values:
		//   * "identity"
		//   * "gzip"
		ContentEncoding string `json:"contentEncoding,omitempty"`

		// Explicitly set the value of the HTTP `Content-Type` response header when the artifact(s)
		// is/are served over HTTP(S). If not provided (this property is optional) the worker will
		// guess the content type of artifacts based on the filename extension of the file storing
		// the artifact content. It does this by looking at the system filename-to-mimetype mappings
		// defined in multiple `mime.types` files located under `/etc`. Note, setting `contentType`
		// on a directory artifact will apply the same contentType to all files contained in the
		// directory.
		//
		// See [mime.TypeByExtension](https://pkg.go.dev/mime#TypeByExtension).
		//
		// Since: generic-worker 10.4.0
		ContentType string `json:"contentType,omitempty"`

		// Date when artifact should expire must be in the future, no earlier than task deadline, but
		// no later than task expiry. If not set, defaults to task expiry.
		//
		// Since: generic-worker 1.0.0
		Expires tctime.Time `json:"expires,omitempty"`

		// Name of the artifact, as it will be published. If not set, `path` will be used.
		// Conventionally (although not enforced) path elements are forward slash separated. Example:
		// `public/build/a/house`. Note, no scopes are required to read artifacts beginning `public/`.
		// Artifact names not beginning `public/` are scope-protected (caller requires scopes to
		// download the artifact). See the Queue documentation for more information.
		//
		// Since: generic-worker 8.1.0
		Name string `json:"name,omitempty"`

		// Relative path of the file/directory from the task directory. Note this is not an absolute
		// path as is typically used in docker-worker, since the absolute task directory name is not
		// known when the task is submitted. Example: `dist\regedit.exe`. It doesn't matter if
		// forward slashes or backslashes are used.
		//
		// Since: generic-worker 1.0.0
		Path string `json:"path"`

		// Artifacts can be either an individual `file` or a `directory` containing
		// potentially multiple files with recursively included subdirectories.
		//
		// Since: generic-worker 1.0.0
		//
		// Possible values:
		//   * "file"
		//   * "directory"
		Type string `json:"type"`
	}

	// Requires scope `queue:get-artifact:<artifact-name>`.
	//
	// Since: generic-worker 5.4.0
	ArtifactContent struct {

		// Max length: 1024
		Artifact string `json:"artifact"`

		// The required SHA 256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		Sha256 string `json:"sha256,omitempty"`

		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	// Base64 encoded content of file/archive, up to 64KB (encoded) in size.
	//
	// Since: generic-worker 11.1.0
	Base64Content struct {

		// Base64 encoded content of file/archive, up to 64KB (encoded) in size.
		//
		// Since: generic-worker 11.1.0
		//
		// Syntax:     ^[A-Za-z0-9/+]+[=]{0,2}$
		// Max length: 65536
		Base64 string `json:"base64"`
	}

	// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
	// if all task commands have a zero exit code, or `failed/failed` if any command has a
	// non-zero exit code. This payload property allows customsation of the task resolution
	// based on exit code of task commands.
	ExitCodeHandling struct {

		// Exit codes for any command in the task payload to cause this task to
		// be resolved as `exception/intermittent-task`. Typically the Queue
		// will then schedule a new run of the existing `taskId` (rerun) if not
		// all task runs have been exhausted.
		//
		// See [itermittent tasks](https://docs.taskcluster.net/docs/reference/platform/taskcluster-queue/docs/worker-interaction#intermittent-tasks) for more detail.
		//
		// Since: generic-worker 10.10.0
		//
		// Array items:
		// Mininum:    1
		Retry []int64 `json:"retry,omitempty"`
	}

	// Feature flags enable additional functionality.
	//
	// Since: generic-worker 5.3.0
	FeatureFlags struct {

		// Artifacts named `public/chain-of-trust.json` and
		// `public/chain-of-trust.json.sig` should be generated which will
		// include information for downstream tasks to build a level of trust
		// for the artifacts produced by the task and the environment it ran in.
		//
		// Since: generic-worker 5.3.0
		ChainOfTrust bool `json:"chainOfTrust,omitempty"`

		// The taskcluster proxy provides an easy and safe way to make authenticated
		// taskcluster requests within the scope(s) of a particular task. See
		// [the github project](https://github.com/taskcluster/taskcluster/tree/main/tools/taskcluster-proxy) for more information.
		//
		// Since: generic-worker 10.6.0
		TaskclusterProxy bool `json:"taskclusterProxy,omitempty"`
	}

	FileMount struct {

		// One of:
		//   * ArtifactContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the file.
		//
		// Since: generic-worker 5.4.0
		File string `json:"file"`
	}

	// This schema defines the structure of the `payload` property referred to in a
	// Taskcluster Task definition.
	GenericWorkerPayload struct {

		// Artifacts to be published.
		//
		// Since: generic-worker 1.0.0
		Artifacts []Artifact `json:"artifacts,omitempty"`

		// One array per command (each command is an array of arguments). Several arrays
		// for several commands.
		//
		// Since: generic-worker 0.0.1
		//
		// Array items:
		// Array items:
		Command [][]string `json:"command"`

		// Env vars must be string to __string__ mappings (not number or boolean). For example:
		// ```
		// {
		//   "PATH": "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
		//   "GOOS": "darwin",
		//   "FOO_ENABLE": "true",
		//   "BAR_TOTAL": "3"
		// }
		// ```
		//
		// Note, the following environment variables will automatically be set in the task
		// commands, but may be overridden by environment variables in the task payload:
		//   * `HOME` - the home directory of the task user
		//   * `PATH` - `/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin`
		//   * `USER` - the name of the task user
		//
		// The following environment variables will automatically be set in the task
		// commands, and may not be overridden by environment variables in the task payload:
		//   * `DISPLAY` - `:0` (Linux only)
		//   * `TASK_ID` - the task ID of the currently running task
		//   * `RUN_ID` - the run ID of the currently running task
		//   * `TASKCLUSTER_ROOT_URL` - the root URL of the taskcluster deployment
		//   * `TASKCLUSTER_PROXY_URL` (if taskcluster proxy feature enabled) - the
		//      taskcluster authentication proxy for making unauthenticated taskcluster
		//      API calls
		//   * `TASK_USER_CREDENTIALS` (if config property `runTasksAsCurrentUser` set to
		//     `true` in `generic-worker.config` file - the absolute file location of a
		//     json file containing the current task OS user account name and password.
		//     This is only useful for the generic-worker multiuser CI tasks, where
		//     `runTasksAsCurrentUser` is set to `true`.
		//   * `TASKCLUSTER_WORKER_LOCATION`. See
		//     [RFC #0148](https://github.com/taskcluster/taskcluster-rfcs/blob/master/rfcs/0148-taskcluster-worker-location.md)
		//     for details.
		//
		// Since: generic-worker 0.0.1
		//
		// Map entries:
		Env map[string]string `json:"env,omitempty"`

		// Feature flags enable additional functionality.
		//
		// Since: generic-worker 5.3.0
		Features FeatureFlags `json:"features,omitempty"`

		// Maximum time the task container can run in seconds.
		//
		// Since: generic-worker 0.0.1
		//
		// Mininum:    1
		// Maximum:    86400
		MaxRunTime int64 `json:"maxRunTime"`

		// Directories and/or files to be mounted.
		//
		// Since: generic-worker 5.4.0
		//
		// Array items:
		// One of:
		//   * FileMount
		//   * WritableDirectoryCache
		//   * ReadOnlyDirectory
		Mounts []json.RawMessage `json:"mounts,omitempty"`

		// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
		// if all task commands have a zero exit code, or `failed/failed` if any command has a
		// non-zero exit code. This payload property allows customsation of the task resolution
		// based on exit code of task commands.
		OnExitStatus ExitCodeHandling `json:"onExitStatus,omitempty"`

		// A list of OS Groups that the task user should be a member of. Not yet implemented on
		// non-Windows platforms, therefore this optional property may only be an empty array if
		// provided.
		//
		// Since: generic-worker 6.0.0
		//
		// Array items:
		OSGroups []string `json:"osGroups,omitempty"`

		// This property is allowed for backward compatibility, but is unused.
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	// Byte-for-byte literal inline content of file/archive, up to 64KB in size.
	//
	// Since: generic-worker 11.1.0
	RawContent struct {

		// Byte-for-byte literal inline content of file/archive, up to 64KB in size.
		//
		// Since: generic-worker 11.1.0
		//
		// Max length: 65536
		Raw string `json:"raw"`
	}

	ReadOnlyDirectory struct {

		// One of:
		//   * ArtifactContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of content for read only directory.
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "tar.xz"
		//   * "tar.zst"
		//   * "zip"
		Format string `json:"format"`
	}

	// URL to download content from.
	//
	// Since: generic-worker 5.4.0
	URLContent struct {

		// The required SHA 256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		Sha256 string `json:"sha256,omitempty"`

		// URL to download content from.
		//
		// Since: generic-worker 5.4.0
		URL string `json:"url"`
	}

	WritableDirectoryCache struct {

		// Implies a read/write cache directory volume. A unique name for the
		// cache volume. Requires scope `generic-worker:cache:<cache-name>`.
		// Note if this cache is loaded from an artifact, you will also require
		// scope `queue:get-artifact:<artifact-name>` to use this cache.
		//
		// Since: generic-worker 5.4.0
		CacheName string `json:"cacheName"`

		// One of:
		//   * ArtifactContent
		//   * URLContent
		//   * RawContent
		//   * Base64Content
		Content json.RawMessage `json:"content,omitempty"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of the preloaded content (if `content` provided).
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "tar.xz"
		//   * "tar.zst"
		//   * "zip"
		Format string `json:"format,omitempty"`
	}
)
