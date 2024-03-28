package networking

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

type Tap struct {
	BridgeName string `json:"bridge_name"`
}

func createTap(tapName, bridgeName string) error {
	// Create a new TAP interface
	tapLink := &netlink.Tuntap{
		LinkAttrs: netlink.LinkAttrs{
			Name: tapName,
		},
		Mode: netlink.TUNTAP_MODE_TAP,
	}
	if err := netlink.LinkAdd(tapLink); err != nil {
		return fmt.Errorf("failed to create tap: %v", err)
	}

	// Bring up the TAP interface
	if err := netlink.LinkSetUp(tapLink); err != nil {
		return fmt.Errorf("failed to bring up tap: %v", err)
	}

	// Get the bridge link
	bridgeLink, err := netlink.LinkByName(bridgeName)
	if err != nil {
		return fmt.Errorf("failed to get bridge link: %v", err)
	}

	// Assign the TAP interface to the bridge
	if err := netlink.LinkSetMaster(tapLink, bridgeLink.(*netlink.Bridge)); err != nil {
		return fmt.Errorf("failed to assign tap to bridge: %v", err)
	}

	fmt.Printf("Tap %s assigned to Bridge %s\n", tapName, bridgeName)
	return nil
}

func SetupTapNetwork(u *Tap) (string, string, string, error) {
	fmt.Println("Setting up tap")

	tapName1 := "tap-" + u.BridgeName + "-1"
	tapName2 := "tap-" + u.BridgeName + "-2"
	fmt.Println("Tap1:", tapName1)
	fmt.Println("Tap2:", tapName2)

	if err := createTap(tapName1, u.BridgeName); err != nil {
		fmt.Println("Error creating tap for VM1:", err)
		return "", "", "", err
	}
	if err := createTap(tapName2, u.BridgeName); err != nil {
		fmt.Println("Error creating tap for VM2:", err)
		return "", "", "", err
	}
	return u.BridgeName, tapName1, tapName2, nil
}
