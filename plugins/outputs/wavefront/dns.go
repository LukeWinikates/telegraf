package wavefront

import (
	"github.com/miekg/dns"
)

// TODO: actually implement
func discoverServersViaDNS(dnsQuery string) ([]string, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(dnsQuery), dns.TypeSRV)
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, "DNSHOST???")
	if err != nil {
		return nil, err
	}
	if r.Rcode != dns.RcodeSuccess {
		return nil, err
	}

	// something something get the DNS server responses from rr.Answer
	return []string{}, nil
}
