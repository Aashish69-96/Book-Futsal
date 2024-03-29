package Book

import(
  "bookfutsal/models/Booking"
  "bookfutsal/database"
  "fmt"
  "time"
  "github.com/gin-gonic/gin"
  "net/http"
)
type futsaldetails struct{
   Opentime int `json:"opentime"`
   Closetime int `json:"closetime"`
   Bookedtimes []int `json:"bookedtimes"`
}

func HandelBook(c *gin.Context){
  var data Booking.FormData  
  id := c.Param("id")
  user,_:=c.Get("user")
  fmt.Println(id)
  fmt.Println(user)
  if err:=c.ShouldBind(&data); err!=nil {
    fmt.Println(err)
    c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
    return
  }
	time := data.Time
  query :=`insert into "bookings"("user_id","ground_id","time_interval_id") values($1, $2, $3)`
  if len(time)>2{
    c.JSON(http.StatusBadRequest,gin.H{"Error":"You can't book more than 2 intervals"})
  }else{

    for i:=0; i<len(time); i++{

       err := database.MakeInsertQuery(query,user,id,time[i])
       if err!=nil{
       c.JSON(http.StatusBadRequest,gin.H{"Error":"Already booked"})
       return
       }else{
       
      c.JSON(http.StatusOK, gin.H{"message": "Booked successfully"})
      }
  }
  }
}

func ThrowFutsalDetails(c *gin.Context) {
	query:=`SELECT "open" FROM "ground" WHERE "id"=$1`
  id := c.Param("id")
  var opentime string
  var closetime string
  opentime, err := database.Searchsmt(query,id)
  if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grounds"})
		return
	}
  parsedTime1, err := time.Parse(time.RFC3339, opentime)
    if err != nil {
        fmt.Println("Error parsing timestamp:", err)
        return
  }
  query =`SELECT "close" FROM "ground" WHERE "id"=$1`
  closetime, err = database.Searchsmt(query,id)

	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grounds"})
		return
	}
  
  parsedTime2, err := time.Parse(time.RFC3339, closetime)
    if err!=nil {
      fmt.Println("Error parsing timestamp:",err)
      return
    }

  query =`SELECT "time_interval_id" FROM "bookings" WHERE  "ground_id"=$1`
  rows,err:=database.MakeSearchQuery(query,id)
  if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grounds"})
		return
	}
	defer rows.Close()

	var bookedtimes []int
	for rows.Next() {
		var bookedtime int 
		if err := rows.Scan(&bookedtime); err != nil {
			fmt.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		bookedtimes = append(bookedtimes, bookedtime)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during iteration"})
		return
	}
    var futsaldetail futsaldetails
    futsaldetail.Opentime=parsedTime1.Hour()
    futsaldetail.Closetime=parsedTime2.Hour()
    futsaldetail.Bookedtimes=bookedtimes
  c.JSON(http.StatusOK, futsaldetail)
}

