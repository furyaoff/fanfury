package keeper_test

import (
	apptesting "github.com/furyaoff/fanfury/app/testing"
	"github.com/furyaoff/fanfury/x/merkledrop/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
