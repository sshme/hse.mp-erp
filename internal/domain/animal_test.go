package domain

import (
	"testing"
)

func TestRabbit_CanContact(t *testing.T) {
	tests := []struct {
		name     string
		kindness int
		want     bool
	}{
		{
			name:     "Kind rabbit",
			kindness: 7,
			want:     true,
		},
		{
			name:     "Neutral rabbit",
			kindness: 5,
			want:     false,
		},
		{
			name:     "Unkind rabbit",
			kindness: 3,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rabbit := Rabbit{
				Herbo{
					Animal: Animal{
						name:   "Test Rabbit",
						foodKg: 2,
						number: 1,
					},
					kindness: tt.kindness,
				},
			}
			if got := rabbit.CanContact(); got != tt.want {
				t.Errorf("Rabbit.CanContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonkey_CanContact(t *testing.T) {
	tests := []struct {
		name     string
		kindness int
		want     bool
	}{
		{
			name:     "Kind monkey",
			kindness: 8,
			want:     true,
		},
		{
			name:     "Neutral monkey",
			kindness: 6,
			want:     true,
		},
		{
			name:     "Unkind monkey",
			kindness: 4,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkey := Monkey{
				Herbo{
					Animal: Animal{
						name:   "Test Monkey",
						foodKg: 3,
						number: 1,
					},
					kindness: tt.kindness,
				},
			}
			if got := monkey.CanContact(); got != tt.want {
				t.Errorf("Monkey.CanContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnimal_Getters(t *testing.T) {
	animal := Animal{
		name:   "Test Animal",
		foodKg: 5,
		number: 42,
	}

	if got := animal.GetName(); got != "Test Animal" {
		t.Errorf("Animal.GetName() = %v, want %v", got, "Test Animal")
	}

	if got := animal.GetNumber(); got != 42 {
		t.Errorf("Animal.GetNumber() = %v, want %v", got, 42)
	}

	if got := animal.GetFood(); got != 5 {
		t.Errorf("Animal.GetFood() = %v, want %v", got, 5)
	}
}

func TestPredator_Interface(t *testing.T) {
	tests := []struct {
		name     string
		predator IPredator
		wantFood int
		wantName string
		wantNum  int
	}{
		{
			name: "Tiger predator",
			predator: &Tiger{
				Predator: Predator{
					Animal: Animal{
						name:   "Test Tiger",
						foodKg: 10,
						number: 1,
					},
				},
			},
			wantFood: 10,
			wantName: "Test Tiger",
			wantNum:  1,
		},
		{
			name: "Wolf predator",
			predator: &Wolf{
				Predator: Predator{
					Animal: Animal{
						name:   "Test Wolf",
						foodKg: 8,
						number: 2,
					},
				},
			},
			wantFood: 8,
			wantName: "Test Wolf",
			wantNum:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.predator.GetFood(); got != tt.wantFood {
				t.Errorf("Predator.GetFood() = %v, want %v", got, tt.wantFood)
			}
			if got := tt.predator.GetName(); got != tt.wantName {
				t.Errorf("Predator.GetName() = %v, want %v", got, tt.wantName)
			}
			if got := tt.predator.GetNumber(); got != tt.wantNum {
				t.Errorf("Predator.GetNumber() = %v, want %v", got, tt.wantNum)
			}
		})
	}
}
