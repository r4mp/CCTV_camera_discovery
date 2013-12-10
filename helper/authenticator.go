package helper

/*import (
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateToken(username string, password string) {

	var userData [3]byte

	t := time.Now()

	_, nonce, err := NewUuid()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	var b bytes.Buffer
	w := base64.NewEncoder(base64.URLEncoding, &b)
	w.Write(data)
	w.Close()
	data := b.Bytes()

	userData[0] = nonce
	userData[1] = t.Format("yyyy-MM-ddThh:mm:ssZ")
	userData[2] = password

}*/
