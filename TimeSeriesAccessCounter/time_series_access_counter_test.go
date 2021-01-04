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
	MockTime *MockTime
	MockTime2 *MockTime
}

func TestTimeSeriesAccessCounterSuiteInit(t *testing.T) {
	suite.Run(t, new(TimeSeriesAccessCounterSuite))
}

func (t *TimeSeriesAccessCounterSuite) SetupTest() {
	t.MockTime = new(MockTime)
	t.MockTime2 = new(MockTime)
	t.sut = new(TimeSeriesAccessCounter)
	t.sut.UnixTime = t.MockTime
}

func (t *TimeSeriesAccessCounterSuite) TestInsertData() {

	t.MockTime.On("GetUnixNow").Return(0)

	t.sut.Insert("127.0.0.1")

	actual := t.sut.Count("127.0.0.1", 0)
	expected := 1

	t.Equal(expected, actual)
}

func (t *TimeSeriesAccessCounterSuite) TestQueryWithTimeRange() {
	t.MockTime.On("GetUnixNow").Return(5)
	t.sut.Insert("127.0.0.1")

	t.sut.UnixTime = t.MockTime2
	t.MockTime2.On("GetUnixNow").Return(60)
	t.sut.Insert("127.0.0.1")

	actual := t.sut.Count("127.0.0.1", 10)
	expected := 1

	t.Equal(expected, actual)
}
