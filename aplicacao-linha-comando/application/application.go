package application

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line - Cícero"
	app.Usage = "Buscando ips"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de ips na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
