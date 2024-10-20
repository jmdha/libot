package api

type OpponentGoneEvent struct {
	Type              string
	Gone              bool
	ClaimWinInSeconds int
}

type ChatLineEvent struct {
	Type     string
	Room     string
	Username string
	Text     string
}

type GameStateEvent struct {
	Type      string
	Moves     string
	WTime     int
	BTime     int
	WInc      int
	BInc      int
	Status    string
	Winner    string
	WDraw     bool
	BDraw     bool
	WTakeBack bool
	BTakeBack bool
}

type GameFullEvent struct {
	Type         string
	ID           string
	Variant      VariantInfo
	Clock        ClockInfo
	Speed        string
	Perf         PerfInfo
	Rated        bool
	CreatedAt    int
	White        GamePlayerInfo
	Black        GamePlayerInfo
	InitialFen   string
	TournamentID string
	State        GameStateEvent
}

type ClockInfo struct {
	Limit     string
	Increment string
}

type GamePlayerInfo struct {
	AILevel     int
	ID          string
	Name        string
	Title       string
	Rating      int
	Provisional bool
}
