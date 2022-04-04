package pumps

import (
	"os"
)

type Pumps struct {
    Count string
}

func (p *Pumps) GetPumps() error {
    body, err := os.ReadFile("./pumps")

    if err != nil {
        return err
    }

    p.Count = string(body)

    return nil
}

func (p *Pumps) UpdatePumpCount(count string) error {
    p.Count = count
    return os.WriteFile("./pumps", []byte(p.Count), 0600)
}

func (p *Pumps) ResetFileCount() error {
    return os.WriteFile("./pumps", []byte("0"), 0600)
}
