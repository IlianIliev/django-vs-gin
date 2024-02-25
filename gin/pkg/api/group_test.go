package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/entities"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/services"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGroupMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedGroupService := services.NewMockGroup(ctrl)

	testCases := []struct {
		name               string
		urlParams          gin.Params
		expectedStatusCode int
		setUpMocks         func()
		expectedBody       string
	}{
		{
			name: "AllOk",
			urlParams: gin.Params{
				{Key: "name", Value: "John"},
			},
			expectedStatusCode: http.StatusOK,
			setUpMocks: func() {
				mockedGroupService.EXPECT().GetMember("John").Return(
					entities.Member{Name: "John", Role: "Leader"}, nil,
				)
			},
			expectedBody: `{"name":"John","role":"Leader"}`,
		},
		{
			name: "Member Not Found",
			urlParams: gin.Params{
				{Key: "name", Value: "Jane"},
			},
			expectedStatusCode: http.StatusNotFound,
			setUpMocks: func() {
				mockedGroupService.EXPECT().GetMember("Jane").Return(
					entities.Member{}, services.ErrMemberNotFound,
				)
			},
		},
		{
			name:               "Invalid URI",
			urlParams:          gin.Params{},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `["Validation for 'Name' failed on 'required'"]`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setUpMocks != nil {
				testCase.setUpMocks()
			}

			recorder := httptest.NewRecorder()
			testContext, _ := gin.CreateTestContext(recorder)

			testContext.Request, _ = http.NewRequest(
				"GET",
				"/", // fmt.Sprintf("/%s", tt.queryParams),
				nil,
			)
			testContext.Params = testCase.urlParams
			testContext.Request.Header = http.Header{
				"Content-Type": []string{"application/json"},
			}

			handler := GetGroupMember(mockedGroupService)
			handler(testContext)

			testContext.Writer.WriteHeaderNow()

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)

			actualBody := recorder.Body.String()
			if testCase.expectedBody != "" || actualBody != "" {
				require.JSONEq(t, testCase.expectedBody, actualBody)
			}
		})
	}
}
