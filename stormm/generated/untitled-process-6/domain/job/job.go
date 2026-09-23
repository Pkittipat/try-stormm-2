package job

// jobState is yours: add the fields your invariants need, and set them in the apply hooks.
type jobState struct{}

// decidePostJob holds the invariants of "Post Job". Return the events that happen, or an error.
func (a *Job) decidePostJob(cmd PostJob) ([]event, error) {
	// TODO: add the invariants for PostJob: return domain.ErrNotAllowed when the command must be refused.
	return []event{JobPublished{ID: a.id}}, nil
}

// applyJobPublished updates the state when "Job Published" has happened. It must not fail.
func (a *Job) applyJobPublished(e JobPublished) {}
