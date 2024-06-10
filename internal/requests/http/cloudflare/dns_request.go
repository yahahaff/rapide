package cloudflare

type DnsCreateRequest struct {
	Content string `json:"content" validate:"required,ipv4"`
	Name    string `json:"name" validate:"required,max=255"`
	Proxied bool   `json:"proxied,bool"`
	Type    string `json:"type" validate:"required,oneof=A AAAA CNAME MX NS TXT"`
	Comment string `json:"comment" validate:"omitempty,max=100"`
	TTL     uint   `json:"ttl" validate:"min=0,max=3600"`
	ZoneId  string `json:"zone_id" validate:"required,min=0,max=50"`
}

type DnsUpdateRequest struct {
	Content string `json:"content" validate:"required"`
	Name    string `json:"name" validate:"required,max=255"`
	Proxied bool   `json:"proxied"`
	Type    string `json:"type" validate:"required,oneof=A AAAA CNAME MX NS TXT"`
	Comment string `json:"comment" validate:"omitempty,max=100"`
	TTL     uint   `json:"ttl" validate:"min=0,max=3600"`
	ZoneId  string `json:"zone_id" validate:"required,min=0,max=50"`
	DnsId   string `json:"dns_id" validate:"required,min=0,max=50"`
}

type DnsDeleteRequest struct {
	ZoneId string `json:"zone_id" validate:"required,min=1,max=32"`
	DnsId  string `json:"dns_id" validate:"required,min=1,max=32"`
}
