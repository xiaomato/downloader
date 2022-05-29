package donwloader

type dld struct {
	concurrent uint32
}

func NewDownloader(concurrence uint32) *dld {
	return &dld{
		concurrent: concurrence,
	}
}

func (d *dld) Start(filename string)
