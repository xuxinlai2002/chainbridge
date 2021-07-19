#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

export KEYSTORE_PASSWORD="xxl123456";

# Exit on failure
pm2 start ./build/chainbridge --watch -- --config config.json --verbosity trace --latest
