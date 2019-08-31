package ngbanks

import (
	"reflect"
	"testing"	
)

func TestGetBank(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name string
		args args
		want Bank
	}{
		// TODO: Add test cases.
		{
			name: "Guaranty trust bank plc",
			args: args{slug: "GTB"},
			want: Bank{"GUARANTY TRUST BANK PLC", "058", "GTB", Ussd{Code: "*737#"}},
		},
		{
			name: "Diamond Bank PLC",
			args: args{slug: "DMB"},
			want: Bank{"DIAMOND BANK PLC", "063", "DMB", Ussd{Code: "*710#"}},
		},
		{
			name: "Union Bank PLC",
			args: args{slug: "UBN"},
			want: Bank{"UNION BANK OF NIGERIA PLC", "032", "UBN", Ussd{Code: "*826#"}},
		},
		{
			name: "Providus Bank Limited",
			args: args{slug: "PVB"},
			want: Bank{"PROVIDUS BANK LIMITED", "101", "PVB", Ussd{Code: ""}},
		},
		{
			name: "Access Bnak Plc",
			args: args{slug: "ACC"},
			want: Bank{"ACCESS BANK PLC", "044", "ACC", Ussd{Code: "*901#"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBank(tt.args.slug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBank() = %v, want %v", got, tt.want)
			}
		})
	}
}
