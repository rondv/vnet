// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id errorThread -d VecType=errorThreadVec -d Type=*errorThread github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type errorThreadVec []*errorThread

func (p *errorThreadVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]*errorThread, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *errorThreadVec) validate(new_len uint, zero **errorThread) **errorThread {
	c := elib.Index(cap(*p))
	lʹ := elib.Index(len(*p))
	l := elib.Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *errorThreadVec) validateSlowPath(zero **errorThread,
	c, l, lʹ elib.Index) **errorThread {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]*errorThread, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *errorThreadVec) Validate(i uint) **errorThread {
	return p.validate(i+1, (**errorThread)(nil))
}

func (p *errorThreadVec) ValidateInit(i uint, zero *errorThread) **errorThread {
	return p.validate(i+1, &zero)
}

func (p *errorThreadVec) ValidateLen(l uint) (v **errorThread) {
	if l > 0 {
		v = p.validate(l, (**errorThread)(nil))
	}
	return
}

func (p *errorThreadVec) ValidateLenInit(l uint, zero *errorThread) (v **errorThread) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p errorThreadVec) Len() uint { return uint(len(p)) }