package llm

type Client interface {
	GenerateCommitMessage(diff string) (string, error)
}
