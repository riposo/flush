package internal_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/riposo/flush/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/conn/cache"
	"github.com/riposo/riposo/pkg/conn/permission"
	"github.com/riposo/riposo/pkg/conn/storage"
	"github.com/riposo/riposo/pkg/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flush handler", func() {
	var txn *api.Txn

	BeforeEach(func() {
		txn = mock.Txn()
		txn.Store = &mockStore{}
		txn.Perms = &mockPerms{}
		txn.Cache = &mockCache{}
	})

	AfterEach(func() {
		Expect(txn.Abort()).To(Succeed())
	})

	It("should flush and serve", func() {
		w := httptest.NewRecorder()
		r := mock.Request(txn, "POST", "/flush", nil)
		api.HandlerFunc(internal.Handler).ServeHTTP(w, r)
		Expect(w.Code).To(Equal(http.StatusAccepted))
		Expect(w.Body.Bytes()).To(MatchJSON(`{}`))

		Expect(txn.Store.(*mockStore).flushed).To(BeTrue())
		Expect(txn.Perms.(*mockPerms).flushed).To(BeTrue())
		Expect(txn.Cache.(*mockCache).flushed).To(BeTrue())
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal")
}

type mockStore struct {
	storage.Backend
	storage.Transaction
	flushed bool
}

func (m *mockStore) Close() error                                       { return nil }
func (m *mockStore) Begin(context.Context) (storage.Transaction, error) { return m, nil }
func (m *mockStore) Rollback() error                                    { return nil }
func (m *mockStore) Flush() error {
	m.flushed = true
	return nil
}

type mockPerms struct {
	permission.Backend
	permission.Transaction
	flushed bool
}

func (m *mockPerms) Close() error                                          { return nil }
func (m *mockPerms) Begin(context.Context) (permission.Transaction, error) { return m, nil }
func (m *mockPerms) Rollback() error                                       { return nil }
func (m *mockPerms) Flush() error {
	m.flushed = true
	return nil
}

type mockCache struct {
	cache.Backend
	cache.Transaction
	flushed bool
}

func (m *mockCache) Close() error                                     { return nil }
func (m *mockCache) Begin(context.Context) (cache.Transaction, error) { return m, nil }
func (m *mockCache) Rollback() error                                  { return nil }
func (m *mockCache) Flush() error {
	m.flushed = true
	return nil
}
