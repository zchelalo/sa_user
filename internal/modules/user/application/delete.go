package userApplication

import "context"

func (useCase *UserUseCases) Delete(ctx context.Context, id string) error {
	_, err := useCase.Get(ctx, id)
	if err != nil {
		return err
	}

	return useCase.userRepository.Delete(ctx, id)
}
