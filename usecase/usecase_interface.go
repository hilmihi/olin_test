package usecase

type AppUseCases struct {
	TransactionUseCase GoAssestmentUseCaseInterface
}

func NewAppUseCase() AppUseCases {
	return AppUseCases{
		TransactionUseCase: NewTransactionUseCase(),
	}
}
