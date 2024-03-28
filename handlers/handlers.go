package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sagoresarker/firecracker-rest-api/networking"
)

type BridgeResponse struct {
	Bridge          string `json:"bridge"`
	UserID          string `json:"user_id"`
	BridgeIPAddress string `json:"bridge_ip_address"`
	BridgeGatewayIP string `json:"bridge_gateway_ip"`
	Error           string `json:"error,omitempty"`
}

type TapResponse struct {
	Bridge   string `json:"bridge"`
	TapName1 string `json:"tap_name_1"`
	TapName2 string `json:"tap_name_2"`
	Error    string `json:"error,omitempty"`
}

func CreateBridge(c echo.Context) error {
	u := new(networking.Bridge)
	if err := c.Bind(u); err != nil {
		return err
	}

	bridge, userID, bridge_ip_address, bridge_gateway_ip, err := networking.SetupBridgeNetwork(u)
	response := BridgeResponse{
		Bridge:          bridge,
		UserID:          userID,
		BridgeIPAddress: bridge_ip_address,
		BridgeGatewayIP: bridge_gateway_ip,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return c.JSON(http.StatusOK, response)
}

func CreateTap(c echo.Context) error {
	u := new(networking.Tap)
	if err := c.Bind(u); err != nil {
		return err
	}
	bridgeName, tapName1, tapname2, err := networking.SetupTapNetwork(u)

	response := TapResponse{
		Bridge:   bridgeName,
		TapName1: tapName1,
		TapName2: tapname2,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return c.JSON(http.StatusOK, response)
}
