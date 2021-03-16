package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/duyquang6/go-rbac-practice/internal/controller"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

const (
	// sessionName is the name of the session.
	sessionName = "rbac-server-session"
)

// PopulateSessionIfNotExist
func PopulateSessionIfNotExist(store sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, w := c.Request, c.Writer
		ctx := c.Request.Context()
		logger := logging.FromContext(ctx).Named("middleware.PopulateSessionIfNotExist")
		// Get or create a session from the store.
		session, err := store.Get(r, sessionName)
		if err != nil {
			logger.Errorw("failed to get session", "error", err)

			// We couldn't get a session (invalid cookie, can't talk to redis,
			// whatever). According to the spec, this can return an error but can never
			// return an empty session. We intentionally discard the error to ensure we
			// have a session.
			session, _ = store.New(r, sessionName)
		}

		// Save the session on the context.
		ctx = controller.WithSession(ctx, session)
		c.Request = r.Clone(ctx)
		// Ensure the session is saved at least once. This is passed to the
		// before-first-byte writer AND called after the middleware executes to
		// ensure the session is always sent.
		var once sync.Once
		saveSession := func() error {
			var err error
			once.Do(func() {
				session := controller.SessionFromContext(ctx)
				if session != nil {
					controller.StoreSessionLastActivity(session, time.Now())
					err = session.Save(r, w)
				}
			})
			return err
		}

		bfbw := &beforeFirstByteWriter{
			ResponseWriter: w,
			before:         saveSession,
			logger:         logger,
		}
		c.Writer = bfbw
		c.Next()

		// Ensure the session is saved - this will happen if no bytes were
		// written (perhaps due to a redirect or empty body).
		if err := saveSession(); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

// beforeFirstByteWriter is a custom http.ResponseWriter with a hook to run
// before the first byte is written. This is useful if you want to store a
// cookie or some other information that must be sent before any body bytes.
type beforeFirstByteWriter struct {
	gin.ResponseWriter

	before func() error
	logger *zap.SugaredLogger
}

func (w *beforeFirstByteWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *beforeFirstByteWriter) WriteHeader(c int) {
	if err := w.before(); err != nil {
		w.logger.Errorw("failed to invoke before() in beforeFirstByteWriter", "error", err)
	}
	w.ResponseWriter.WriteHeader(c)
}

func (w *beforeFirstByteWriter) Write(b []byte) (int, error) {
	if err := w.before(); err != nil {
		return 0, err
	}
	return w.ResponseWriter.Write(b)
}
