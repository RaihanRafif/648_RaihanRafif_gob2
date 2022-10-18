package main

import Routers "ass2/routers"

func main() {

	var PORT = ":3000"

	Routers.StartServer().Run(PORT)
}
