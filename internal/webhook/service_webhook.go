package webhook

type ServiceWebHook struct {
	repo *RepoWebHook
}

func NewService(s *RepoWebHook) *ServiceWebHook {
	return &ServiceWebHook{repo: s}
}
