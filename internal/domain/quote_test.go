package domain

import (
	"math"
	"strings"
	"testing"
)

func Test_Quote_ValidateID(t *testing.T) {
	tests := []struct {
		testName string
		wantErr  bool
		field    int
	}{
		{"ID не должен быть равен 0", true, 0},
		{"ID не должен быть отрицательным числом", true, -100},
		{"ID может быть целым числом", false, 1},
		{"ID может быть целым числом", false, 100},
		{"ID может быть целым числом", false, 99999},
		{"ID может быть максимальным значением для типа int", false, math.MaxInt},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			quote := Quote{
				ID: tt.field,
			}
			if err := quote.ValidateID(); tt.wantErr == (err != nil) {
			} else {
				t.Errorf("unexpected error %v", err)
			}
		})
	}

}

func Test_Quote_ValidateAuthor(t *testing.T) {
	tests := []struct {
		testName string
		wantErr  bool
		field    string
	}{
		{testName: "пустое имя", wantErr: true, field: ""},
		{testName: "толкьо с пробелами", field: " ", wantErr: true},
		{testName: "управляющие символы", field: "q\n\t\r\a\f\vq", wantErr: true},
		{testName: "в начале пробел", field: " Name", wantErr: true},
		{testName: "в конце пробел", field: "Name ", wantErr: true},
		{testName: "длина имени превышет дозволенную длину", field: strings.Repeat("名", AuthorNameMaxLen+1), wantErr: true},

		{testName: "один символ", field: "q", wantErr: false},
		{testName: "одна цифра", field: "1", wantErr: false},
		{testName: "нижний регистр", field: "name", wantErr: false},
		{testName: "буквы и цифры", field: "Name2", wantErr: false},
		{testName: "имя с пробелом", field: "first last", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			quote := Quote{
				Author: tt.field,
			}
			if err := quote.ValidateAuthor(); tt.wantErr == (err != nil) {
			} else {
				t.Errorf("unexpected error %v", err)
			}
		})
	}

}

func Test_Quote_ValidateText(t *testing.T) {
	tests := []struct {
		testName string
		wantErr  bool
		field    string
	}{
		{"Text не может быть пустым", true, ""},
		{"Text не может содержать только пробелы", true, "     "},

		{"Text может содержать в начала и в конце пробелы", false, "  test   "},
		{"Text может не содержать пробелы", false, "test1234"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			quote := Quote{
				Text: tt.field,
			}
			if err := quote.ValidateText(); tt.wantErr == (err != nil) {
			} else {
				t.Errorf("unexpected error %v", err)
			}
		})
	}

}
