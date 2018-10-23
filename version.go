package config

type Version struct {
	versionID string
	projectID string
}

func NewVersion(projectId, id string) Version {
	return Version{
		versionID: id,
		projectID: projectId,
	}
}

func (v Version) ID() string {
	return v.versionID
}

func (v Version) ProjectID() string {
	return v.projectID
}
