package handler

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/akshaym-3255/go-go-learning/crud-go/models"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCreateNewArticle(t *testing.T) {
// 	mockAlbumService := &AlbumService{}
// 	albumHandler := NewAlbumHandler(mockAlbumService)
// 	mockAlbumService.On("Create", new(models.Article)).Return(nil)
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	var jsonData = []byte(`{Id: "1", Title: "hello", Desc: "test", Contnet: "this is first article"}`)
// 	r, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonData))
// 	w := httptest.NewRecorder()

// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			"sucessfully created a new album return new album in json",
// 			args{
// 				w,
// 				r
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			albumHandler.CreateNewArticle(tt.args.w, tt.args.r)

// 		})
// 	}
// }

// type AlbumService struct {
// 	mock.Mock
// }

// // Create provides a mock function with given fields: _a0
// func (_m *AlbumService) Create(_a0 models.Article) error {
// 	ret := _m.Called(_a0)

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(models.Article) error); ok {
// 		r0 = rf(_a0)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

// // Get provides a mock function with given fields: id
// func (_m *AlbumService) Get(id string) (models.Article, error) {
// 	ret := _m.Called(id)

// 	var r0 models.Article
// 	if rf, ok := ret.Get(0).(func(string) models.Article); ok {
// 		r0 = rf(id)
// 	} else {
// 		r0 = ret.Get(0).(models.Article)
// 	}

// 	var r1 error
// 	if rf, ok := ret.Get(1).(func(string) error); ok {
// 		r1 = rf(id)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }

// // GetAll provides a mock function with given fields:
// func (_m *AlbumService) GetAll() ([]models.Article, error) {
// 	ret := _m.Called()

// 	var r0 []models.Article
// 	if rf, ok := ret.Get(0).(func() []models.Article); ok {
// 		r0 = rf()
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).([]models.Article)
// 		}
// 	}

// 	var r1 error
// 	if rf, ok := ret.Get(1).(func() error); ok {
// 		r1 = rf()
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }
