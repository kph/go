// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package mod

import (
	"fmt"
	"net"

	"github.com/platinasystems/go/goes/cmd/ip/options"
	"github.com/platinasystems/go/goes/lang"
	"github.com/platinasystems/go/internal/parms"
)

const (
	Name    = "add | change | delete | replace"
	Apropos = "link address"
	Usage   = `
	ip neighbor { add | change | delete | replace }
		{ ADDR [ lladdr LLADDR ] [ nud STATE ] | proxy ADDR }
		[ dev DEV ]

	STATE := { permanent | noarp | stale | reachable | none | incomplete |
		delay | probe | failed }
	`
	Man = `
SEE ALSO
	ip man neighbor || ip neighbor -man
`
)

var (
	apropos = lang.Alt{
		lang.EnUS: Apropos,
	}
	man = lang.Alt{
		lang.EnUS: Man,
	}
	theseParms = []string{"lladdr", "nud", "proxy", "dev"}
)

func New(s string) Command { return Command(s) }

type Command string

func (Command) Apropos() lang.Alt { return apropos }
func (Command) Man() lang.Alt     { return man }
func (c Command) String() string  { return string(c) }
func (Command) Usage() string     { return Usage }

func (c Command) Main(args ...string) error {
	var ifaddr net.IP
	ipFlag, ipParm, args := options.New(args)
	parm, args := parms.New(args, theseParms...)

	switch len(args) {
	case 0:
		return fmt.Errorf("ADDR: missing")
	case 1:
		ifaddr = net.ParseIP(args[0])
	default:
		return fmt.Errorf("%v: unexpected", args[1:])
	}

	if ifaddr == nil {
		return fmt.Errorf("%s: invalid", args[0])
	}

	fmt.Print("FIXME ", c, " ", ifaddr)

	_ = ipFlag
	_ = ipParm
	_ = parm

	return nil
}
