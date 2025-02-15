package service

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
	"github.com/andys920605/meme-coin/internal/mock"
	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/pkg/database"
	"github.com/andys920605/meme-coin/pkg/logging"
	"github.com/andys920605/meme-coin/pkg/snowflake"
)

func TestMemeCoinDomainServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Meme Coin Domain Service Suite")
}

var _ = Describe("MemeCoinDomainService CreateMemeCoin", func() {
	var (
		ctrl          *gomock.Controller
		mockRepo      *mock.MockMemeCoinRepository
		mockTM        *mock.MockTransactionManager
		domainService *MemeCoinDomainService
		ctx           context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mock.NewMockMemeCoinRepository(ctrl)
		mockTM = mock.NewMockTransactionManager(ctrl)
		domainService = &MemeCoinDomainService{
			transactionManager: mockTM,
			memeCoinRepository: mockRepo,
		}
		ctx = context.Background()
		snowflake.Init(logging.New())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("成功建立 MemeCoin", func() {
		It("應該正確建立並回傳 MemeCoin", func() {
			cmd := message.CreateMemeCoinCommand{
				Name:        "TestCoin",
				Description: "This is a test meme coin",
			}

			mockTM.EXPECT().Execute(ctx, gomock.Any()).DoAndReturn(
				func(ctx context.Context, fn database.TxFunc) error {
					return fn(ctx)
				},
			).Times(1)

			mockRepo.EXPECT().Save(ctx, gomock.Any()).Do(func(_ context.Context, coin *meme_coin.MemeCoin) {
				Expect(coin.Name).To(Equal(cmd.Name))
				Expect(coin.Description).To(Equal(cmd.Description))
			}).Return(nil).Times(1)

			result, err := domainService.CreateMemeCoin(ctx, cmd)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).NotTo(BeNil())
			Expect(result.Name).To(Equal(cmd.Name))
			Expect(result.Description).To(Equal(cmd.Description))
		})
	})
	// todo: GetMemeCoin, UpdateMemeCoin, DeleteMemeCoin, PokeMemeCoin
})
