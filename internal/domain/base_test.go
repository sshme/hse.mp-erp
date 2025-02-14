package domain

import "testing"

func TestNewAnimal(t *testing.T) {
	tests := []struct {
		name       string
		animalType string
		animalName string
		food       int
		number     int
		kindness   int
		wantErr    bool
	}{
		{
			name:       "Valid rabbit",
			animalType: "rabbit",
			animalName: "Bunny",
			food:       2,
			number:     1,
			kindness:   7,
			wantErr:    false,
		},
		{
			name:       "Valid tiger",
			animalType: "tiger",
			animalName: "Tiger",
			food:       10,
			number:     2,
			kindness:   0, // ignored for predators
			wantErr:    false,
		},
		{
			name:       "Invalid animal type",
			animalType: "cat",
			animalName: "Kitty",
			food:       1,
			number:     3,
			kindness:   5,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAnimal(tt.animalType, tt.animalName, tt.food, tt.number, tt.kindness)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAnimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("NewAnimal() returned nil, want non-nil")
			}
		})
	}
}
