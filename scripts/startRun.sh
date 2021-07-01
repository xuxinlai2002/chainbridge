#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


export KEYSTORE_PASSWORD="xxl123456";

# Exit on failure
./build/chainbridge --config config.json --verbosity trace --latest
