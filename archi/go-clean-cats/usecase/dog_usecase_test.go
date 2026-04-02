package usecase_test

import (
	"go-clean-cats/domain/entity"
	"go-clean-cats/domain/repository"
	"go-clean-cats/usecase"
	"testing"
)

func TestDogUseCase_UpdateDog(t *testing.T) {
	tests := []struct {
		// Named input parameters for receiver constructor.
		dogRepository repository.DogRepository
		// Named input parameters for target function.
		id      string
		name    string
		breed   string
		age     int
		color   string
		want    *entity.Dog
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewDogUseCase(tt.dogRepository)
			got, gotErr := uc.UpdateDog(tt.id, tt.name, tt.breed, tt.age, tt.color)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateDog() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateDog() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("UpdateDog() = %v, want %v", got, tt.want)
			}
		})
	}
}
