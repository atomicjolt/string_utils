package canvas

import (
	"reflect"
	"testing"
)

func TestInclude(t *testing.T) {
	strings := []string{"red", "green", "blue"}

	if !Include(strings, "red") {
		t.Errorf("expected Include to return true")
	}

	if Include(strings, "purple") {
		t.Errorf("expected Include to return false")
	}
}

func TestUnique(t *testing.T) {
	strings := []string{"red", "green", "blue", "blue"}

	ret := Unique(strings)
	if len(ret) != 3 {
		t.Errorf("failed remove duplicate values")
	}
}

func TestRemoveAt(t *testing.T) {
	type args struct {
		vs []string
		i  int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "remove item from slice",
			args: args{
				vs: []string{"one", "two"},
				i:  1,
			},
			want: []string{"one"},
		},
		{
			name: "index out of range",
			args: args{
				vs: []string{"one", "two"},
				i:  2,
			},
			want: []string{"one", "two"},
		},
		{
			name: "negative index",
			args: args{
				vs: []string{"one", "two"},
				i:  -1,
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAt(tt.args.vs, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		vs []string
		s  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "remove item from slice",
			args: args{
				vs: []string{"one", "two"},
				s:  "two",
			},
			want: []string{"one"},
		},
		{
			name: "item not in slice",
			args: args{
				vs: []string{"one", "two"},
				s:  "three",
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.vs, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
