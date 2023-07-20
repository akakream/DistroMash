package controllers

import (
	"github.com/akakream/DistroMash/pkg/repository/peer"
	"github.com/gofiber/fiber/v2"
)

func GetPeersListUI(c *fiber.Ctx) error {
    identity, err := peer.GetIdentity()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    peers, err := peer.GetPeersList()
	// Return status 500 Internal Server Error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// peers := strings.Split(data.Peers, ",")

	return c.Render("peers", fiber.Map{
		"Peers": peers.Peers,
        "HostID": identity.ID,
        "HostAddrs": identity.Addrs,
	}, "base")
}

