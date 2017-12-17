#!/bin/bash
gnome-terminal -x bash -c "./bin/main_go 00000000 8080 8081 2>&1; bash"
gnome-terminal -x bash -c "./bin/main_go 00000001 8082 8083 2>&1; bash"
gnome-terminal -x bash -c "./bin/main_go 00000002 8084 8085 2>&1; bash"

