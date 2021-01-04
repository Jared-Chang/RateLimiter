package TimeSeriesAccessCounter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TimeSeriesAccessCounterSuite struct {
	suite.Suite
	*TimeSeriesAccessCounter
}

func TestTimeSeriesAccessCounterSuiteInit(t *testing.T) {
	suite.Run(t, new(TimeSeriesAccessCounterSuite))
}

func (t *TimeSeriesAccessCounterSuite) SetupTest() {
	t.TimeSeriesAccessCounter = new(TimeSeriesAccessCounter)
}

func (t *TimeSeriesAccessCounterSuite) TestInsertData() {
	t.TimeSeriesAccessCounter.Insert("127.0.0.1")

	actual := t.TimeSeriesAccessCounter.Count("127.0.0.1", 0)
	expected := 1

	t.Equal(expected, actual)
}
