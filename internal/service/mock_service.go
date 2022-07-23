package service

type MockService struct {
}

func (s MockService) GenerateJwtToken(id string, secret string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0", nil
}

func (s MockService) RefreshJwtToken(jwt string) (string, error) {
	return "ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjU4MDY1ODY0fQ._Q9NIu1anMwPzZ3w0gvQbRQVlHRyZUnyd60LzhfNyL0", nil
}
