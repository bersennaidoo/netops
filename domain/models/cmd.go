package models

type CMD struct {
	Cmd string
}

func (c CMD) CreateCMDList(l LoopBackRequest) []CMD {
	var cmds []CMD
	for i := 0; i < 1; i++ {
		cmds = append(cmds, CMD{Cmd: l.Enable})
		cmds = append(cmds, CMD{Cmd: l.Pwd})
		cmds = append(cmds, CMD{Cmd: l.Conft})
		cmds = append(cmds, CMD{Cmd: l.Loop})
		cmds = append(cmds, CMD{Cmd: l.Ipaddr})
		cmds = append(cmds, CMD{Cmd: l.End})
		cmds = append(cmds, CMD{Cmd: l.Wr})
	}

	return cmds
}
