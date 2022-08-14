#!/bin/sh

if [ -z "`/bin/pidof maxscale`" ]; then
  exit 1
fi
