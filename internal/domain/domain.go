package domain

type Location struct {
	ID       uint32 // unique
	Place    string
	Country  string // len <= 50
	City     string // len <= 50
	Distance uint32
}

type User struct {
	ID        uint32 // unique
	Email     string // unique, len <= 100
	FirstName string // len <= 50
	LastName  string // len <= 50
	Gender    string // m/f
	BirthDate int64  // timestamp
}

type Visit struct {
	ID         uint32 // unique
	LocationID uint32
	UserID     uint32
	VisitedAt  int64 // timestamp
	Mark       int   // 0-5
}
