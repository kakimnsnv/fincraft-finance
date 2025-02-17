package usecases

import "context"

//go:generate mockgen -source=income_repository.go -destination=mocks/income_repository_mock.go -package=mocks

// IncomeRepository репозиторий для работы с доходами
type IncomeRepository interface {
	AddIncome(ctx context.Context, userID int64, categoryID int, amount float64, description string) error
}
