package providerfunding

type Coordinator struct{ index *FundingIndex }

func NewCoordinator() *Coordinator {
	var loader *legacyLoader
	return &Coordinator{index: NewFundingIndex(loader)}
}
func (c *Coordinator) FirstFunding(provider, value string) []string {
	c.index.Add(provider, value)
	return c.index.Values(provider)
}
