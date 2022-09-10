package module

import (
	//"github.com/aristanetworks/goeapi"
)

type ShowBGPEVPNSummary struct {
	VRFs map[string]VRF
}

type VRFS struct {
	RouterID string                        `json:"routerId"`
	Peers    map[string]BGPNeighborSummary `json:"peers"`
	VRF      string                        `json:"vrf"`
	ASN      string                          `json:"asn"`
}

type BGPEVPNNeighborSummary struct {
	MsgSent             int     `json:"msgSent"`
	InMsgQueue          int     `json:"inMsgQueue"`
	PrefixReceived      int     `json:"prefixReceived"`
	UpDownTime          float64 `json:"upDownTime"`
	Version             int     `json:"version"`
	MsgReceived         int     `json:"msgReceived"`
	PrefixAccepted      int     `json:"prefixAccepted"`
	PeerState           string  `json:"peerState"`
	PeerStateIdleReason string  `json:"peerStateIdleReason,omitempty"`
	OutMsgQueue         int     `json:"outMsgQueue"`
	UnderMaintenance    bool    `json:"underMaintenance"`
	ASN                 int64   `json:"asn"`
}

func (b *ShowBGPEVPNSummary) GetCmd() string {
	return "show bgp evpn summary"
}

func (s *ShowEntity) ShowBGPEVPNSummary() (ShowBGPEVPNSummary, error) {
	handle, _ := s.node.GetHandle("json")
	var showbgpevpnsmmary ShowBGPEVPNSummary
	handle.AddCommand(&showbgpevpnsmmary)

	if err := handle.Call(); err != nil {
		return showbgpevpnsmmary, err
	}

	handle.Close()
	return showbgpevpnsmmary, nil
}
