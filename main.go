package main

var size int = 100
var factor float64 = 0.1
var steps int = 100

func main() {

}

func step(in Net, fc FireCalculater, f Firer) (out Net) {
	for i, node := range in.Nodes {
		if fc.Calculate(node.Neuron) {
			f.Fire(out.Nodes[i])
		} else {
			f.UnFire(out.Nodes[i])
		}
	}
	return
}

func simulate(nc NetCreater, np NetPrinter, fc FireCalculater, f Firer) {
	var nets [2]Net = [2]Net{nc.Create(size, factor), nc.Create(size, factor)}
	var curr int = 0
	var next int = 1
	np.Print(nets[curr])
	for i := 1; i < steps; i += 1 {
		nets[next] = step(nets[curr], fc, f)
		next, curr = curr, next
		np.Print(nets[curr])
	}
}
