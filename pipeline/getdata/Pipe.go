package getdata

import "sync"

type Inchan chan interface{}
type Outchan chan interface{}
type CmdFunc func(args ...interface{}) Inchan
type PipeCmdFunc func(in Inchan) Outchan
type Pipline struct {
	Cmd     CmdFunc
	PipeCmd PipeCmdFunc
	Count   int
}

func NewPipe() *Pipline {
	return &Pipline{Count: 1}
}

func (p *Pipline) SetCmd(c CmdFunc) {
	p.Cmd = c
}

func (p *Pipline) SetPipeCmd(c PipeCmdFunc, count int) {
	p.PipeCmd = c
	p.Count = count
}

func (p *Pipline) Exec(args ...interface{}) Outchan {
	in := p.Cmd(args)
	out := make(Outchan)
	wg := sync.WaitGroup{}
	for i := 0; i < p.Count; i++ {
		getChan := p.PipeCmd(in)
		wg.Add(1)
		go func(input Outchan) {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out

}
