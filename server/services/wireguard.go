package services

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"net"
	"time"

	"github.com/spiron09/burnervpn/server/models"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func GenerateKeyPair() (*models.KeyPair, error) {
	privateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	return &models.KeyPair{
		PrivateKey: privateKey.String(),
		PublicKey:  privateKey.PublicKey().String(),
	}, nil
}

func BuildServerConfig(serverPrivateKey, clientPublicKey string) string {
	return fmt.Sprintf(
		`#cloud-config
package_update: true
packages:
  - wireguard

write_files:
  - path: /etc/wireguard/wg0.conf
    content: |
      [Interface]
      PrivateKey = %s
      Address = 10.0.0.1/24
      ListenPort = 51820
      PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
      PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE

      [Peer]
      PublicKey = %s
      AllowedIPs = 10.0.0.2/32

runcmd:
  - echo net.ipv4.ip_forward=1 >> /etc/sysctl.conf
  - sysctl -p
  - chmod 600 /etc/wireguard/wg0.conf
  - wg-quick up wg0
  - systemctl enable wg-quick@wg0
`, serverPrivateKey, clientPublicKey)

}

func BuildClientConfig(serverPublicKey, clientPrivateKey, serverIP string) string {
	return fmt.Sprintf(`[Interface]
PrivateKey = %s
Address = 10.0.0.2/32
DNS = 1.1.1.1

[Peer]
PublicKey = %s
Endpoint = %s:51820
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`, clientPrivateKey, serverPublicKey, serverIP)
}

func WaitForWireGuardReady(ip string) error {
	for i := 0; i < 60; i++ {
		conn, err := net.DialTimeout("tcp", ip+":51820", 2*time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("WireGuard server did not start in time")
}

func WaitForDropletProvision(client *godo.Client, dropletID int) (*godo.Droplet, error) {
	maxAttempts := 30
	for i := 0; i < maxAttempts; i++ {
		droplet, _, err := client.Droplets.Get(context.Background(), dropletID)

		if err != nil {
			return nil, err
		}

		if droplet.Status == "active" {
			return droplet, nil
		}

		time.Sleep(10 * time.Second)
	}
	return nil, fmt.Errorf("timeout waiting for droplet to get IP address")
}
