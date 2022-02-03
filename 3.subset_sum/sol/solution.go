package sol

func Solution(a []int,k int) bool{
	mem := make([]bool,k+1)
	mem[0] = true
	for _,v := range a{
		for i:=k;i>=v;i--{
			if mem[i-v]{
				mem[i] = mem[i-v]
			}
		}
	}

	return mem[k]
}
