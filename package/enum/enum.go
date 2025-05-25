package enum

const (
	_ = 5 * iota
	UserTypeBasic
	UserTypeAdmin
	UserTypeModerator
)

const (
	_ = 5 * iota
	UserStatusActive
	UserStatusSuspended
	UserStatusBanned
)
