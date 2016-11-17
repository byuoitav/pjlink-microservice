package health

import (
	"net/http"

	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

// Check returns a message indicating that the service is awake and listening
func Check(context echo.Context) error {
	return jsonresp.New(context, http.StatusOK, "Uh, we had a slight weapons malfunction, but uh...everything's perfectly all right now. We're fine. We're all fine here now, thank you. How are you?")
}
