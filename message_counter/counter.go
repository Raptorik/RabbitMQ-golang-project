package message_counter

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

var counter int64

func IncrementCounter() {
	atomic.AddInt64(&counter, 1)
}

func GetCounter() int64 {
	return atomic.LoadInt64(&counter)
}

func GetRecordCountHandler(c *gin.Context) {
	// Connect to database
	db, err := sql.Open("postgres", "postgres://user:postgres@localhost/postgres")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Retrieve number of records
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM messages")
	err = row.Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return count as JSON response
	c.JSON(http.StatusOK, gin.H{"count": count})
}
