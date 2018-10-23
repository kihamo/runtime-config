package internal

type version struct {
	versionID string
	projectID string
}

func NewVersion(projectId, id string) version {
	return version{
		versionID: id,
		projectID: projectId,
	}
}

func (v version) ID() string {
	return v.versionID
}

func (v version) ProjectID() string {
	return v.projectID
}
