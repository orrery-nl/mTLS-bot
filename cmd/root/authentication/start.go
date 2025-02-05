package authentication

var (
	startedAuthenticationFlows StartedAuthenticationFlows
)

type StartedAuthenticationFlows struct {
	Flows map[string]struct {
		State     string
		StartedAt int64
	}
}

func init() {
	startedAuthenticationFlows = StartedAuthenticationFlows{
		Flows: make(map[string]struct {
			State     string
			StartedAt int64
		}),
	}
}

// IsIdUnique - Check if the ID is unique in the `startedAuthenticationFlows` slice.
func IsIdUnique(id string) bool {
	_, exists := startedAuthenticationFlows.Flows[id]
	return !exists
}
