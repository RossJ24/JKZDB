#!/bin/bash
lsof -ti:$1 | xargs kill