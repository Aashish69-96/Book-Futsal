package Auth

import(
  "bookfutsal/models/User"
  "bookfutsal/models/Ground"
  "bookfutsal/database"
  
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v5"
  "github.com/joho/godotenv"
  "fmt"
  "os"
  "log"
  "net/http"
  "time"
)
func HandelLogin(c *gin.Context) {
    var data User.FormData

    if err := c.ShouldBind(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    email := data.Email
    password := data.Password
    id, err := database.LoginQuery(email, password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    if id != 0 {
        // Load the secret key from environment variable
      if err := godotenv.Load(); err != nil {
		      log.Fatal("Error loading .env file")
	      }
        secret := os.Getenv("SECRET")
        if secret == "" {
            log.Fatal("Secret key not found")
        }

        // Create JWT token
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "userid": id,
            "exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
        })

        // Generate token string
        tokenString, err := token.SignedString([]byte(secret))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Token string could not be created"})
            return
        }

        // Set JWT token as a cookie
        c.SetSameSite(http.SameSiteNoneMode)
        c.SetCookie("Auth", tokenString, 3600*24*30, "", "", true, true)
        c.JSON(http.StatusOK, gin.H{"message": "Logged in as user", "user_id": id})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect"})
    }
}

func HandelSignUP(c *gin.Context){
  var data User.FormDataS
  body := c.Request.Body
  fmt.Println(body)

  if err:=c.ShouldBind(&data); err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
    return
  }
  name := data.Name
  contact :=data.Number
	email := data.Email
	password := data.Password
  query:=`insert into "users"("name","contact", "email","password") values($1, $2, $3, $4)`
  err := database.MakeInsertQuery(query,name,contact,email,password)
  if err!=nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
  }
	c.JSON(http.StatusOK, gin.H{"message": "Form submitted successfully"})
}

func HandelFutsalRegister(c *gin.Context){
  var data Ground.Grounddata

  if err:=c.ShouldBind(&data); err!=nil {
    fmt.Println("Test4 ",err)
    c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
    return
  }
  name := data.Name
  location := data.Location
  contact :=data.Number
	email := data.Email
	password := data.Password
  open := data.Open
  end := data.End
  doc := data.Doc
  path := "assets/ground"+name+doc.Filename
  var status bool = false
  err := c.SaveUploadedFile(doc, path)
  if err != nil {
    fmt.Println("Test1")
     c.JSON(http.StatusInternalServerError, "Something went wrong!")
  return
 }
  query:=`insert into "ground"("name","location","contact", "email","password","open","close","document_path","verification_status") values($1, $2, $3, $4, $5, $6, $7, $8, $9)`
  err= database.MakeInsertQuery(query,name,location,contact,email,password,open,end,path,status)
  if err!=nil{
    fmt.Println("Test2")
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
  }
	c.JSON(http.StatusOK, gin.H{"message": "Form submitted successfully"})
}

