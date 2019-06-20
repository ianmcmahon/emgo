package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"
)

type Bits struct {
	Name  string
	Mask  uint32
	LSL   uint
	Descr string
	Val   bool
}

type Register struct {
	Offset  int
	BitSiz  int
	Name    string
	Len     int
	Descr   string
	Bits    []*Bits
	SubRegs []*Register
}

type Instance struct {
	Name  string
	Base  string
	Descr string
}

type Periph struct {
	Name  string
	Descr string
	Insts []*Instance
	Regs  []*Register
}

func (p *Periph) Save(base, pkgname string) {
	w := create(strings.ToLower(p.Name + ".go"))
	defer w.Close()
	fmt.Fprintf(w, "// Peripheral: %s_Periph  %s.\n", p.Name, p.Descr)
	fmt.Fprintln(w, "// Instances:")
	tw := new(tabwriter.Writer)
	tw.Init(w, 0, 0, 1, ' ', 0)
	for _, inst := range p.Insts {
		fmt.Fprintf(tw, "//  %s\t %s\t", inst.Name, inst.Base)
		if inst.Descr != "" {
			fmt.Fprintf(tw, "%s.\n", inst.Descr)
		} else {
			fmt.Fprintln(tw)
		}
	}
	tw.Flush()
	fmt.Fprintln(w, "// Registers:")
	for _, r := range p.Regs {
		if r == nil {
			continue
		}
		fmt.Fprintf(tw, "//  0x%02X\t%2d\t ", r.Offset, r.BitSiz)
		name := r.Name
		if len(r.SubRegs) > 0 {
			subregs := r.SubRegs[0].Name
			for _, sr := range r.SubRegs[1:] {
				subregs += "," + sr.Name
			}
			name += "{" + subregs + "}"
		}
		if r.Len == 0 {
			fmt.Fprintf(tw, "%s\t", name)
		} else {
			fmt.Fprintf(tw, "%s[%d]\t", name, r.Len)
		}
		if r.Descr != "" {
			fmt.Fprintf(tw, "%s.\n", r.Descr)
		} else {
			fmt.Fprintln(tw)
		}
	}
	tw.Flush()
	fmt.Fprintln(w, "// Import:")
	fmt.Fprintln(w, "// ", base+"/mmap")

	fmt.Fprintln(w, "package", pkgname)
	w.donotedit()
	saveBits(w, p.Regs)
}

func saveBits(w io.Writer, regs []*Register) {
	for _, r := range regs {
		if r == nil {
			continue
		}
		if len(r.SubRegs) > 0 {
			saveBits(w, r.SubRegs)
			continue
		}
		if len(r.Bits) == 0 {
			continue
		}
		fmt.Fprintln(w, "\nconst (")
		for _, b := range r.Bits {
			fmt.Fprintf(
				w, "\t%s %s = 0x%02X << %d",
				b.Name, r.Name, b.Mask, b.LSL,
			)
			switch {
			case b.Descr != "":
				if b.Val {
					fmt.Fprintf(w, " //  %s.\n", b.Descr)
				} else {
					fmt.Fprintf(w, " //+ %s.\n", b.Descr)
				}
			case !b.Val:
				fmt.Fprintln(w, " //+")
			default:
				fmt.Fprintln(w)
			}
		}
		fmt.Fprintln(w, ")")
		fmt.Fprintln(w, "\nconst (")
		for _, b := range r.Bits {
			if !b.Val {
				fmt.Fprintf(w, "\t%sn = %d\n", b.Name, b.LSL)
			}
		}
		fmt.Fprintln(w, ")")
	}
}

type Package struct {
	Name    string
	Descr   string
	Periphs []*Periph
}

func (pkg *Package) saveDoc() {
	w := create("0_doc.go")
	defer w.Close()
	fmt.Fprintf(
		w, "// Package %s provides interface to %s.\npackage %s\n",
		pkg.Name, pkg.Descr, pkg.Name,
	)
	w.donotedit()
}

func (pkg *Package) Save(base string) {
	mkdir(pkg.Name)
	chdir(pkg.Name)
	defer chdir("..")
	pkg.saveDoc()
	for _, periph := range pkg.Periphs {
		periph.Save(base, pkg.Name)
	}
}

func peripherals(r *scanner) []*Package {
	var (
		pkgs   []*Package
		pkg    *Package
		brief  string
		pbase  string
		regs   []*Register
		offset int
	)
	for r.Scan() {
		line := strings.TrimSpace(r.Text())
		if bri := doxy(line, "@brief"); bri != "" {
			brief = bri
			continue
		}
		if uintx := strings.Index(line, "uint"); uintx >= 0 {
			ioreg := strings.HasPrefix(line, "__I") // True for IO register.
			line = line[uintx:]
			n := strings.IndexByte(line, ';')
			if n < 0 {
				r.Die("';' expected after register name")
			}
			tr := strings.Fields(line[:n])
			if len(tr) != 2 {
				r.Die("wrong number of fields before ';'")
			}
			line = line[n+1:]
			typ, reg := tr[0], tr[1]
			var size, bitsiz int
			switch typ {
			case "uint32_t":
				bitsiz = 32
				size = 4
			case "uint16_t":
				bitsiz = 16
				size = 2
			case "uint8_t":
				bitsiz = 8
				size = 1
			default:
				r.Die("unknown type:", typ)
			}
			var length int
			if n := len(reg) - 1; n >= 0 && reg[n] == ']' {
				m := strings.Index(reg, "[")
				if m < 0 {
					die("Bad register name:", reg)
				}
				n, err := strconv.ParseUint(reg[m+1:n], 0, 32)
				checkErr(err)
				reg = reg[:m]
				length = int(n)
				size *= length
			}
			if !ioreg {
				offset += size
				continue
			}
			var descr string
			if n := strings.Index(line, "/*"); n > 0 {
				descr = strings.TrimPrefix(line[n+2:], "!<")
				if n := strings.LastIndex(descr, "ddress offset:"); n > 0 {
					descr = descr[:n-1]
				} else if n := strings.LastIndex(descr, "*/"); n >= 0 {
					descr = descr[:n]
				}
				descr = strings.TrimSpace(descr)
				if n := len(descr); n > 0 {
					n--
					switch descr[n] {
					case '.', ',', ';':
						descr = descr[:n]
					}
				}
			}
			regs = append(
				regs,
				&Register{
					Offset: offset,
					BitSiz: bitsiz,
					Name:   reg,
					Len:    length,
					Descr:  descr,
				},
			)
			offset += size
			continue
		}
		if strings.HasPrefix(line, "}") {
			line = strings.TrimSpace(line[1:])
			var n int
			if n = strings.Index(line, "TypeDef;"); n > 0 {
				if line[n-1] == '_' {
					n--
				}
			} else {
				r.Die("name of type (*TypeDef) expected after '}'")
			}
			periph := line[:n]
			pb := periph
			if n := strings.IndexByte(pb, '_'); n > 0 {
				pb = periph[:n]
			}
			if pbase != pb {
				pbase = pb
				pkg = &Package{Name: strings.ToLower(pbase)}
				pkgs = append(pkgs, pkg)
			}
			if periph == pbase {
				pkg.Descr = brief
			}
			for _, reg := range regs {
				d := reg.Descr
				if n := strings.IndexFunc(d, unicode.IsSpace); n > 0 {
					if d[:n] == periph {
						d = strings.TrimSpace(d[n+1:])
						reg.Descr = upperFirst(d)
					}
				}
			}
			p := &Periph{Name: periph, Descr: brief, Regs: regs}
			tweakPeriph(p)
			pkg.Periphs = append(pkg.Periphs, p)
			regs = nil
			offset = 0
			continue
		}
		if doxy(line, "@addtogroup") != "" {
			break
		}
	}
	checkErr(r.Err())
	return pkgs
}

func tweakPeriph(p *Periph) {
	switch p.Name {
	case "FSMC_Bank1", "FMC_Bank1":
		fsmcBank1(p)
	case "DSI":
		tweakDsi(p)
	}
}

func fsmcBank1(p *Periph) {
	for _, r := range p.Regs {
		if r.Name == "BTCR" {
			bcr := &Register{
				Offset: r.Offset,
				BitSiz: r.BitSiz,
				Name:   "BCR1",
				Descr:  "NOR/PSRAM chip-select control register",
			}
			btr := &Register{
				Offset: r.Offset,
				BitSiz: r.BitSiz,
				Name:   "BTR1",
				Descr:  "NOR/PSRAM chip-select timing register",
			}
			p.Regs = append(p.Regs, bcr, btr)
			return
		}
	}
}

func tweakDsi(p *Periph) {
	for _, r := range p.Regs {
		// without this, the autogenerated R version of this, RMCR, conflicts with the autogenerated RM version of CR, RMCR
		if r.Name == "MCR" {
			r.Name = "M_CR"
		}
	}
}
