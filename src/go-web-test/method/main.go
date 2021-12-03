package method

const (
	WHITE  = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color  Color
}

func (b Box) Color() Color {
	return b.color
}


type Option func(*Box)

func NewBox(opt ...Option) (Box) {
	r := new(Box)
	for _, o := range opt {
		o(r)
	}
	return *r
}


func WriteWidth(s float64) Option {
	return func(o *Box) {
		o.width = s
	}
}

func WriteHeight(s float64) Option {
	return func(o *Box) {
		o.height = s
	}
}

func WriteDepth(s float64) Option {
	return func(o *Box) {
		o.depth = s
	}
}


func WriteColor(s Color) Option {
	return func(o *Box) {
		o.color = s
	}
}


type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack()  {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}




