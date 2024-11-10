# Arikr

**Arikr** is a CPU emulator in Go.

```
type Line struct { Cells [256]Cell }
type Grid struct { Lines [256]*Line }
type Pipe struct { Reader io.Reader, Writer io.Writer }
type Core struct { Grid *Grid, Pipe *Pipe, Opers map[Cell]Oper }
```
