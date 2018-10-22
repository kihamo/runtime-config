package config

type Version struct {
	versionID string
	projectID uint64
}

func NewVersion(projectId uint64, id string) Version {
	return Version{
		versionID: id,
		projectID: projectId,
	}
}

func (v Version) ID() string {
	return v.versionID
}

func (v Version) ProjectID() uint64 {
	return v.projectID
}
