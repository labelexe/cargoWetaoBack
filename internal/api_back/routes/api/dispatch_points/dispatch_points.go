package dispatch_points

import (
	"Wetao/internal/api_back/app/http/handlers/dispatch_points"
	"github.com/gofiber/fiber/v2"
)

// ApiCargoGroup /api/v1/dispatch_p/::
func ApiDispatchPointsGroup(dispatchPointGr fiber.Router) {
	// Get All Category and category-types
	//dispatchPointsGr := dispatchPointGr.Group("dispatch_p")
	dispatchPointGr.Get("get", dispatch_points.GetDispatchesPoints)
}
