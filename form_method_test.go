package formmethod

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestMiddleware(t *testing.T) {
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "PUT",
			args: args{request: func() *http.Request {
				values := url.Values{
					defaultInputName: {"PUT"},
				}
				r, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(values.Encode()))
				if err != nil {
					panic(fmt.Errorf("create request: %w", err))
				}
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}()},
			want: "PUT",
		},
		{
			name: "PUT but lowercase",
			args: args{request: func() *http.Request {
				values := url.Values{
					defaultInputName: {"put"},
				}
				r, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(values.Encode()))
				if err != nil {
					panic(fmt.Errorf("create request: %w", err))
				}
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}()},
			want: "PUT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.want {
					t.Errorf("Middleware() = %v, want %v", r.Method, tt.want)
				}
			}))
			h.ServeHTTP(nil, tt.args.request)
		})
	}
}

func TestTemplateField(t *testing.T) {
	type args struct {
		method string
	}
	tests := []struct {
		name string
		args args
		want template.HTML
	}{
		{
			name: "PUT",
			args: args{method: "PUT"},
			want: template.HTML(`<input type="hidden" name="_method" value="PUT">`),
		},
		{
			name: "PUT but lowercase",
			args: args{method: "put"},
			want: template.HTML(`<input type="hidden" name="_method" value="put">`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TemplateField(tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TemplateField() = %v, want %v", got, tt.want)
			}
		})
	}
}
