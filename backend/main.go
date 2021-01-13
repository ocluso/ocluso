/*
 * Copyright (C) 2021 The ocluso Authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ocluso/ocluso/backend/core"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ocluso PATH-TO-CONFIG-FILE")
		os.Exit(1)
	}

	config, err := core.LoadConfig(os.Args[1])
	handleErr(err)

	server, err := core.NewServer(*config)
	handleErr(err)

	go func() {
		err := server.Serve()
		handleErr(err)
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	signal := <-quitChannel
	log.Println("Received", signal)

	err = server.Shutdown()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
