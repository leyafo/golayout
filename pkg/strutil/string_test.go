package strutil

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{str: "test_case"}, want: "TestCase"},
		{name: "test2", args: args{str: "_test_case"}, want: "TestCase"},
		{name: "test3", args: args{str: "test case"}, want: "TestCase"},
		{name: "test4", args: args{str: " test_case"}, want: "TestCase"},
		{name: "test5", args: args{str: ",test2324_,^$^!@&case"}, want: "TestCase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.str); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
