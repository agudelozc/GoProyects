package facade

type CPU struct{}

func (c *CPU) Freeze() {}
func (c *CPU) Jump()   {}
func (c *CPU) Execute() {}

type Memory struct{}

func (m *Memory) Load(addr uint64) {}

type HardDrive struct{}

func (h *HardDrive) Read() {}

type ComputerFacade struct {
    cpu       *CPU
    memory    *Memory
    hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:       &CPU{},
        memory:    &Memory{},
        hardDrive: &HardDrive{},
    }
}

func (cf *ComputerFacade) Start() {
    cf.cpu.Freeze()
    cf.memory.Load(0)
    cf.cpu.Jump()
    cf.cpu.Execute()
}
