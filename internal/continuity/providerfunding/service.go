package providerfunding

type SnapshotLoader interface{ Buckets() map[string][]string }
type legacyLoader struct{ buckets map[string][]string }

func (l *legacyLoader) Buckets() map[string][]string { return l.buckets }

type FundingIndex struct{ buckets map[string][]string }

func NewFundingIndex(loader SnapshotLoader) *FundingIndex {
	values := loader.Buckets()
	if values == nil {
		values = make(map[string][]string)
	}
	return &FundingIndex{buckets: values}
}
func (i *FundingIndex) Add(provider, value string) {
	i.buckets[provider] = append(i.buckets[provider], value)
}
func (i *FundingIndex) Values(provider string) []string {
	return append([]string(nil), i.buckets[provider]...)
}
