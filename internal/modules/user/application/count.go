package userApplication

import "context"

func (useCase *UserUseCases) Count(ctx context.Context) (int32, error) {
	return useCase.userRepository.Count(ctx)
}
