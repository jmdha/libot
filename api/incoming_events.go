package api

type GameStartEvent struct {
	Type string
	Game GameEventinfo
}

type GameFinsihEvent struct {
	Type string
	Game GameEventinfo
}

type GameEventinfo struct {
	ID     string
	Source string
	Winner string
	Compat CompatInfo
}

type CompatInfo struct {
	Bot   bool
	Board bool
}

type ChallengeEvent struct {
	Type      string
	Challenge ChallengeInfo
}

type ChallengeCanceledEvent struct {
	Type      string
	Challenge ChallengeInfo
}

type ChallengeDeclinedEvent struct {
	Type      string
	Challenge ChallengeDeclinedInfo
}

type ChallengeInfo struct {
	ID               string
	URL              string
	Status           string
	Challenger       UserInfo
	DestUser         UserInfo
	Variant          VariantInfo
	Rated            bool
	Speed            string
	TimeControl      TimeControlInfo
	Color            string
	FinalColor       string
	Perf             PerfInfo
	Direction        string
	InitialFen       string
	DeclineReason    string
	DeclineReasonKey string
}

type UserInfo struct {
	Rating      int
	Provisional bool
	Online      bool
	Lag         int
	Name        string
	Title       string
	Patron      bool
}

type VariantInfo struct {
	Key   string
	Name  string
	Short string
}

type TimeControlInfo struct {
	Type      string
	Limit     int
	Increment int
	Show      string
}

type PerfInfo struct {
	Icon string
	Name string
}

type ChallengeDeclinedInfo struct {
	ID string
}
