package output

import "testing"

func TestOptions_providesTemplateData(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{name: "Opt is nil", otpions: nil, want: false},
		{name: "Opt is empty", otpions: &Options{}, want: false},
		{name: "TemplateData is nil", otpions: &Options{TemplateData: nil}, want: false},
		{name: "TemplateData is empty", otpions: &Options{TemplateData: [][]byte{}}, want: false},
		{name: "TemplateData is OK", otpions: &Options{TemplateData: [][]byte{[]byte("ok")}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateData(); got != tt.want {
				t.Errorf("Options.isProvidesTemplateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplate(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{name: "Opt is nil", otpions: nil, want: false},
		{name: "Opt is empty", otpions: &Options{}, want: false},
		{name: "Template is empty (implicit_zero_val)", otpions: &Options{Template: ""}, want: false},
		{name: "Template is OK", otpions: &Options{Template: "testify"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplate(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplateDir(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{name: "Opt is nil", otpions: nil, want: false},
		{name: "Opt is empty", otpions: &Options{}, want: false},
		{name: "Template is empty", otpions: &Options{TemplateDir: ""}, want: false},
		{name: "Template is OK", otpions: &Options{TemplateDir: "testify"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateDir(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
