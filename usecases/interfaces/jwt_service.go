package interfaces

type IJwtService interface {
	GenerateToken(email string, isSeller bool) (string, string, error)
	ValidateToken(token string, isRefresh bool) (string, bool, error)
}