package TimeSeriesAccessCounter

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockTime struct {
	mock.Mock
}

func (m *MockTime) GetUnixNow() int64 {
	args := m.Called()
	return int64(args.Int(0))
}

type TimeSeriesAccessCounterSuite struct {
	suite.Suite
	sut      *TimeSeriesAccessCounter
}

func TestTimeSeriesAccessCounterSuiteInit(t *testing.T) {
	suite.Run(t, new(TimeSeriesAccessCounterSuite))
}

func (t *TimeSeriesAccessCounterSuite) SetupTest() {
	t.sut = new(TimeSeriesAccessCounter)
	t.sut.UnixTime = new(MockTime)
	t.sut.BufferRange = 60
}

func (t *TimeSeriesAccessCounterSuite) TestInsertData() {

	InsertDataWithTimes(t,5)
	SetCurrentTimeTo(t, 5)

	actual := t.sut.Count("127.0.0.1", 0)
	expected := 1

	t.Equal(expected, actual)
}

func (t *TimeSeriesAccessCounterSuite) TestQueryWithTimeRange() {
	InsertDataWithTimes(t,5, 60)
	SetCurrentTimeTo(t, 60)

	actual := t.sut.Count("127.0.0.1", 10)
	expected := 1

	t.Equal(expected, actual)
}

func SetCurrentTimeTo(t *TimeSeriesAccessCounterSuite, currentTime int) {
	mock := new(MockTime)
	t.sut.UnixTime = mock
	mock.On("GetUnixNow").Return(currentTime)
}

func InsertDataWithTimes(t *TimeSeriesAccessCounterSuite, times ...int) {
	currentMock := t.sut.UnixTime
	defer func() {t.sut.UnixTime = currentMock} ()

	for _, time := range times {
		mock := new(MockTime)
		t.sut.UnixTime = mock
		mock.On("GetUnixNow").Return(time)
		t.sut.Insert("127.0.0.1")
	}
}
